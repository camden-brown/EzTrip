export const environment = {
  production: false,

  /**
   * GraphQL endpoint used by the frontend.
   *
   * In local development we rely on the Angular dev-server proxy:
   * `/api-go/graphql` -> `http://localhost:8080/graphql`
   */
  graphqlEndpoint: '/api-go/graphql',
};
