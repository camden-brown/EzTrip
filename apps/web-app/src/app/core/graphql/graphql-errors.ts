/**
 * Error codes from the backend API
 */
export enum GraphQLErrorCode {
  VALIDATION_ERROR = 'VALIDATION_ERROR',
  NOT_FOUND = 'NOT_FOUND',
  UNAUTHORIZED = 'UNAUTHORIZED',
  FORBIDDEN = 'FORBIDDEN',
  INTERNAL_ERROR = 'INTERNAL_ERROR',
  BAD_REQUEST = 'BAD_REQUEST',
}

/**
 * Extensions that can be attached to GraphQL errors
 */
export interface GraphQLErrorExtensions<
  ErrorCode extends GraphQLErrorCode = GraphQLErrorCode,
> {
  code?: ErrorCode;
  field?: string;
  [key: string]: unknown;
}

/**
 * GraphQL error structure from the API
 */
export interface GraphQLError {
  message: string;
  path?: Array<string | number>;
  extensions?: GraphQLErrorExtensions;
}

/**
 * GraphQL response structure
 */
export interface GraphQLResponse<TData> {
  data?: TData;
  errors?: GraphQLError[];
}

/**
 * Enhanced error class for GraphQL requests with structured error support
 * @template TErrorCode - Specific error codes expected for this operation (extends base ErrorCode)
 */
export class GraphQLRequestError<
  TErrorCode extends GraphQLErrorCode = GraphQLErrorCode,
> extends Error {
  public readonly errors: GraphQLError[];
  public readonly code?: TErrorCode;
  public readonly field?: string;

  constructor(message: string, errors: GraphQLError[]) {
    super(message);
    this.name = 'GraphQLRequestError';
    this.errors = errors;

    if (errors.length > 0 && errors[0].extensions) {
      this.code = errors[0].extensions.code as TErrorCode;
      this.field = errors[0].extensions.field;
    }
  }

  /**
   * Check if the error is of a specific type
   */
  isErrorCode(code: TErrorCode): boolean {
    return this.code === code;
  }

  /**
   * Check if the error is a validation error
   */
  isValidationError(): boolean {
    return this.code === GraphQLErrorCode.VALIDATION_ERROR;
  }

  /**
   * Check if the error is a not found error
   */
  isNotFoundError(): boolean {
    return this.code === GraphQLErrorCode.NOT_FOUND;
  }

  /**
   * Check if the error is an authorization error (unauthorized or forbidden)
   */
  isAuthError(): boolean {
    return (
      this.code === GraphQLErrorCode.UNAUTHORIZED ||
      this.code === GraphQLErrorCode.FORBIDDEN
    );
  }

  /**
   * Check if the error is related to a specific field
   */
  isFieldError(fieldName?: string): boolean {
    if (!this.field) {
      return false;
    }
    return fieldName ? this.field === fieldName : true;
  }

  /**
   * Get all errors for a specific field
   */
  getFieldErrors(fieldName: string): GraphQLError[] {
    return this.errors.filter((error) => error.extensions?.field === fieldName);
  }

  /**
   * Get a user-friendly error message
   */
  getUserMessage(): string {
    if (this.isAuthError()) {
      return (
        this.message || 'You do not have permission to perform this action'
      );
    }

    if (this.isNotFoundError()) {
      return this.message || 'The requested resource was not found';
    }

    if (this.isValidationError()) {
      return this.message || 'Please check your input and try again';
    }

    return this.message || 'An unexpected error occurred';
  }
}
