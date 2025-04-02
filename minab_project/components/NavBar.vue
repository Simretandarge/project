<template>
  <nav class="bg-white shadow-md px-8 py-4 fixed top-0 left-0 w-full z-50">
    <div class="flex items-center justify-between w-full gap-4">
      <!-- Left: Logo -->
      <NuxtLink
        to="/"
        class="flex items-center gap-2 text-xl font-bold text-green-600 hover:text-green-700"
        @click="refreshHomePage"
      >
        🍲<span class="truncate">Food Recipe</span>
      </NuxtLink>

      <!-- Search (desktop only) -->
      <div class="hidden md:flex items-center border border-gray-300 rounded-md px-3 py-1 bg-gray-50 flex-1 max-w-sm mx-4">
        <input
          type="text"
          placeholder="Search..."
          class="bg-transparent outline-none text-sm text-gray-700 w-full"
        />
        <button class="p-1 text-gray-600">
          <MagnifyingGlassIcon class="w-5 h-5" />
        </button>
      </div>

      <!-- Right: Nav or Toggle -->
      <div class="flex items-center gap-4">
        <button class="md:hidden" @click="toggleMenu" aria-label="Toggle Menu">
          <Bars3Icon class="w-6 h-6 text-gray-700 hover:text-green-600"/>
        </button>

        <!-- Desktop Nav Links -->
        <div class="hidden md:flex items-center gap-6">
          <ul class="flex items-center gap-8 text-gray-700 text-base">
            <li>
              <NuxtLink
                to="/"
                class="flex items-center gap-4 hover:text-green-600"
                @click="refreshHomePage"
              >
                <HomeIcon class="w-5 h-5" /> Home
              </NuxtLink>
            </li>
            <li>
              <NuxtLink
                to="/categories"
                class="flex items-center gap-4 hover:text-green-600"
              >
                <Squares2X2Icon class="w-5 h-5" /> Categories
              </NuxtLink>
            </li>
            <li>
              <NuxtLink
                to="/create"
                class="flex items-center gap-4 hover:text-green-600"
              >
                <PlusCircleIcon class="w-5 h-5" /> Create
              </NuxtLink>
            </li>
            <li>
              <NuxtLink
                to="/login"
                class="flex items-center gap-4 hover:text-green-600"
              >
                <ArrowRightOnRectangleIcon class="w-4 h-4" /> Login
              </NuxtLink>
            </li>
            <li>
              <NuxtLink
                to="/signup"
                class="flex items-center gap-4 hover:text-green-600"
              >
                <UserPlusIcon class="w-5 h-5" /> Signup
              </NuxtLink>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <!-- Mobile Menu Dropdown -->
    <transition name="slide-fade">
      <div v-if="menuOpen" class="md:hidden absolute left-0 top-full w-full bg-[#e2d4c1] shadow-lg border-t z-40 px-4 py-4 max-h-[80vh] overflow-y-auto rounded-b-xl">
        <div class="flex flex-col gap-4 mb-4 text-gray-800 text-sm">
          <NuxtLink
            to="/"
            class="px-3 py-2 rounded-md hover:bg-green-100"
            @click="refreshHomePage"
          >
            <HomeIcon class="w-5 h-5" /> Home
          </NuxtLink>
          <NuxtLink
            to="/categories"
            class="px-3 py-2 rounded-md hover:bg-green-100"
          >
            <Squares2X2Icon class="w-5 h-5" /> Categories
          </NuxtLink>
          <NuxtLink
            to="/create"
            class="px-3 py-2 rounded-md hover:bg-green-100"
          >
            <PlusCircleIcon class="w-5 h-5" /> Create
          </NuxtLink>
          <NuxtLink
            to="/login"
            class="px-3 py-2 rounded-md hover:bg-green-100"
          >
            <ArrowRightOnRectangleIcon class="w-5 h-5" /> Login
          </NuxtLink>
          <NuxtLink
            to="/signup"
            class="px-3 py-2 rounded-md hover:bg-green-100"
          >
            <UserPlusIcon class="w-5 h-5" /> Signup
          </NuxtLink>
        </div>
      </div>
    </transition>
  </nav>
</template>

<script setup>
import { ref } from "vue";
import {
  HomeIcon,
  Squares2X2Icon,
  PlusCircleIcon,
  ArrowRightOnRectangleIcon,
  UserPlusIcon,
  MagnifyingGlassIcon,
  Bars3Icon,
} from "@heroicons/vue/24/outline";

const menuOpen = ref(false);

function toggleMenu() {
  menuOpen.value = !menuOpen.value;
}

function refreshHomePage(event) {
  if (window.location.pathname !== "/") {
    window.location.href = "/";  // This forces a page reload when you click Home or Food Recipe.
  }
  event.preventDefault(); // Prevent the default behavior of the link and manually reload
}
</script>

<style scoped>
/* Slide transition for mobile menu */
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

/* Fixed navbar */
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


/* Hover effects for nav links */
a:hover {
  color: #4caf50; /* Green color */
}

/* Active link styling */
a.active {
  background-color: #e8f5e9; /* Light green */
}
</style>
