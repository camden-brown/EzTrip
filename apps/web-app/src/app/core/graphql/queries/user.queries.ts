import { gql } from 'graphql-tag';
import { User, SignupCredentials } from '../../models/user.model';

export interface CurrentUserQuery {
  currentUser: User | null;
}

export interface CreateUserMutation {
  createUser: User;
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

export const CREATE_USER_MUTATION = gql`
  mutation CreateUser($input: CreateUserInput!) {
    createUser(input: $input) {
      id
      firstName
      lastName
      email
    }
  }
`;
