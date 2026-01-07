package seeds

import (
	"time"

	"eztrip/api-go/logger"
	"eztrip/api-go/trip"
	"eztrip/api-go/user"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// SeedTrips populates the trips, itinerary_days, and activities tables with sample data
func SeedTrips(db *gorm.DB) error {
	logger.Log.Info("Seeding trips...")

	// Check if trips already exist
	var count int64
	if err := db.Model(&trip.Trip{}).Count(&count).Error; err != nil {
		logger.Log.WithError(err).Error("Failed to count existing trips")
		return err
	}

	if count > 0 {
		logger.Log.WithField("count", count).Info("Trips already exist, skipping seed")
		return nil
	}

	// Get the first user to assign as trip owner
	var owner user.User
	if err := db.First(&owner).Error; err != nil {
		logger.Log.WithError(err).Error("No users found - run user seeds first")
		return err
	}

	ownerID, err := uuid.Parse(owner.ID.String())
	if err != nil {
		logger.Log.WithError(err).Error("Failed to parse owner ID")
		return err
	}

	// Create Kauai trip
	kauaiTrip := createKauaiTrip(ownerID)

	if err := db.Create(&kauaiTrip).Error; err != nil {
		logger.Log.WithFields(logrus.Fields{
			"trip":  kauaiTrip.Title,
			"error": err.Error(),
		}).Error("Failed to create trip")
		return err
	}

	logger.Log.WithFields(logrus.Fields{
		"trip_id": kauaiTrip.ID,
		"title":   kauaiTrip.Title,
		"days":    len(kauaiTrip.Itinerary),
	}).Info("Trip seeded successfully")

	return nil
}

func createKauaiTrip(ownerID uuid.UUID) trip.Trip {
	startDate := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2026, 1, 16, 0, 0, 0, 0, time.UTC)

	return trip.Trip{
		OwnerID:     ownerID,
		Title:       "Kauai Family Adventure",
		Destination: "Kauai, Hawaii",
		StartDate:   startDate,
		EndDate:     endDate,
		Travelers:   4,
		Itinerary: []trip.ItineraryDay{
			createDay1(startDate),
			createDay2(startDate.AddDate(0, 0, 1)),
			createDay3(startDate.AddDate(0, 0, 2)),
			createDay4(startDate.AddDate(0, 0, 3)),
			createDay5(startDate.AddDate(0, 0, 4)),
			createDay6(startDate.AddDate(0, 0, 5)),
			createDay7(startDate.AddDate(0, 0, 6)),
		},
	}
}

func createDay1(date time.Time) trip.ItineraryDay {
	return trip.ItineraryDay{
		Date:      date,
		DayNumber: 1,
		Activities: []trip.Activity{
			{
				Type:        trip.ActivityTypeTransport,
				Time:        time.Date(2026, 1, 10, 10, 30, 0, 0, time.UTC),
				Title:       "Arrive at Lihue Airport",
				Location:    "Lihue Airport (LIH)",
				Category:    trip.ActivityCategoryTransport,
				Description: "Pick up rental car and head to the hotel",
				Notes:       "Car rental counters are at the airport. Pre-book for best rates.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 10, 12, 30, 0, 0, time.UTC),
				Title:       "Kalapaki Beach Hut",
				Location:    "3474 Rice St, Lihue",
				Category:    trip.ActivityCategoryFood,
				Description: "Casual beachside spot with great fish tacos - perfect for the kids",
				Notes:       "Outdoor seating with ocean breeze. Try the fish tacos and the poke bowl.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 10, 14, 0, 0, 0, time.UTC),
				Title:       "Kalapaki Beach",
				Location:    "Kalapaki Beach, Lihue",
				Category:    trip.ActivityCategoryBeach,
				Description: "Gentle waves perfect for the boys. Boogie boards available for rent. Protected swimming area with lifeguard on duty.",
				Notes:       "Bring sunscreen and water shoes. Restrooms and showers available.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 10, 18, 0, 0, 0, time.UTC),
				Title:       "Dinner at Hukilau Lanai",
				Location:    "520 Aleka Loop, Kapaa",
				Category:    trip.ActivityCategoryFood,
				Description: "Family-friendly Hawaiian cuisine with live music",
				Notes:       "Reservations recommended. Live music on certain nights. Great for families.",
			},
		},
	}
}

