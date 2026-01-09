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
    clientId: 'zpnPxpqpJXv6LEgpoEEuU7gCPHuzVhwd',
    authorizationParams: {
      redirect_uri: 'https://ez-trip.ai',
      audience: 'https://api.ez-trip.ai',
    },
    httpInterceptor: {
      allowedList: ['/graphql', '/api/*'],
    },
  },
};
