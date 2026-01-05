import { Injectable } from '@angular/core';
import { Trip } from '../models/trip.model';

@Injectable({
  providedIn: 'root',
})
export class TripMockService {
  private mockTrip: Trip = {
    id: '1',
    title: 'Kauai Family Adventure',
    destination: 'Kauai, Hawaii',
    startDate: '2026-01-10',
    endDate: '2026-01-16',
    travelers: 4,
    itinerary: [
      {
        date: '2026-01-10',
        activities: [
          {
            id: '1-1',
            time: '2026-01-10T10:30:00',
            title: 'Arrive at Lihue Airport',
            location: 'Lihue Airport (LIH)',
            category: 'transport',
            description: 'Pick up rental car and head to the hotel',
            images: [],
            rating: 4.2,
            reviewCount: 856,
            address: '3901 Mokulele Loop, Lihue, HI 96766',
            website: 'https://airports.hawaii.gov/lih',
            notes:
              'Car rental counters are at the airport. Pre-book for best rates.',
          },
          {
            id: '1-2',
            time: '2026-01-10T12:30:00',
            title: 'Kalapaki Beach Hut',
            location: '3474 Rice St, Lihue',
            category: 'food',
            description:
              'Casual beachside spot with great fish tacos - perfect for the kids',
            rating: 4.5,
            reviewCount: 523,
            address: '3474 Rice St, Lihue, HI 96766',
            website: 'http://www.kalapakibeachhut.com',
            images: [
              'https://images.unsplash.com/photo-1551504734-5ee1c4a1479b?w=400',
              'https://images.unsplash.com/photo-1565299624946-b28f40a0ae38?w=400',
            ],
            notes:
              'Outdoor seating with ocean breeze. Try the fish tacos and the poke bowl.',
          },
          {
            id: '1-3',
            time: '2026-01-10T14:00:00',
            title: 'Kalapaki Beach',
            location: 'Kalapaki Beach, Lihue',
            category: 'beach',
            description:
              'Gentle waves perfect for the boys. Boogie boards available for rent. Protected swimming area with lifeguard on duty.',
            rating: 4.7,
            reviewCount: 1832,
            address: 'Rice St, Lihue, HI 96766',
            website:
              'https://www.kauai.gov/Government/Departments-Agencies/Parks-Recreation',
            images: [
              'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400',
              'https://images.unsplash.com/photo-1505142468610-359e7d316be0?w=400',
              'https://images.unsplash.com/photo-1473496169904-658ba7c44d8a?w=400',
            ],
            notes:
              'Bring sunscreen and water shoes. Restrooms and showers available.',
          },
          {
            id: '1-4',
            time: '2026-01-10T18:00:00',
            title: 'Dinner at Hukilau Lanai',
            location: '520 Aleka Loop, Kapaa',
            category: 'food',
            description: 'Family-friendly Hawaiian cuisine with live music',
            images: [
              'https://images.unsplash.com/photo-1555939594-58d7cb561ad1?w=400',
              'https://images.unsplash.com/photo-1544025162-d76694265947?w=400',
            ],
            rating: 4.6,
            reviewCount: 1245,
            address: '520 Aleka Loop, Kapaa, HI 96746',
            website: 'https://www.hukilaukauai.com',
            notes:
              'Reservations recommended. Live music on certain nights. Great for families.',
          },
        ],
      },
      {
        date: '2026-01-11',
        activities: [
          {
            id: '2-1',
            time: '2026-01-11T08:00:00',
            title: 'Breakfast at Tip Top Cafe',
            location: '3173 Akahi St, Lihue',
            category: 'food',
            description: 'Local favorite - try the macadamia nut pancakes!',
            images: [
              'https://images.unsplash.com/photo-1533089860892-a7c6f0a88666?w=400',
            ],
            rating: 4.4,
            reviewCount: 678,
            address: '3173 Akahi St, Lihue, HI 96766',
            website: 'https://www.tiptopmotel.com',
            notes: 'Cash only. Opens early at 6:30 AM. Popular with locals.',
          },
          {
            id: '2-2',
            time: '2026-01-11T09:30:00',
            title: 'Sleeping Giant Trail (Nounou East)',
            location: 'Nounou Trail East, Kapaa',
            category: 'hike',
            description:
              '2-mile moderate hike with stunning views. Doable for 7-year-old, might need to carry the 4-year-old for parts. The trail winds through a forest before opening up to panoramic views of the coastline.',
            rating: 4.6,
            reviewCount: 945,
            address: 'Haleilio Rd, Kapaa, HI 96746',
            website:
              'https://www.alltrails.com/trail/hawaii/kauai/sleeping-giant-trail',
            images: [
              'https://images.unsplash.com/photo-1551632811-561732d1e306?w=400',
              'https://images.unsplash.com/photo-1464822759023-fed622ff2c3b?w=400',
            ],
            notes:
              'Start early to avoid heat. Bring plenty of water and snacks. Trail can be muddy after rain.',
          },
          {
            id: '2-3',
            time: '2026-01-11T13:00:00',
            title: 'Lunch at Opakapaka Grill',
            location: '4-1543 Kuhio Hwy, Kapaa',
            category: 'food',
            description: 'Fresh poke bowls and local plates',
            images: [
              'https://images.unsplash.com/photo-1546069901-ba9599a7e63c?w=400',
            ],
            rating: 4.3,
            reviewCount: 412,
            address: '4-1543 Kuhio Hwy, Kapaa, HI 96746',
            website: 'https://www.opakapakagrill.com',
            notes: 'Great ahi poke. Casual outdoor seating.',
          },
          {
            id: '2-4',
            time: '2026-01-11T15:00:00',
            title: 'Lydgate Beach Park',
            location: 'Lydgate State Park, Kapaa',
            category: 'beach',
            description:
              'Protected swimming area perfect for young kids. Has a great playground too!',
            images: [
              'https://images.unsplash.com/photo-1471922694854-ff1b63b20054?w=400',
              'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=400',
            ],
            rating: 4.8,
            reviewCount: 1543,
            address: 'Leho Dr, Kapaa, HI 96746',
            website:
              'https://www.kauai.gov/Government/Departments-Agencies/Parks-Recreation',
            notes:
              'Large playground, picnic areas, and protected pools make this ideal for families. Showers and restrooms available.',
          },
        ],
      },
      {
        date: '2026-01-12',
        activities: [
          {
            id: '3-1',
            time: '2026-01-12T07:30:00',
            title: 'Drive to North Shore',
            location: 'Kapaa to Hanalei',
            category: 'transport',
            description: 'Scenic 45-minute drive with beautiful coastal views',
            images: [],
            rating: 4.9,
            reviewCount: 2345,
            address: 'Kuhio Hwy, Kauai, HI',
            website:
              'https://www.gohawaii.com/islands/kauai/regions/north-shore',
            notes:
              'Stop at scenic overlooks along the way. One-lane bridges require courtesy driving.',
          },
          {
            id: '3-2',
            time: '2026-01-12T08:30:00',
            title: 'Breakfast at Hanalei Bread Company',
            location: '5-5161 Kuhio Hwy, Hanalei',
            category: 'food',
            description: 'Amazing pastries and breakfast sandwiches',
            images: [
              'https://images.unsplash.com/photo-1509440159596-0249088772ff?w=400',
            ],
            rating: 4.7,
            reviewCount: 892,
            address: '5-5161 Kuhio Hwy, Hanalei, HI 96714',
            website: 'https://www.hanaleibroadcompany.com',
            notes: 'Arrive early as they often sell out. Try the guava danish.',
          },
          {
            id: '3-3',
            time: '2026-01-12T10:00:00',
            title: 'Hanalei Bay Beach',
            location: 'Hanalei Bay, Hanalei',
            category: 'beach',
            description:
              'One of the most beautiful beaches in Hawaii. Great for swimming and building sandcastles. Two-mile crescent bay with spectacular mountain backdrop.',
            rating: 4.9,
            reviewCount: 2156,
            address: 'Weke Rd, Hanalei, HI 96714',
            website:
              'https://www.kauai.gov/Government/Departments-Agencies/Parks-Recreation',
            images: [
              'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400',
              'https://images.unsplash.com/photo-1473496169904-658ba7c44d8a?w=400',
              'https://images.unsplash.com/photo-1505142468610-359e7d316be0?w=400',
            ],
            notes:
              'Can get crowded in peak season. Best swimming on calm days - check surf conditions.',
          },
          {
            id: '3-4',
            time: '2026-01-12T12:30:00',
            title: 'Lunch at Tahiti Nui',
            location: '5-5134 Kuhio Hwy, Hanalei',
            category: 'food',
            description: 'Historic Hawaiian restaurant with great burgers',
            images: [
              'https://images.unsplash.com/photo-1568901346375-23c9450c58cd?w=400',
            ],
            rating: 4.4,
            reviewCount: 734,
            address: '5-5134 Kuhio Hwy, Hanalei, HI 96714',
            website: 'https://www.thenui.com',
            notes:
              'Featured in the movie "The Descendants". Lively atmosphere with local charm.',
          },
          {
            id: '3-5',
            time: '2026-01-12T14:30:00',
            title: "Queen's Bath (viewpoint only)",
            location: "Queen's Bath Trail, Princeville",
            category: 'hike',
            description:
              'Short walk to see the famous tidal pool. Too dangerous for swimming with kids, but beautiful views.',
            images: [
              'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400',
            ],
            rating: 4.3,
            reviewCount: 567,
            address: 'Kapiolani Rd, Princeville, HI 96722',
            website:
              'https://www.alltrails.com/trail/hawaii/kauai/queens-bath-trail',
            notes:
              'Trail can be slippery. Only visit during calm ocean conditions. Do NOT enter the water with children.',
          },
        ],
      },
      {
        date: '2026-01-13',
        activities: [
          {
            id: '4-1',
            time: '2026-01-13T09:00:00',
            title: 'Wailua River Kayak Adventure',
            location: 'Wailua River State Park',
            category: 'activity',
            description:
              'Family-friendly kayak trip to Secret Falls. The 7-year-old can help paddle! Includes kayak rental, paddles, and life jackets. 4-5 hour round trip with short hike to waterfall.',
            rating: 4.8,
            reviewCount: 1234,
            address: '174 Wailua Rd, Kapaa, HI 96746',
            website: 'http://www.kayakkauai.com',
            images: [
              'https://images.unsplash.com/photo-1544551763-46a013bb70d5?w=400',
              'https://images.unsplash.com/photo-1502933691298-84fc14542831?w=400',
            ],
            notes:
              'Book in advance. Bring waterproof bag for phones/cameras. Wear water shoes.',
          },
          {
            id: '4-2',
            time: '2026-01-13T13:00:00',
            title: 'Lunch at Wailua Shave Ice',
            location: '4-831 Kuhio Hwy, Kapaa',
            category: 'food',
            description:
              'Best shave ice on the island - the boys will love this!',
            images: [
              'https://images.unsplash.com/photo-1563805042-7684c019e1cb?w=400',
            ],
            rating: 4.8,
            reviewCount: 1123,
            address: '4-831 Kuhio Hwy, Kapaa, HI 96746',
            website: 'https://www.uncleswailua.com',
            notes:
              'Try the lilikoi or coconut flavors. Add ice cream and mochi for extra treat.',
          },
          {
            id: '4-3',
            time: '2026-01-13T15:00:00',
            title: 'Smith Family Garden Luau',
            location: '3-5971 Kuhio Hwy, Lihue',
            category: 'entertainment',
            description:
              'Family-friendly luau with boat ride, gardens, and traditional show. Buffet dinner with Hawaiian specialties. Cultural demonstrations and hula performances.',
            rating: 4.5,
            reviewCount: 867,
            address: '3-5971 Kuhio Hwy, Kapaa, HI 96746',
            website: 'http://www.smithskauai.com',
            images: [
              'https://images.unsplash.com/photo-1533174072545-7a4b6ad7a6c3?w=400',
              'https://images.unsplash.com/photo-1530789253388-582c481c54b0?w=400',
            ],
            notes:
              'Reservations required. Show starts at 5 PM, arrive by 4:30 PM for garden tour.',
          },
        ],
      },
      {
        date: '2026-01-14',
        activities: [
          {
            id: '5-1',
            time: '2026-01-14T08:00:00',
            title: 'Drive to Waimea Canyon',
            location: 'Waimea Canyon Drive',
            category: 'transport',
            description: "The 'Grand Canyon of the Pacific' - stunning views!",
            images: [],
            rating: 4.9,
            reviewCount: 1987,
            address: 'Waimea Canyon Dr, Waimea, HI 96796',
            website:
              'https://dlnr.hawaii.gov/dsp/parks/kauai/waimea-canyon-state-park',
            notes:
              '90-minute drive from East Side. Bring layers as it gets cooler at elevation.',
          },
          {
            id: '5-2',
            time: '2026-01-14T10:00:00',
            title: 'Waimea Canyon Lookout',
            location: 'Waimea Canyon State Park',
            category: 'hike',
            description:
              'Multiple viewpoints along the drive. Short walks to overlooks.',
            images: [
              'https://images.unsplash.com/photo-1542259009477-d625272157b7?w=400',
              'https://images.unsplash.com/photo-1580837119756-563d608dd119?w=400',
            ],
            rating: 4.8,
            reviewCount: 2234,
            address: 'Waimea Canyon Dr, Waimea, HI 96796',
            website:
              'https://dlnr.hawaii.gov/dsp/parks/kauai/waimea-canyon-state-park',
            notes:
              'Best lighting in morning or late afternoon. Several viewpoints accessible with kids.',
          },
          {
            id: '5-3',
            time: '2026-01-14T12:00:00',
            title: 'Lunch at Waimea Brewing Company',
            location: '9400 Kaumualii Hwy, Waimea',
            category: 'food',
            description: 'Good food with outdoor seating. Kid-friendly menu.',
            images: [
              'https://images.unsplash.com/photo-1513104890138-7c749659a591?w=400',
            ],
            rating: 4.3,
            reviewCount: 645,
            address: '9400 Kaumualii Hwy, Waimea, HI 96796',
            website: 'https://www.waimeabrewing.com',
            notes:
              'Try the Hawaiian pizza. Good selection of local beers for adults.',
          },
          {
            id: '5-4',
            time: '2026-01-14T14:30:00',
            title: 'Salt Pond Beach Park',
            location: 'Salt Pond Beach Park, Hanapepe',
            category: 'beach',
            description:
              'Calm, shallow water perfect for little ones. Natural tide pools to explore.',
            images: [
              'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=400',
            ],
            rating: 4.6,
            reviewCount: 723,
            address: 'Lokokai Rd, Hanapepe, HI 96716',
            website:
              'https://www.kauai.gov/Government/Departments-Agencies/Parks-Recreation',
            notes:
              'Protected swimming area. Lifeguard on duty. Historic salt ponds nearby.',
          },
          {
            id: '5-5',
            time: '2026-01-14T17:00:00',
            title: 'Hanapepe Art Walk',
            location: 'Hanapepe Town',
            category: 'shopping',
            description:
              "Friday night art walk in 'Kauai's Biggest Little Town'. Ice cream shops for the boys!",
            images: [
              'https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400',
            ],
            rating: 4.5,
            reviewCount: 456,
            address: 'Hanapepe Rd, Hanapepe, HI 96716',
            website: 'https://www.hanapepe.org',
            notes:
              'Only on Fridays 6-9 PM. Free parking. Street vendors and live music.',
          },
        ],
      },
      {
        date: '2026-01-15',
        activities: [
          {
            id: '6-1',
            time: '2026-01-15T09:00:00',
            title: 'Poipu Beach',
            location: 'Poipu Beach Park',
            category: 'beach',
            description:
              "Excellent snorkeling and often monk seals resting on the beach. Kids' favorite!",
            images: [
              'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400',
              'https://images.unsplash.com/photo-1505142468610-359e7d316be0?w=400',
            ],
            rating: 4.8,
            reviewCount: 1876,
            address: 'Hoowili Rd, Koloa, HI 96756',
            website:
              'https://www.kauai.gov/Government/Departments-Agencies/Parks-Recreation',
            notes:
              'Stay 50 feet away from monk seals. Great for snorkeling in protected cove. Showers available.',
          },
          {
            id: '6-2',
            time: '2026-01-15T12:00:00',
            title: 'Lunch at Brenneckes Beach Broiler',
            location: '2100 Hoone Rd, Koloa',
            category: 'food',
            description: 'Right across from the beach with great ocean views',
            images: [
              'https://images.unsplash.com/photo-1414235077428-338989a2e8c0?w=400',
            ],
            rating: 4.4,
            reviewCount: 1089,
            address: '2100 Hoone Rd, Koloa, HI 96756',
            website: 'https://www.brenneckes.com',
            notes:
              'Upstairs has best views. Famous for steaks and seafood. Keiki menu available.',
          },
          {
            id: '6-3',
            time: '2026-01-15T14:00:00',
            title: 'Spouting Horn',
            location: 'Spouting Horn Park, Poipu',
            category: 'activity',
            description:
              'Natural blowhole that shoots water 50 feet in the air. Kids love it!',
            images: [
              'https://images.unsplash.com/photo-1582407947304-fd86f028f716?w=400',
            ],
            rating: 4.5,
            reviewCount: 1567,
            address: 'Lawai Rd, Poipu, HI 96756',
            website:
              'https://www.gohawaii.com/islands/kauai/regions/south-shore/spouting-horn',
            notes:
              'Stay behind barriers. Best during high tide. Parking lot can fill up.',
          },
          {
            id: '6-4',
            time: '2026-01-15T16:00:00',
            title: 'Shops at Kukuiula',
            location: '2829 Ala Kalanikaumaka, Koloa',
            category: 'shopping',
            description: 'Upscale shopping village with playground for kids',
            images: [
              'https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=400',
            ],
            rating: 4.6,
            reviewCount: 534,
            address: '2829 Ala Kalanikaumaka St, Koloa, HI 96756',
            website: 'https://www.theshopsatkukuiula.com',
            notes:
              'Nice playground and lawn area. Weekly farmers market on Wednesdays.',
          },
          {
            id: '6-5',
            time: '2026-01-15T18:30:00',
            title: 'Farewell Dinner at Merriman Fish House',
            location: '2829 Ala Kalanikaumaka, Koloa',
            category: 'food',
            description:
              'Farm-to-table Hawaiian cuisine for a special last dinner',
            images: [
              'https://images.unsplash.com/photo-1550966871-3ed3cdb5ed0c?w=400',
              'https://images.unsplash.com/photo-1559339352-11d035aa65de?w=400',
            ],
            rating: 4.7,
            reviewCount: 945,
            address: '2829 Ala Kalanikaumaka St, Koloa, HI 96756',
            website: 'https://www.merrimanshawaii.com/merriman-s-fish-house',
            notes:
              'Reservations essential. Sunset views. Kids menu available but upscale atmosphere.',
          },
        ],
      },
      {
        date: '2026-01-16',
        activities: [
          {
            id: '7-1',
            time: '2026-01-16T08:00:00',
            title: 'Breakfast at Kalypso',
            location: '4-1384 Kuhio Hwy, Kapaa',
            category: 'food',
            description: 'Last breakfast with ocean views',
            images: [
              'https://images.unsplash.com/photo-1525351484163-7529414344d8?w=400',
            ],
            rating: 4.4,
            reviewCount: 578,
            address: '4-1384 Kuhio Hwy, Kapaa, HI 96746',
            website: 'https://www.kalypsokauai.com',
            notes: 'Oceanfront dining. Try the acai bowls. Relaxed atmosphere.',
          },
          {
            id: '7-2',
            time: '2026-01-16T10:00:00',
            title: 'Last beach time at Lydgate',
            location: 'Lydgate State Park',
            category: 'beach',
            description: 'Quick morning swim before heading to the airport',
            images: [],
            rating: 4.8,
            reviewCount: 1543,
            address: 'Leho Dr, Kapaa, HI 96746',
            website:
              'https://www.kauai.gov/Government/Departments-Agencies/Parks-Recreation',
            notes:
              'One last splash in the protected pools. Showers available before leaving.',
          },
          {
            id: '7-3',
            time: '2026-01-16T13:00:00',
            title: 'Return rental car & depart',
            location: 'Lihue Airport (LIH)',
            category: 'transport',
            description: 'Aloha, Kauai! Until next time.',
            images: [],
            rating: 4.2,
            reviewCount: 856,
            address: '3901 Mokulele Loop, Lihue, HI 96766',
            website: 'https://airports.hawaii.gov/lih',
            notes:
              'Allow extra time for car return. Small airport, arrive 2 hours before departure.',
          },
        ],
      },
    ],
  };

  getTripById(id: string): Trip | undefined {
    if (id === '1') {
      return this.mockTrip;
    }
    return undefined;
  }
}
