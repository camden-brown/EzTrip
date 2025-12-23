import { InjectionToken } from '@angular/core';
import { environment } from '../../../environments/environment';

/**
 * Base URL for the GraphQL endpoint.
 *
 * Default assumes Angular dev-server proxy routes `/api-go` to the Go API.
 */
export const GRAPHQL_ENDPOINT = new InjectionToken<string>('GRAPHQL_ENDPOINT', {
  factory: () => environment.graphqlEndpoint,
});
