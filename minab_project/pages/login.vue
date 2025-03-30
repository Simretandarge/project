<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="bg-white p-8 rounded-2xl shadow-lg w-full max-w-md animate-fade-in-up">
        <h2 class="text-2xl font-bold mb-6 text-center">Login to Your Account</h2>
  
        <form @submit.prevent="handleLogin" class="space-y-4">
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input
              type="email"
              id="email"
              v-model="email"
              class="mt-1 block w-full px-4 py-2 border rounded-xl shadow-sm focus:ring focus:ring-blue-200 focus:outline-none"
              required
            />
          </div>
  
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
              type="password"
              id="password"
              v-model="password"
              class="mt-1 block w-full px-4 py-2 border rounded-xl shadow-sm focus:ring focus:ring-blue-200 focus:outline-none"
              required
            />
          </div>
  
          <button
            type="submit"
            class="w-full bg-blue-600 text-white py-2 rounded-xl hover:bg-blue-700 transition"
          >
            Login
          </button>
        </form>
  
        <p class="mt-4 text-center text-sm text-gray-600">
          Don't have an account?
          <NuxtLink to="/signup" class="text-blue-600 hover:underline">Sign up</NuxtLink>
        </p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const email = ref('')
  const password = ref('')
  
  const handleLogin = () => {
    // 🔒 You'll integrate Hasura/Backend login later here
    console.log('Logging in with:', email.value, password.value)
  }
  </script>
  