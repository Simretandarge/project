<template>
    <div class="max-w-4xl mx-auto p-6">
      <div v-if="loading">Loading recipes...</div>
      <div v-else-if="error">❌ Error: {{ error.message }}</div>
      <div v-else>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div
            v-for="recipe in recipes"
            :key="recipe.id"
            class="p-4 bg-white shadow rounded-xl transition hover:shadow-md"
          >
            <img
              :src="recipe.image_url"
              alt="Recipe Image"
              class="w-full h-48 object-cover rounded-md mb-3"
            />
            <h2 class="text-xl font-bold mb-1">{{ recipe.title }}</h2>
            <p class="text-gray-600 text-sm mb-2">{{ recipe.description }}</p>
            <p class="text-gray-500 text-xs">Prep Time: {{ recipe.prep_time }} mins</p>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { useQuery } from '@vue/apollo-composable';
  import { computed } from 'vue';
  import gql from 'graphql-tag'
  
  const GET_RECIPES = gql`
    query GetRecipes {
      recipes(order_by: { created_at: desc }) {
        id
        title
        description
        image_url
        prep_time
        created_at
      }
    }
  `;
  
  const { result, loading, error } = useQuery(GET_RECIPES);
  const recipes = computed(() => result.value?.recipes || []);
  </script>
  
  <style scoped>
  /* You can add custom styles here if needed */
  </style>
  