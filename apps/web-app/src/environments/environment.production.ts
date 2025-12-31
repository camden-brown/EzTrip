export const environment = {
  production: true,

  /**
   * Production should typically point at the same origin (reverse-proxied)
   * or a full URL if hosted separately.
   */
  graphqlEndpoint: '/graphql',

  /**
   * Auth0 configuration
   */
  auth0: {
    domain: 'eztrip.us.auth0.com',
    clientId: '8gQ8nnbeHrzzLgMLER5I81tUlJsDY5E9',
    authorizationParams: {
      redirect_uri: 'https://eztrip.com',
      audience: 'https://eztrip.us.auth0.com/api/v2/',
    },
    httpInterceptor: {
      allowedList: ['/graphql', '/api/*'],
    },
  },
};
