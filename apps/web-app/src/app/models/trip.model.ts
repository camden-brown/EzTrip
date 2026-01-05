export type ActivityCategory =
  | 'beach'
  | 'hike'
  | 'food'
  | 'hotel'
  | 'activity'
  | 'transport'
  | 'shopping'
  | 'entertainment';

export interface Activity {
  id: string;
  time: string; // ISO 8601 datetime string
  title: string;
  location: string;
  category: ActivityCategory;
  description?: string;
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
