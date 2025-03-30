<template>
  <div>
    <!-- ✅ Hero Section -->
    <section
      class="relative bg-cover bg-center bg-no-repeat h-[90vh] flex items-center justify-center text-center text-white"
      style="background-image: url('/images/hero.jpg');"
    >
      <div class="absolute inset-0 bg-black/40 z-0"></div>
      <div class="relative z-10 px-4">
        <h1 class="text-4xl md:text-6xl font-extrabold mb-4 leading-tight">
          🍲 <span class="text-green-300 animate-pulse">Welcome to Minab Recipe</span>
        </h1>
        <p class="text-xl md:text-2xl font-medium mb-6">
          <span class="text-white">Discover,</span>
          <span class="text-green-200 hover:text-green-300 transition duration-300">share,</span>
          <span class="text-white">and enjoy recipes crafted by food lovers</span>
        </p>
      </div>
    </section>

    <!-- ✅ Recent Recipes Section -->
    <section class="py-12 px-6 max-w-7xl mx-auto">
      <h2 class="text-3xl font-bold text-center mb-10 text-green-700">Recent Recipes</h2>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
        <NuxtLink
          v-for="recipe in sampleRecipes"
          :key="recipe.id"
          :to="`/recipes/${recipe.id}`"
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
            <p class="text-xs text-gray-400">👨‍🍳 {{ recipe.chef }} • 🕒 {{ recipe.prepTime }} min</p>
          </div>
        </NuxtLink>
      </div>
    </section>

    <!-- ✅ Create Recipe CTA Section -->
    <section class="bg-white py-16 text-center">
      <h2 class="text-3xl font-bold mb-4 text-gray-800">Start Sharing Deliciousness</h2>
      <p class="text-gray-600 max-w-xl mx-auto mb-8">
        Got a recipe? Inspire others by adding your unique touch to the food world.
      </p>
      <NuxtLink to="/create" class="bg-green-600 hover:bg-green-700 text-white px-6 py-3 rounded-full font-semibold transition">
        ➕ Create Your Own Recipe
      </NuxtLink>
    </section>

    <!-- ✅ Categories Section -->
    <section class="py-12 px-6 max-w-7xl mx-auto">
      <NuxtLink to="/categories">
        <h2 class="text-3xl font-bold text-center mb-10 text-green-700 hover:underline cursor-pointer">
          Browse by Categories
        </h2>
      </NuxtLink>
      <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-6">
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
          <span class="text-sm font-semibold text-gray-700 group-hover:text-green-600">{{ category.name }}</span>
        </NuxtLink>
      </div>
    </section>
  </div>
</template>

<script setup>
const sampleRecipes = [
  {
    id: 1,
    title: 'Spaghetti Bolognese',
    chef: 'Chef Mario',
    description: 'A classic Italian pasta dish with rich meat sauce.',
    image: '../images/featured-2.jpg',
    prepTime: 25,
  },
  {
    id: 2,
    title: 'Avocado Toast',
    chef: 'Chef Lila',
    description: 'Simple and healthy breakfast with avocado and egg.',
    image: '../images/featured-1.jpg',
    prepTime: 10,
  },
  {
    id: 3,
    title: 'Berry Smoothie Bowl',
    chef: 'Chef Maya',
    description: 'Refreshing smoothie topped with fresh fruits.',
    image: '../images/featured-3.jpg',
    prepTime: 15,
  },
  {
    id: 1,
    title: 'Spaghetti Bolognese',
    chef: 'Chef Mario',
    description: 'A classic Italian pasta dish with rich meat sauce.',
    image: '../images/featured-4.jpg',
    prepTime: 25,
  }
];

const categories = [
  { name: 'Breakfast', slug: 'breakfast', image: '/images/categories/breakfast.jpg' },
  { name: 'Lunch', slug: 'lunch', image: '/images/categories/lunch.jpg' },
  { name: 'Dinner', slug: 'dinner', image: '/images/categories/dinner.jpg' },
  { name: 'Dessert', slug: 'dessert', image: '/images/categories/dessert.jpg' },
  { name: 'Vegan', slug: 'vegan', image: '/images/categories/vegan.jpg' },
  { name: 'Drinks', slug: 'drinks', image: '/images/categories/drinks.jpg' },
  { name: 'Traditional', slug: 'traditional', image: '/images/categories/traditional.jpg' },
  { name: 'Snacks', slug: 'snacks', image: '/images/categories/snacks.jpg' },
  { name: 'Soup', slug: 'soup', image: '/images/categories/soup.jpg' },
  { name: 'BBQ', slug: 'bbq', image: '/images/categories/bbq.jpg' },
  { name: 'Seafood', slug: 'seafood', image: '/images/categories/seafood.jpg' },
  { name: 'Salads', slug: 'salads', image: '/images/categories/salads.jpg' },
];
</script>

<style scoped>
/* Add your custom styles here if needed */
</style>
