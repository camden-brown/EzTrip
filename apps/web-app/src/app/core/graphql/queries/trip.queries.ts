import { gql } from 'graphql-tag';

export interface TripSuggestionQuery {
  tripSuggestion: string;
}

export interface TripSuggestionVariables {
  prompt: string;
}

export const TRIP_SUGGESTION_QUERY = gql`
  query TripSuggestion($prompt: String!) {
    tripSuggestion(prompt: $prompt)
  }
`;