func createDay2(date time.Time) trip.ItineraryDay {
	return trip.ItineraryDay{
		Date:      date,
		DayNumber: 2,
		Activities: []trip.Activity{
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 11, 8, 0, 0, 0, time.UTC),
				Title:       "Breakfast at Tip Top Cafe",
				Location:    "3173 Akahi St, Lihue",
				Category:    trip.ActivityCategoryFood,
				Description: "Local favorite - try the macadamia nut pancakes!",
				Notes:       "Cash only. Opens early at 6:30 AM. Popular with locals.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 11, 9, 30, 0, 0, time.UTC),
				Title:       "Sleeping Giant Trail (Nounou East)",
				Location:    "Nounou Trail East, Kapaa",
				Category:    trip.ActivityCategoryHike,
				Description: "2-mile moderate hike with stunning views. Doable for 7-year-old, might need to carry the 4-year-old for parts.",
				Notes:       "Start early to avoid heat. Bring plenty of water and snacks. Trail can be muddy after rain.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 11, 13, 0, 0, 0, time.UTC),
				Title:       "Lunch at Opakapaka Grill",
				Location:    "4-1543 Kuhio Hwy, Kapaa",
				Category:    trip.ActivityCategoryFood,
				Description: "Fresh poke bowls and local plates",
				Notes:       "Great ahi poke. Casual outdoor seating.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 11, 15, 0, 0, 0, time.UTC),
				Title:       "Lydgate Beach Park",
				Location:    "Lydgate State Park, Kapaa",
				Category:    trip.ActivityCategoryBeach,
				Description: "Protected swimming area perfect for young kids. Has a great playground too!",
				Notes:       "Large playground, picnic areas, and protected pools make this ideal for families. Showers and restrooms available.",
			},
		},
	}
}

func createDay3(date time.Time) trip.ItineraryDay {
	return trip.ItineraryDay{
		Date:      date,
		DayNumber: 3,
		Activities: []trip.Activity{
			{
				Type:        trip.ActivityTypeTransport,
				Time:        time.Date(2026, 1, 12, 7, 30, 0, 0, time.UTC),
				Title:       "Drive to North Shore",
				Location:    "Kapaa to Hanalei",
				Category:    trip.ActivityCategoryTransport,
				Description: "Scenic 45-minute drive with beautiful coastal views",
				Notes:       "Stop at scenic overlooks along the way. One-lane bridges require courtesy driving.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 12, 8, 30, 0, 0, time.UTC),
				Title:       "Breakfast at Hanalei Bread Company",
				Location:    "5-5161 Kuhio Hwy, Hanalei",
				Category:    trip.ActivityCategoryFood,
				Description: "Amazing pastries and breakfast sandwiches",
				Notes:       "Get there early - popular spot. Try the chocolate croissants.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 12, 10, 0, 0, 0, time.UTC),
				Title:       "Hanalei Bay Beach",
				Location:    "Hanalei Bay, Hanalei",
				Category:    trip.ActivityCategoryBeach,
				Description: "Iconic crescent bay with calm waters - perfect for families. Stunning mountain backdrop.",
				Notes:       "Arrive early for parking. Lifeguard on duty. Great for bodyboarding and swimming.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 12, 13, 0, 0, 0, time.UTC),
				Title:       "Lunch at Hanalei Taro & Juice",
				Location:    "5-5070 Kuhio Hwy, Hanalei",
				Category:    trip.ActivityCategoryFood,
				Description: "Fresh smoothie bowls and healthy island fare",
				Notes:       "Try the acai bowl. Outdoor seating with mountain views.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 12, 15, 30, 0, 0, time.UTC),
				Title:       "Anini Beach",
				Location:    "Anini Beach Park",
				Category:    trip.ActivityCategoryBeach,
				Description: "Shallow protected reef - safest swimming on the island for young kids",
				Notes:       "Best snorkeling spot for beginners. Very shallow water. Bring reef-safe sunscreen.",
			},
		},
	}
}

