package seeds

import (
	"eztrip/api-go/auth0"
	"eztrip/api-go/logger"
	"eztrip/api-go/rbac"
	"eztrip/api-go/user"

	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	defaultPassword = "TempPassword123!"
	roleAdmin       = "admin"
	roleUser        = "user"
)

type userWithRole struct {
	user user.User
	role string
}

// SeedUsers populates the users table with sample data and assigns RBAC roles
func SeedUsers(db *gorm.DB, enforcer *casbin.Enforcer) error {
	if usersAlreadyExist(db) {
		logger.Log.Info("Users already seeded, skipping...")
		return nil
	}

	logger.Log.Info("Seeding users...")

	auth0Client := initializeAuth0Client()
	usersToSeed := buildUserList()

	seedAllUsers(db, enforcer, auth0Client, usersToSeed)

	logSeedingSummary(usersToSeed)
	return nil
}

func usersAlreadyExist(db *gorm.DB) bool {
	var count int64
	db.Model(&user.User{}).Count(&count)
	return count > 0
}

func initializeAuth0Client() *auth0.Client {
	auth0Client, err := auth0.NewClient()
	if err != nil {
		logger.Log.WithField("error", err.Error()).Warn("Auth0 client not configured, seeding without Auth0 IDs")
		return nil
	}
	return auth0Client
}

func buildUserList() []userWithRole {
	adminUsers := []user.User{
		{FirstName: "Admin", LastName: "One", Email: "admin1@example.com"},
		{FirstName: "Admin", LastName: "Two", Email: "admin2@example.com"},
	}

	regularUsers := []user.User{
		{FirstName: "John", LastName: "Doe", Email: "john@example.com"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
		{FirstName: "Alice", LastName: "Johnson", Email: "alice@example.com"},
		{FirstName: "Bob", LastName: "Williams", Email: "bob@example.com"},
		{FirstName: "Charlie", LastName: "Brown", Email: "charlie@example.com"},
	}

	usersToSeed := make([]userWithRole, 0, len(adminUsers)+len(regularUsers))

	for _, u := range adminUsers {
		usersToSeed = append(usersToSeed, userWithRole{user: u, role: roleAdmin})
	}

	for _, u := range regularUsers {
		usersToSeed = append(usersToSeed, userWithRole{user: u, role: roleUser})
	}

	return usersToSeed
}

func seedAllUsers(db *gorm.DB, enforcer *casbin.Enforcer, auth0Client *auth0.Client, usersToSeed []userWithRole) {
	for _, uwr := range usersToSeed {
		if err := seedUserWithRole(db, enforcer, auth0Client, &uwr.user, uwr.role); err != nil {
			logger.Log.WithFields(logrus.Fields{
				"email": uwr.user.Email,
				"role":  uwr.role,
				"error": err.Error(),
			}).Warn("Failed to seed user")
		}
	}
}

func logSeedingSummary(usersToSeed []userWithRole) {
	adminCount := 0
	userCount := 0

	for _, uwr := range usersToSeed {
		if uwr.role == roleAdmin {
			adminCount++
		} else {
			userCount++
		}
	}

	logger.Log.WithFields(logrus.Fields{
		"total":  len(usersToSeed),
		"admins": adminCount,
		"users":  userCount,
	}).Info("Successfully seeded users with roles")
}

func seedUserWithRole(db *gorm.DB, enforcer *casbin.Enforcer, auth0Client *auth0.Client, u *user.User, role string) error {
	auth0ID := getOrCreateAuth0User(auth0Client, u)
	u.Auth0UserID = auth0ID

	if err := createUserInDatabase(db, u); err != nil {
		return err
	}

	if err := assignRoleToUser(enforcer, u, role); err != nil {
		return err
	}

	logSuccessfulSeed(u, role, auth0ID)
	return nil
}

func getOrCreateAuth0User(auth0Client *auth0.Client, u *user.User) *string {
	if auth0Client == nil {
		return nil
	}

	auth0User, err := auth0Client.CreateUser(u.Email, defaultPassword, u.FirstName, u.LastName)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"email": u.Email,
			"error": err.Error(),
		}).Warn("Failed to create user in Auth0, checking if user already exists")

		return fetchExistingAuth0User(auth0Client, u.Email)
	}

	logger.Log.WithFields(logrus.Fields{
		"email":    u.Email,
		"auth0_id": auth0User.UserID,
	}).Info("Created user in Auth0")

	return &auth0User.UserID
}

func fetchExistingAuth0User(auth0Client *auth0.Client, email string) *string {
	existingUser, err := auth0Client.GetUserByEmail(email)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"email": email,
			"error": err.Error(),
		}).Warn("Failed to fetch existing user from Auth0, continuing without Auth0 ID")
		return nil
	}

	logger.Log.WithFields(logrus.Fields{
		"email":    email,
		"auth0_id": existingUser.UserID,
	}).Info("Found existing user in Auth0")

	return &existingUser.UserID
}

func createUserInDatabase(db *gorm.DB, u *user.User) error {
	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func assignRoleToUser(enforcer *casbin.Enforcer, u *user.User, role string) error {
	if err := rbac.AddRoleForUser(enforcer, u.ID, role); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"user_id": u.ID,
			"email":   u.Email,
			"role":    role,
			"error":   err.Error(),
		}).Warn("Failed to assign role")
		return err
	}
	return nil
}

func logSuccessfulSeed(u *user.User, role string, auth0ID *string) {
	logger.Log.WithFields(logrus.Fields{
		"user_id":  u.ID,
		"email":    u.Email,
		"role":     role,
		"auth0_id": auth0ID,
	}).Info("Seeded user successfully")
}
