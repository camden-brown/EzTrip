import { gql } from 'graphql-tag';
import { User } from '../../../models/user.model';

export interface CurrentUserQuery {
  currentUser: User | null;
}

export const CURRENT_USER_QUERY = gql`
  query CurrentUser {
    currentUser {
      id
      firstName
      lastName
      email
    }
  }
`;
