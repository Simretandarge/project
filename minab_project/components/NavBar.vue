<template>
    <nav class="bg-white shadow-md px-4 py-4 relative z-50">
      <!-- Top Navbar -->
      <div class="flex items-center justify-between w-full gap-6">
        <!-- Left: Logo -->
        <NuxtLink
          to="/"
          class="flex items-center gap-4 text-xl font-bold text-green-600 hover:text-green-700 transition whitespace-nowrap"
        >
          <MagnifyingGlassIcon class="w-6 h-6 text-green-600 shrink-0" />
          <span class="truncate">Minab Recipe</span>
        </NuxtLink>
  
        <!-- Center: Search (desktop only) -->
        <div
          class="hidden md:flex items-center border border-gray-300 rounded-md px-3 py-2 bg-gray-50 flex-1 max-w-sm mx-4"
        >
          <MagnifyingGlassIcon class="w-6 h-6 text-gray-500 mr-2" />
          <input
            type="text"
            placeholder="Search..."
            class="bg-transparent outline-none text-sm text-gray-700 w-full"
          />
        </div>
  
        <!-- Right: Nav or Toggle -->
        <div class="flex items-center gap-4">
          <!-- Mobile toggle -->
          <button class="md:hidden" @click="toggleMenu" aria-label="Toggle Menu">
            <Bars3Icon class="w-8 h-8 text-gray-700" />
          </button>
  
          <!-- Desktop Nav Links -->
          <div class="hidden md:block max-w-[60vw] overflow-x-auto whitespace-nowrap">
            <ul class="flex items-center gap-6 text-gray-700 font-medium text-base">
              <li>
                <NuxtLink to="/" class="flex items-center gap-2 hover:text-green-600 transition">
                  <HomeIcon class="w-6 h-6" /> Home
                </NuxtLink>
              </li>
              <li>
                <NuxtLink to="/categories" class="flex items-center gap-2 hover:text-green-600 transition">
                  <Squares2X2Icon class="w-6 h-6" /> Categories
                </NuxtLink>
              </li>
              <li>
                <NuxtLink to="/create" class="flex items-center gap-2 hover:text-green-600 transition">
                  <PlusCircleIcon class="w-6 h-6" /> Create
                </NuxtLink>
              </li>
              <li>
                <NuxtLink to="/login" class="flex items-center gap-2 hover:text-green-600 transition">
                  <ArrowRightOnRectangleIcon class="w-5 h-5" /> Login
                </NuxtLink>
              </li>
              <li>
                <NuxtLink to="/signup" class="flex items-center gap-2 hover:text-green-600 transition">
                  <UserPlusIcon class="w-6 h-6" /> Signup
                </NuxtLink>
              </li>
            </ul>
          </div>
        </div>
      </div>
  
      <!-- Mobile Menu Dropdown -->
      <transition name="slide-fade">
  <div
    v-if="menuOpen"
    class="md:hidden absolute left-0 top-full w-full bg-[#e2d4c1] shadow-lg border-t border-green-200 z-40 px-4 py-4 max-h-[80vh] overflow-y-auto transition-all duration-300 ease-in-out rounded-b-xl"
  >
    <!-- Vertical Nav Links with hover -->
    <div class="flex flex-col gap-4 mb-4 text-gray-800 font-medium text-base">
      <NuxtLink
        to="/"
        class="flex items-center gap-2 px-3 py-2 rounded-md hover:bg-green-100 transition"
      >
        <HomeIcon class="w-5 h-5" /> Home
      </NuxtLink>
      <NuxtLink
        to="/categories"
        class="flex items-center gap-2 px-3 py-2 rounded-md hover:bg-green-100 transition"
      >
        <Squares2X2Icon class="w-5 h-5" /> Categories
      </NuxtLink>
      <NuxtLink
        to="/create"
        class="flex items-center gap-2 px-3 py-2 rounded-md hover:bg-green-100 transition"
      >
        <PlusCircleIcon class="w-5 h-5" /> Create
      </NuxtLink>
      <NuxtLink
        to="/login"
        class="flex items-center gap-2 px-3 py-2 rounded-md hover:bg-green-100 transition"
      >
        <ArrowRightOnRectangleIcon class="w-5 h-5" /> Login
      </NuxtLink>
      <NuxtLink
        to="/signup"
        class="flex items-center gap-2 px-3 py-2 rounded-md hover:bg-green-100 transition"
      >
        <UserPlusIcon class="w-5 h-5" /> Signup
      </NuxtLink>
    </div>

    <!-- Mobile Search -->
    <div class="flex items-center border border-blue-300 rounded-md px-3 py-2 bg-white w-full shadow-inner">
      <MagnifyingGlassIcon class="w-5 h-5 text-gray-500 mr-2" />
      <input
        type="text"
        placeholder="Search..."
        class="bg-transparent outline-none text-sm text-gray-700 w-full"
      />
    </div>
  </div>
</transition>

    </nav>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import {
    HomeIcon,
    Squares2X2Icon,
    PlusCircleIcon,
    ArrowRightOnRectangleIcon,
    UserPlusIcon,
    MagnifyingGlassIcon,
    Bars3Icon,
  } from '@heroicons/vue/24/outline'
  
  const menuOpen = ref(false)
  
  function toggleMenu() {
    menuOpen.value = !menuOpen.value
  }
  </script>
  
  <style>
  .slide-fade-enter-active,
  .slide-fade-leave-active {
    transition: all 0.3s ease;
  }
  .slide-fade-enter-from {
    opacity: 0;
    transform: translateY(-10px);
  }
  .slide-fade-leave-to {
    opacity: 0;
    transform: translateY(-10px);
  }
  
  /* Optional: Hide horizontal scrollbar if nav links overflow */
  ::-webkit-scrollbar {
    display: none;
  }
  </style>
  