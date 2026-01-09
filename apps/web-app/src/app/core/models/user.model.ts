export interface User {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
}

export interface SignupCredentials {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}

/**
 * User-specific error codes
 */
export enum UserErrorCode {
  DUPLICATE_EMAIL = 'USER_DUPLICATE_EMAIL',
}

/**
 * User-related messages
 */
export enum UserErrorMessage {
  DUPLICATE_EMAIL = 'An account with this email already exists.',
  GENERIC_ERROR = 'An error occurred during signup. Please try again later.',
}

/**
 * Gets the appropriate user error message based on error code
 */
export function getUserErrorMessage(errorCode?: string): string {
  switch (errorCode) {
    case UserErrorCode.DUPLICATE_EMAIL:
      return UserErrorMessage.DUPLICATE_EMAIL;
    default:
      return UserErrorMessage.GENERIC_ERROR;
  }
}
