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
          },
          {
            id: '1-2',
            time: '2026-01-10T12:30:00',
            title: 'Lunch at Kalapaki Beach Hut',
            location: '3474 Rice St, Lihue',
            category: 'food',
            description:
              'Casual beachside spot with great fish tacos - perfect for the kids',
          },
          {
            id: '1-3',
            time: '2026-01-10T14:00:00',
            title: 'Kalapaki Beach',
            location: 'Kalapaki Beach, Lihue',
            category: 'beach',
            description:
              'Gentle waves perfect for the boys. Boogie boards available for rent.',
          },
          {
            id: '1-4',
            time: '2026-01-10T18:00:00',
            title: 'Dinner at Hukilau Lanai',
            location: '520 Aleka Loop, Kapaa',
            category: 'food',
            description: 'Family-friendly Hawaiian cuisine with live music',
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
          },
          {
            id: '2-2',
            time: '2026-01-11T09:30:00',
            title: 'Sleeping Giant Trail (Nounou East)',
            location: 'Nounou Trail East, Kapaa',
            category: 'hike',
            description:
              '2-mile moderate hike with stunning views. Doable for 7-year-old, might need to carry the 4-year-old for parts.',
          },
          {
            id: '2-3',
            time: '2026-01-11T13:00:00',
            title: 'Lunch at Opakapaka Grill',
            location: '4-1543 Kuhio Hwy, Kapaa',
            category: 'food',
            description: 'Fresh poke bowls and local plates',
          },
          {
            id: '2-4',
            time: '2026-01-11T15:00:00',
            title: 'Lydgate Beach Park',
            location: 'Lydgate State Park, Kapaa',
            category: 'beach',
            description:
              'Protected swimming area perfect for young kids. Has a great playground too!',
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
          },
          {
            id: '3-2',
            time: '2026-01-12T08:30:00',
            title: 'Breakfast at Hanalei Bread Company',
            location: '5-5161 Kuhio Hwy, Hanalei',
            category: 'food',
            description: 'Amazing pastries and breakfast sandwiches',
          },
          {
            id: '3-3',
            time: '2026-01-12T10:00:00',
            title: 'Hanalei Bay Beach',
            location: 'Hanalei Bay, Hanalei',
            category: 'beach',
            description:
              'One of the most beautiful beaches in Hawaii. Great for swimming and building sandcastles.',
          },
          {
            id: '3-4',
            time: '2026-01-12T12:30:00',
            title: 'Lunch at Tahiti Nui',
            location: '5-5134 Kuhio Hwy, Hanalei',
            category: 'food',
            description: 'Historic Hawaiian restaurant with great burgers',
          },
          {
            id: '3-5',
            time: '2026-01-12T14:30:00',
            title: "Queen's Bath (viewpoint only)",
            location: "Queen's Bath Trail, Princeville",
            category: 'hike',
            description:
              'Short walk to see the famous tidal pool. Too dangerous for swimming with kids, but beautiful views.',
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
              'Family-friendly kayak trip to Secret Falls. The 7-year-old can help paddle!',
          },
          {
            id: '4-2',
            time: '2026-01-13T13:00:00',
            title: 'Lunch at Wailua Shave Ice',
            location: '4-831 Kuhio Hwy, Kapaa',
            category: 'food',
            description:
              'Best shave ice on the island - the boys will love this!',
          },
          {
            id: '4-3',
            time: '2026-01-13T15:00:00',
            title: 'Smith Family Garden Luau',
            location: '3-5971 Kuhio Hwy, Lihue',
            category: 'entertainment',
            description:
              'Family-friendly luau with boat ride, gardens, and traditional show',
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
          },
          {
            id: '5-2',
            time: '2026-01-14T10:00:00',
            title: 'Waimea Canyon Lookout',
            location: 'Waimea Canyon State Park',
            category: 'hike',
            description:
              'Multiple viewpoints along the drive. Short walks to overlooks.',
          },
          {
            id: '5-3',
            time: '2026-01-14T12:00:00',
            title: 'Lunch at Waimea Brewing Company',
            location: '9400 Kaumualii Hwy, Waimea',
            category: 'food',
            description: 'Good food with outdoor seating. Kid-friendly menu.',
          },
          {
            id: '5-4',
            time: '2026-01-14T14:30:00',
            title: 'Salt Pond Beach Park',
            location: 'Salt Pond Beach Park, Hanapepe',
            category: 'beach',
            description:
              'Calm, shallow water perfect for little ones. Natural tide pools to explore.',
          },
          {
            id: '5-5',
            time: '2026-01-14T17:00:00',
            title: 'Hanapepe Art Walk',
            location: 'Hanapepe Town',
            category: 'shopping',
            description:
              "Friday night art walk in 'Kauai's Biggest Little Town'. Ice cream shops for the boys!",
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
          },
          {
            id: '6-2',
            time: '2026-01-15T12:00:00',
            title: 'Lunch at Brenneckes Beach Broiler',
            location: '2100 Hoone Rd, Koloa',
            category: 'food',
            description: 'Right across from the beach with great ocean views',
          },
          {
            id: '6-3',
            time: '2026-01-15T14:00:00',
            title: 'Spouting Horn',
            location: 'Spouting Horn Park, Poipu',
            category: 'activity',
            description:
              'Natural blowhole that shoots water 50 feet in the air. Kids love it!',
          },
          {
            id: '6-4',
            time: '2026-01-15T16:00:00',
            title: 'Shops at Kukuiula',
            location: '2829 Ala Kalanikaumaka, Koloa',
            category: 'shopping',
            description: 'Upscale shopping village with playground for kids',
          },
          {
            id: '6-5',
            time: '2026-01-15T18:30:00',
            title: 'Farewell Dinner at Merriman Fish House',
            location: '2829 Ala Kalanikaumaka, Koloa',
            category: 'food',
            description:
              'Farm-to-table Hawaiian cuisine for a special last dinner',
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
          },
          {
            id: '7-2',
            time: '2026-01-16T10:00:00',
            title: 'Last beach time at Lydgate',
            location: 'Lydgate State Park',
            category: 'beach',
            description: 'Quick morning swim before heading to the airport',
          },
          {
            id: '7-3',
            time: '2026-01-16T13:00:00',
            title: 'Return rental car & depart',
            location: 'Lihue Airport (LIH)',
            category: 'transport',
            description: 'Aloha, Kauai! Until next time.',
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