func createDay4(date time.Time) trip.ItineraryDay {
	return trip.ItineraryDay{
		Date:      date,
		DayNumber: 4,
		Activities: []trip.Activity{
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 13, 8, 0, 0, 0, time.UTC),
				Title:       "Breakfast at Java Kai",
				Location:    "4-1384 Kuhio Hwy, Kapaa",
				Category:    trip.ActivityCategoryFood,
				Description: "Great coffee and breakfast burritos",
				Notes:       "Convenient drive-through option. Strong coffee for early mornings.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 13, 9, 0, 0, 0, time.UTC),
				Title:       "Wailua Falls",
				Location:    "Wailua Falls Overlook",
				Category:    trip.ActivityCategoryActivity,
				Description: "Stunning 80-foot waterfall - viewable from roadside parking area. No hiking required!",
				Notes:       "Best views in the morning. Short walk from parking. Great photo opportunity.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 13, 11, 0, 0, 0, time.UTC),
				Title:       "Opaekaa Falls",
				Location:    "Kuamoo Rd, Kapaa",
				Category:    trip.ActivityCategoryActivity,
				Description: "Easy roadside viewing of beautiful waterfall",
				Notes:       "Paved parking and viewing area. Combine with Wailua Falls in same morning.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 13, 12, 30, 0, 0, time.UTC),
				Title:       "Lunch at Hamura Saimin",
				Location:    "2956 Kress St, Lihue",
				Category:    trip.ActivityCategoryFood,
				Description: "Famous local saimin noodle shop",
				Notes:       "Cash only. Often a line but moves fast. Kid-friendly noodles.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 13, 14, 30, 0, 0, time.UTC),
				Title:       "Poipu Beach",
				Location:    "Poipu Beach Park",
				Category:    trip.ActivityCategoryBeach,
				Description: "South shore beach with monk seal sightings! Protected cove perfect for kids.",
				Notes:       "Look for sea turtles and monk seals. Lifeguard on duty. Great facilities.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 13, 18, 30, 0, 0, time.UTC),
				Title:       "Dinner at Eating House 1849",
				Location:    "2829 Ala Kalanikaumaka St, Poipu",
				Category:    trip.ActivityCategoryFood,
				Description: "Roy Yamaguchi's restaurant with island-fusion cuisine",
				Notes:       "Reservations required. Great kids menu. Try the macadamia nut crusted fish.",
			},
		},
	}
}

func createDay5(date time.Time) trip.ItineraryDay {
	return trip.ItineraryDay{
		Date:      date,
		DayNumber: 5,
		Activities: []trip.Activity{
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 14, 7, 0, 0, 0, time.UTC),
				Title:       "Early Morning Na Pali Coast Boat Tour",
				Location:    "Port Allen Harbor",
				Category:    trip.ActivityCategoryActivity,
				Description: "Snorkel tour along the Na Pali Coast - dolphins, sea turtles, and dramatic cliffs. Book with Blue Dolphin or Captain Andy's.",
				Notes:       "Take seasickness meds if prone. Bring sunscreen and water. Tours are 4-5 hours. Life jackets provided for kids.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 14, 13, 0, 0, 0, time.UTC),
				Title:       "Late Lunch at Tidepools Restaurant",
				Location:    "1571 Poipu Rd, Koloa",
				Category:    trip.ActivityCategoryFood,
				Description: "Beautiful open-air restaurant over koi ponds",
				Notes:       "Reservations recommended. Scenic setting. Good for a special family lunch.",
			},
			{
				Type:        trip.ActivityTypeCustom,
				Time:        time.Date(2026, 1, 14, 15, 30, 0, 0, time.UTC),
				Title:       "Relax at Hotel Pool",
				Location:    "Hotel",
				Category:    trip.ActivityCategoryActivity,
				Description: "Rest after morning boat tour - kids can enjoy the pool",
				Notes:       "Recovery time. Let the kids burn off energy in the pool.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 14, 18, 0, 0, 0, time.UTC),
				Title:       "Casual Dinner at Brennecke's Beach Broiler",
				Location:    "2100 Hoone Rd, Poipu",
				Category:    trip.ActivityCategoryFood,
				Description: "Casual beach restaurant with ocean views",
				Notes:       "Right across from Poipu Beach. Upstairs has better views. Good fish and chips.",
			},
		},
	}
}

