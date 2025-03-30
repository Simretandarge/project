<template>
    <div class="max-w-6xl mx-auto py-12 px-4">
      <h1 class="text-4xl font-extrabold text-center text-green-700 mb-10">
        🍽️ Explore Recipe Categories
      </h1>
  
      <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-6">
        <NuxtLink
          v-for="category in categories"
          :key="category.slug"
          :to="`/categories/${category.slug}`"
          class="bg-white shadow-md hover:shadow-xl border border-gray-200 hover:border-green-500 rounded-xl p-6 flex flex-col items-center justify-center text-center transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-105"
        >
          <div class="text-3xl mb-3">
            {{ category.emoji }}
          </div>
          <div class="text-lg font-semibold text-gray-800">
            {{ category.name }}
          </div>
        </NuxtLink>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  const categories = [
    { name: 'Breakfast', slug: 'breakfast', emoji: '🍳' },
    { name: 'Lunch', slug: 'lunch', emoji: '🥪' },
    { name: 'Dinner', slug: 'dinner', emoji: '🍝' },
    { name: 'Desserts', slug: 'desserts', emoji: '🍰' },
    { name: 'Snacks', slug: 'snacks', emoji: '🍿' },
    { name: 'Drinks', slug: 'drinks', emoji: '🥤' },
    { name: 'Healthy', slug: 'healthy', emoji: '🥗' },
    { name: 'Fast Food', slug: 'fast-food', emoji: '🍔' },
    { name: 'Vegetarian', slug: 'vegetarian', emoji: '🥦' },
    { name: 'Soups', slug: 'soups', emoji: '🥣' },
    { name: 'Salads', slug: 'salads', emoji: '🥬' },
    { name: 'Grilled', slug: 'grilled', emoji: '🔥' }
  ]
  </script>
  