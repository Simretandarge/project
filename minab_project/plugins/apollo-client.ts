import { defineNuxtPlugin } from '#app'
import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core'
import { DefaultApolloClient } from '@vue/apollo-composable'

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = new HttpLink({
    uri: 'http://localhost:8080/v1/graphql',
    headers: {
      'x-hasura-admin-secret': 'mysecretadmin'
    }
  })

  const apolloClient = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache()
  })

  nuxtApp.vueApp.provide(DefaultApolloClient, apolloClient)
})
