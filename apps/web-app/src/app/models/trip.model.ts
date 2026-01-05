export type ActivityCategory =
  | 'beach'
  | 'hike'
  | 'food'
  | 'hotel'
  | 'activity'
  | 'transport'
  | 'shopping'
  | 'entertainment';

export enum CategoryIcon {
  Beach = 'beach_access',
  Hike = 'terrain',
  Food = 'restaurant',
  Hotel = 'hotel',
  Activity = 'local_activity',
  Transport = 'flight',
  Shopping = 'shopping_bag',
  Entertainment = 'celebration',
  Default = 'place',
}

export enum CategoryLabel {
  Beach = 'Beach',
  Hike = 'Hiking',
  Food = 'Restaurant',
  Hotel = 'Hotel',
  Activity = 'Activity',
  Transport = 'Transportation',
  Shopping = 'Shopping',
  Entertainment = 'Entertainment',
  Default = 'Activity',
}

export interface Activity {
  id: string;
  time: string; // ISO 8601 datetime string
  title: string;
  location: string;
  category: ActivityCategory;
  description: string;
  images: string[];
  rating: number;
  reviewCount: number;
  address: string;
  website: string;
  notes: string;
}

export interface ItineraryDay {
  date: string;
  activities: Activity[];
}

export interface Trip {
  id: string;
  title: string;
  destination: string;
  startDate: string;
  endDate: string;
  travelers: number;
  itinerary: ItineraryDay[];
}
