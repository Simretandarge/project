<template>
  <div class="max-w-6xl mx-auto mt-10 px-4">
    <h1 class="text-3xl font-extrabold mb-6 text-gray-800">
      🍽️ Browse by Category
    </h1>

    <!-- Category Filter Buttons -->
    <div class="flex flex-wrap gap-3 mb-8">
      <button
        v-for="cat in categories"
        :key="cat.category"
        @click="selectedCategory = cat.category"
        class="px-4 py-2 rounded-full bg-blue-100 text-sm text-blue-700 hover:bg-blue-300 transition"
      >
        {{ cat.category }}
      </button>
      <button
        v-if="selectedCategory"
        @click="selectedCategory = null"
        class="px-4 py-2 rounded-full bg-red-100 text-sm text-red-700 hover:bg-red-300 transition"
      >
        Clear Filter
      </button>
    </div>

    <!-- Recipe Grid -->
    <div v-if="loading" class="text-center text-lg text-gray-600">
      Loading...
    </div>
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
      <div
        v-for="recipe in filteredRecipes"
        :key="recipe.id"
        @click="goToRecipe(recipe.id)"
        class="cursor-pointer p-5 rounded-lg shadow hover:shadow-lg bg-white hover:bg-gray-50 transition-all duration-300 border"
      >
        <div class="mb-2">
          <h3 class="text-lg font-semibold text-green-700 truncate">
            {{ recipe.title }}
          </h3>
          <p class="text-sm text-gray-600 truncate">
            {{ recipe.description }}
          </p>
        </div>
        <span
          class="inline-block mt-2 px-3 py-1 bg-green-100 text-green-700 text-xs rounded-full"
        >
          Category: {{ recipe.category }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { useRouter } from "vue-router";

const router = useRouter();
const selectedCategory = ref(null);

// Queries
const GET_RECIPES = gql`
  query {
    recipes {
      id
      title
      description
      category
    }
  }
`;

const GET_CATEGORIES = gql`
  query {
    recipes(distinct_on: category) {
      category
    }
  }
`;

const { result: recipeResult, loading } = useQuery(GET_RECIPES);
const { result: categoryResult } = useQuery(GET_CATEGORIES);

const recipes = computed(() => recipeResult.value?.recipes || []);
const categories = computed(() => categoryResult.value?.recipes || []);

const filteredRecipes = computed(() =>
  selectedCategory.value
    ? recipes.value.filter((r) => r.category === selectedCategory.value)
    : recipes.value
);

const goToRecipe = (id) => {
  router.push(`/recipe/${id}`);
};
</script>
<style>
/* Fixed navbar */
nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 50;
  width: 100%;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding-top: 0.75rem; /* Reduced navbar height */
  padding-bottom: 0.75rem;
}

/* Add space for the body or content to be pushed below the navbar */
body {
  padding-top: 80px; /* Adjust this to match your navbar height */
}

/* Main content area, if needed */
main {
  padding-top: 80px; /* Adjust this as needed */
}

</style>
