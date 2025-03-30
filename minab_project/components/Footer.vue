<template>
    <footer class="bg-gray-100 text-gray-700 mt-12 px-6 py-8 rounded-t-xl shadow-inner">
      <div class="max-w-6xl mx-auto flex flex-col md:flex-row justify-between items-center gap-4">
        <p class="text-center md:text-left text-sm">
          © {{ new Date().getFullYear() }} Minab Recipe. All rights reserved.
        </p>
        <div class="flex gap-6 text-sm">
          <NuxtLink to="/" class="hover:text-green-600 transition">Home</NuxtLink>
          <NuxtLink to="/categories" class="hover:text-green-600 transition">Categories</NuxtLink>
          <NuxtLink to="/create" class="hover:text-green-600 transition">Create</NuxtLink>
          <NuxtLink to="/about" class="hover:text-green-600 transition">About</NuxtLink>
          <NuxtLink to="/contact" class="hover:text-green-600 transition">Contact</NuxtLink>
        </div>
      </div>
    </footer>
  </template>
  