import { gql } from 'graphql-tag';

export const GET_TRIPS = gql`
  query GetTrips {
    trips {
      id
      title
      destination
      startDate
      endDate
      travelers
      itinerary {
        id
        date
        dayNumber
        activities {
          id
          time
          title
          location
          category
          type
          description
          notes
          placeId
        }
      }
    }
  }
`;

export const GET_TRIP = gql`
  query GetTrip($id: ID!) {
    trip(id: $id) {
      id
      title
      destination
      startDate
      endDate
      travelers
      itinerary {
        id
        date
        dayNumber
        activities {
          id
          time
          title
          location
          category
          type
          description
          notes
          placeId
        }
      }
    }
  }
`;

export const GET_ACTIVITY = gql`
  query GetActivity($id: ID!) {
    activity(id: $id) {
      id
      time
      title
      location
      category
      type
      description
      notes
      placeId
    }
  }
`;
