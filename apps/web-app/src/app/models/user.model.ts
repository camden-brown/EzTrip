export interface User {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
}

/**
 * User-specific error codes
 */
export enum UserErrorCode {
  DUPLICATE_EMAIL = 'USER_DUPLICATE_EMAIL',
}
