export const environment = {
  production: false,

  /**
   * GraphQL endpoint used by the frontend.
   *
   * In local development we rely on the Angular dev-server proxy:
   * `/api-go/graphql` -> `http://localhost:8080/graphql`
   */
  graphqlEndpoint: '/api-go/graphql',

  /**
   * Auth0 configuration
   * These values should be set via environment variables in production
   */
  auth0: {
    domain: 'eztrip.us.auth0.com',
    clientId: '8gQ8nnbeHrzzLgMLER5I81tUlJsDY5E9',
    authorizationParams: {
      redirect_uri: 'http://localhost:4200',
      audience: 'https://eztrip.us.auth0.com/api/v2/',
    },
    httpInterceptor: {
      allowedList: ['/api-go/*'],
    },
  },
};
