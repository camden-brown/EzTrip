package seeds

import (
	"fmt"

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
	logger.Log.Info("Seeding users...")

	auth0Client := initializeAuth0Client()
	
	// Sync existing Auth0 users to database
	if auth0Client != nil {
		if err := syncAuth0Users(db, enforcer, auth0Client); err != nil {
			logger.Log.WithError(err).Warn("Failed to sync Auth0 users, continuing with local seeds")
		}
	}

	// Seed additional local test users
	usersToSeed := buildUserList()
	seedAllUsers(db, enforcer, usersToSeed)

	logSeedingSummary(usersToSeed)
	return nil
}

func initializeAuth0Client() *auth0.Client {
	auth0Client, err := auth0.NewClient()
	if err != nil {
		logger.Log.WithField("error", err.Error()).Warn("Auth0 client not configured, skipping Auth0 sync")
		return nil
	}
	return auth0Client
}

func syncAuth0Users(db *gorm.DB, enforcer *casbin.Enforcer, auth0Client *auth0.Client) error {
	logger.Log.Info("Fetching users from Auth0...")
	
	auth0Users, err := auth0Client.ListUsers()
	if err != nil {
		return fmt.Errorf("failed to list Auth0 users: %w", err)
	}

	logger.Log.WithField("count", len(auth0Users)).Info("Found Auth0 users")

	for _, auth0User := range auth0Users {
		// Check if user already exists in database by auth0_user_id
		var existingUser user.User
		err := db.Where("auth0_user_id = ?", auth0User.UserID).First(&existingUser).Error
		if err == nil {
			logger.Log.WithFields(logrus.Fields{
				"email":    auth0User.Email,
				"auth0_id": auth0User.UserID,
			}).Debug("Auth0 user already exists in database")
			continue
		}

		// Check if user exists by email (from local seeds)
		err = db.Where("email = ?", auth0User.Email).First(&existingUser).Error
		if err == nil {
			// Update existing user with Auth0 ID
			existingUser.Auth0UserID = &auth0User.UserID
			existingUser.FirstName = auth0User.FirstName
			existingUser.LastName = auth0User.LastName
			if err := db.Save(&existingUser).Error; err != nil {
				logger.Log.WithFields(logrus.Fields{
					"email": auth0User.Email,
					"error": err.Error(),
				}).Warn("Failed to update user with Auth0 ID")
				continue
			}

			logger.Log.WithFields(logrus.Fields{
				"user_id":  existingUser.ID,
				"email":    auth0User.Email,
				"auth0_id": auth0User.UserID,
			}).Info("Updated existing user with Auth0 ID")
			continue
		}

		// Create new user from Auth0
		newUser := user.User{
			Auth0UserID: &auth0User.UserID,
			FirstName:   auth0User.FirstName,
			LastName:    auth0User.LastName,
			Email:       auth0User.Email,
		}

		if err := db.Create(&newUser).Error; err != nil {
			logger.Log.WithFields(logrus.Fields{
				"email": auth0User.Email,
				"error": err.Error(),
			}).Warn("Failed to create user from Auth0")
			continue
		}

		// Assign default user role
		if err := rbac.AddRoleForUser(enforcer, newUser.ID, roleUser); err != nil {
			logger.Log.WithFields(logrus.Fields{
				"user_id": newUser.ID,
				"error":   err.Error(),
			}).Warn("Failed to assign role to Auth0 user")
		}

		logger.Log.WithFields(logrus.Fields{
			"user_id":  newUser.ID,
			"email":    auth0User.Email,
			"auth0_id": auth0User.UserID,
		}).Info("Synced Auth0 user to database")
	}

	return nil
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

func seedAllUsers(db *gorm.DB, enforcer *casbin.Enforcer, usersToSeed []userWithRole) {
	for _, uwr := range usersToSeed {
		if err := seedUserWithRole(db, enforcer, &uwr.user, uwr.role); err != nil {
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

func seedUserWithRole(db *gorm.DB, enforcer *casbin.Enforcer, u *user.User, role string) error {
	// Check if user already exists in database by email
	var existingUser user.User
	err := db.Where("email = ?", u.Email).First(&existingUser).Error
	if err == nil {
		logger.Log.WithFields(logrus.Fields{
			"email": u.Email,
		}).Info("User already exists in database, skipping")
		return nil
	}

	// Real users will be auto-provisioned with Auth0 ID on first login
	if err := createUserInDatabase(db, u); err != nil {
		return err
	}

	if err := assignRoleToUser(enforcer, u, role); err != nil {
		return err
	}

	logSuccessfulSeed(u, role, nil)
	return nil
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
