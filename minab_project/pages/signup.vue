<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="bg-white p-8 rounded-2xl shadow-lg w-full max-w-md animate-fade-in-up">
        <h2 class="text-2xl font-bold mb-6 text-center">Create an Account</h2>
  
        <form @submit.prevent="handleSignup" class="space-y-4">
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700">Full Name</label>
            <input
              type="text"
              id="name"
              v-model="name"
              class="mt-1 block w-full px-4 py-2 border rounded-xl shadow-sm focus:ring focus:ring-blue-200 focus:outline-none"
              required
            />
          </div>
  
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
            class="w-full bg-green-600 text-white py-2 rounded-xl hover:bg-green-700 transition"
          >
            Sign Up
          </button>
        </form>
  
        <p class="mt-4 text-center text-sm text-gray-600">
          Already have an account?
          <NuxtLink to="/login" class="text-blue-600 hover:underline">Login</NuxtLink>
        </p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const name = ref('')
  const email = ref('')
  const password = ref('')
  
  const handleSignup = () => {
    // ⚙️ Will connect to Hasura registration logic
    console.log('Signing up:', name.value, email.value, password.value)
  }
  </script>
  