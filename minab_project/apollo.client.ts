// plugins/apollo-client.ts
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client/core";

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = createHttpLink({
    uri: "http://localhost:8080/v1/graphql", // Hasura endpoint
  });

  const apolloClient = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache(),
  });

  nuxtApp.provide("apollo", apolloClient);
});