func createDay6(date time.Time) trip.ItineraryDay {
	return trip.ItineraryDay{
		Date:      date,
		DayNumber: 6,
		Activities: []trip.Activity{
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 15, 8, 30, 0, 0, time.UTC),
				Title:       "Breakfast at Kountry Kitchen",
				Location:    "4-1485 Kuhio Hwy, Kapaa",
				Category:    trip.ActivityCategoryFood,
				Description: "Hearty local breakfast - famous for pancakes",
				Notes:       "Often a wait on weekends. Large portions. Cash and cards accepted.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC),
				Title:       "Kilauea Lighthouse",
				Location:    "Kilauea Point National Wildlife Refuge",
				Category:    trip.ActivityCategoryActivity,
				Description: "Historic lighthouse with bird watching and ocean views. Often see whales in winter!",
				Notes:       "Small entrance fee. Paved paths, stroller-friendly. Bring binoculars for whale watching.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 15, 13, 0, 0, 0, time.UTC),
				Title:       "Lunch at Chicken in a Barrel",
				Location:    "4-1586 Kuhio Hwy, Kapaa",
				Category:    trip.ActivityCategoryFood,
				Description: "BBQ rotisserie chicken - simple and delicious",
				Notes:       "Casual outdoor seating. Great for picky eaters. Takeout available.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 15, 14, 30, 0, 0, time.UTC),
				Title:       "Secret Beach (Kauapea Beach)",
				Location:    "End of Kalihiwai Rd",
				Category:    trip.ActivityCategoryBeach,
				Description: "Requires short steep hike down but rewards with stunning beach. Better for older kids.",
				Notes:       "Watch kids carefully - strong currents. Beautiful for photos and exploring tide pools.",
			},
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 15, 18, 30, 0, 0, time.UTC),
				Title:       "Farewell Dinner at Bar Acuda",
				Location:    "5-5161 Kuhio Hwy, Hanalei",
				Category:    trip.ActivityCategoryFood,
				Description: "Tapas-style dining - order multiple dishes to share",
				Notes:       "Reservations essential. Kid-friendly small plates. Great wine selection for adults.",
			},
		},
	}
}

func createDay7(date time.Time) trip.ItineraryDay {
	return trip.ItineraryDay{
		Date:      date,
		DayNumber: 7,
		Activities: []trip.Activity{
			{
				Type:        trip.ActivityTypePlaceBased,
				Time:        time.Date(2026, 1, 16, 7, 0, 0, 0, time.UTC),
				Title:       "Breakfast at Hotel",
				Location:    "Hotel Restaurant",
				Category:    trip.ActivityCategoryFood,
				Description: "Leisurely hotel breakfast before checkout",
				Notes:       "Pack the night before. Check flight time to plan departure.",
			},
			{
				Type:        trip.ActivityTypeTransport,
				Time:        time.Date(2026, 1, 16, 10, 0, 0, 0, time.UTC),
				Title:       "Return Rental Car",
				Location:    "Lihue Airport",
				Category:    trip.ActivityCategoryTransport,
				Description: "Drop off rental car and check in for flight",
				Notes:       "Arrive 2 hours early for domestic flights. Refuel car before returning.",
			},
			{
				Type:        trip.ActivityTypeTransport,
				Time:        time.Date(2026, 1, 16, 13, 0, 0, 0, time.UTC),
				Title:       "Depart Lihue Airport",
				Location:    "Lihue Airport (LIH)",
				Category:    trip.ActivityCategoryTransport,
				Description: "Departure flight home",
				Notes:       "Pack any shells/rocks in checked luggage. Agriculture inspection required.",
			},
		},
	}
}
