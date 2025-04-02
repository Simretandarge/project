<template>
  <div>
    <!-- ✅ Hero Section -->
    <section
      class="relative bg-cover bg-center bg-no-repeat h-[90vh] flex items-center justify-center text-center text-white"
      style="background-image: url('/images/hero.jpg')"
    >
      <div class="absolute inset-0 bg-black/40 z-0"></div>
      <div class="relative z-10 px-4">
        <h1 class="text-4xl md:text-6xl font-extrabold mb-4 leading-tight">
          🍲
          <span class="text-green-300 animate-pulse"
            >Welcome to Minab Recipe</span
          >
        </h1>
        <p class="text-xl md:text-2xl font-medium mb-6">
          <span class="text-white">Discover,</span>
          <span
            class="text-green-200 hover:text-green-300 transition duration-300"
            >share,</span
          >
          <span class="text-white"
            >and enjoy recipes crafted by food lovers</span>
        </p>
      </div>
    </section>

    <!-- ✅ Recent Recipes Section -->
    <section class="py-12 px-6 max-w-7xl mx-auto">
      <h2 class="text-3xl font-bold text-center mb-10 text-green-700">
        Recent Recipes
      </h2>

      <!-- Display loading state -->
      <div v-if="loading" class="text-center text-gray-500">
        Loading Recipes...
      </div>

      <!-- Display error message if no recipes found -->
      <div v-else-if="error" class="text-center text-red-500">
        Error loading recipes. Please try again later.
      </div>

      <!-- Display recipes -->
      <div
        v-else
        class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6"
      >
        <div
          v-for="recipe in recipes"
          :key="recipe.id"
          class="block rounded-xl shadow hover:shadow-lg transition bg-white overflow-hidden"
        >
          <img
            :src="recipe.image"
            alt="Recipe Image"
            class="w-full h-48 object-cover"
          />
          <div class="p-4">
            <h3 class="text-lg font-bold text-gray-800">{{ recipe.title }}</h3>
            <p class="text-gray-500 text-sm mb-2">{{ recipe.description }}</p>
            <p class="text-xs text-gray-400">
              👨‍🍳 Category: {{ recipe.category }}
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- ✅ Categories Section -->
    <section class="py-12 px-6 max-w-7xl mx-auto">
      <NuxtLink to="/categories">
        <h2
          class="text-3xl font-bold text-center mb-10 text-green-700 hover:underline cursor-pointer"
        >
          Browse by Categories
        </h2>
      </NuxtLink>
      <div
        class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-6"
      >
        <NuxtLink
          v-for="category in categories"
          :key="category.name"
          :to="`/categories/${category.slug}`"
          class="flex flex-col items-center justify-center p-4 bg-white rounded-xl shadow-md hover:shadow-lg transition duration-300 text-center group"
        >
          <img
            :src="category.image"
            :alt="category.name"
            class="w-16 h-16 object-cover rounded-full mb-3 border-2 border-green-200 group-hover:scale-105 transition"
          />
          <span
            class="text-sm font-semibold text-gray-700 group-hover:text-green-600"
            >{{ category.name }}</span
          >
        </NuxtLink>
      </div>
    </section>
  </div>
</template>

<script setup>
import { useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";

// GraphQL query to fetch recipes from the backend
const GET_RECIPES = gql`
  query GetRecipes {
    recipes {
      id
      title
      description
      category
      images
    }
  }
`;

// Fetch recipes using Apollo Client
const { result, loading, error } = useQuery(GET_RECIPES);

// Cache recipes data in localStorage for persistence (only on client side)
const recipes = computed(() => {
  if (process.client) {
    if (result.value?.recipes) {
      // Store the recipes in localStorage only on the client-side
      localStorage.setItem('recentRecipes', JSON.stringify(result.value.recipes));
      return result.value.recipes;
    }

    // Retrieve from localStorage if no data fetched
    const cachedRecipes = localStorage.getItem('recentRecipes');
    return cachedRecipes ? JSON.parse(cachedRecipes) : [];
  }
  return [];
});

const categories = [
  {
    name: "Breakfast",
    slug: "breakfast",
    image: "../images/breakfast.jpg",
  },
  { name: "Lunch", slug: "lunch", image: "../images/lunch.jpg" },
  { name: "Dinner", slug: "dinner", image: "../images/lunch.jpg" },
  { name: "Dessert", slug: "dessert", image: "../images/featured-4.jpg" },
  { name: "Vegan", slug: "vegan", image: "../images/vegan.jpg" },
  { name: "Drinks", slug: "drinks", image: "../images/drinks.jpg" },
  {
    name: "Traditional",
    slug: "traditional",
    image: "../images/featured-1.jpg",
  },
  { name: "Snacks", slug: "snacks", image: "../images/bbq.jpg" },
  { name: "Soup", slug: "soup", image: "../images/soup.jpg" },
  { name: "BBQ", slug: "bbq", image: "../images/bbq.jpg" },
  { name: "Seafood", slug: "seafood", image: "../images/vegan.jpg" },
  { name: "Salads", slug: "salads", image: "../images/vegan.jpg" },
];
</script>

<style scoped>
/* Custom styles if needed */
</style>
