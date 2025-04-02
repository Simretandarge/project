<template>
  <div class="max-w-4xl mx-auto mt-10 p-6 bg-white shadow-lg rounded-lg">
    <div v-if="loading" class="text-center text-gray-500">Loading...</div>
    <div v-else-if="!recipe" class="text-center text-red-500">
      Recipe not found.
    </div>
    <div v-else>
      <!-- Title and Category -->
      <h1
        class="text-3xl font-bold text-center text-green-700 mb-4 hover:text-green-600 transition duration-300"
      >
        {{ recipe.title }}
      </h1>
      <p class="text-sm text-green-600 italic mb-4 text-center">
        Category: <span class="text-green-800">{{ recipe.category }}</span>
      </p>

      <!-- Images -->
      <div
        v-if="recipe.images?.length"
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-6"
      >
        <img
          v-for="(img, index) in recipe.images"
          :key="index"
          :src="img"
          alt="Recipe image"
          class="rounded-lg shadow-lg hover:scale-105 transition-transform duration-300"
        />
      </div>

      <!-- Description -->
      <p class="text-gray-700 mb-6 text-lg leading-relaxed">
        {{ recipe.description }}
      </p>

      <!-- Ingredients -->
      <div v-if="recipe.ingredients?.length" class="mb-6">
        <h2 class="text-xl font-semibold mb-2 text-green-700">Ingredients</h2>
        <ul class="list-disc list-inside space-y-1 text-gray-600">
          <li
            v-for="(ing, idx) in recipe.ingredients"
            :key="idx"
            class="hover:text-green-600 transition duration-200"
          >
            {{ ing.name }} – <span class="font-medium">{{ ing.quantity }}</span>
          </li>
        </ul>
      </div>

      <!-- Steps -->
      <div v-if="recipe.steps?.length" class="mb-6">
        <h2 class="text-xl font-semibold mb-2 text-green-700">Steps</h2>
        <ol class="list-decimal list-inside space-y-2 text-gray-600">
          <li
            v-for="(step, idx) in sortedSteps"
            :key="idx"
            class="hover:text-green-600 transition duration-200"
          >
            {{ step.instruction }}
          </li>
        </ol>
      </div>

      <!-- Like, Bookmark, Rate -->
      <div class="flex justify-start items-center space-x-6 mb-6">
        <!-- Like Button -->
        <button
          @click="handleLike"
          class="flex items-center text-red-600 hover:text-red-700 transition duration-300 space-x-2"
        >
          <svg
            class="w-5 h-5"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M11.049 2.927C11.36 1.764 12.639 1.764 12.95 2.927l1.092 2.469a1.75 1.75 0 001.314 1.021l2.707.247c1.152.106 1.622 1.592.78 2.407l-2.042 1.968a1.75 1.75 0 00-.51 1.447l.55 2.634c.221 1.036-.843 1.881-1.756 1.392l-2.366-1.312a1.75 1.75 0 00-1.61 0l-2.366 1.312c-.913.489-1.977-.356-1.756-1.392l.55-2.634a1.75 1.75 0 00-.51-1.447L7.998 7.073c-.843-.815-.372-2.301.78-2.407l2.707-.247a1.75 1.75 0 001.314-1.021L11.049 2.927z"
            />
          </svg>
          <span>{{ recipe.likes.length }} Likes</span>
        </button>

        <!-- Bookmark Button -->
        <button
          @click="handleBookmark"
          class="flex items-center text-blue-600 hover:text-blue-700 transition duration-300 space-x-2"
        >
          <svg
            class="w-5 h-5"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 3a2 2 0 00-2 2v14l7-4 7 4V5a2 2 0 00-2-2H5z"
            />
          </svg>
          <span>{{ recipe.bookmarks.length }} Bookmarks</span>
        </button>

        <!-- Rating -->
        <div class="flex items-center space-x-2">
          <span>Average Rating:</span>
          <div class="flex space-x-1">
            <svg
              @click="setRating(1)"
              class="w-5 h-5"
              :class="rating >= 1 ? 'text-yellow-500' : 'text-gray-300'"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M10 15l-5.447 2.857L5.927 12 1 7.907l6.309-.55L10 2l2.691 5.357 6.309.55L14.073 12l.374 5.857L10 15z"
                clip-rule="evenodd"
              />
            </svg>
            <svg
              @click="setRating(2)"
              class="w-5 h-5"
              :class="rating >= 2 ? 'text-yellow-500' : 'text-gray-300'"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M10 15l-5.447 2.857L5.927 12 1 7.907l6.309-.55L10 2l2.691 5.357 6.309.55L14.073 12l.374 5.857L10 15z"
                clip-rule="evenodd"
              />
            </svg>
            <svg
              @click="setRating(3)"
              class="w-5 h-5"
              :class="rating >= 3 ? 'text-yellow-500' : 'text-gray-300'"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M10 15l-5.447 2.857L5.927 12 1 7.907l6.309-.55L10 2l2.691 5.357 6.309.55L14.073 12l.374 5.857L10 15z"
                clip-rule="evenodd"
              />
            </svg>
            <svg
              @click="setRating(4)"
              class="w-5 h-5"
              :class="rating >= 4 ? 'text-yellow-500' : 'text-gray-300'"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M10 15l-5.447 2.857L5.927 12 1 7.907l6.309-.55L10 2l2.691 5.357 6.309.55L14.073 12l.374 5.857L10 15z"
                clip-rule="evenodd"
              />
            </svg>
            <svg
              @click="setRating(5)"
              class="w-5 h-5"
              :class="rating >= 5 ? 'text-yellow-500' : 'text-gray-300'"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M10 15l-5.447 2.857L5.927 12 1 7.907l6.309-.55L10 2l2.691 5.357 6.309.55L14.073 12l.374 5.857L10 15z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
        </div>
      </div>

      <!-- Comments Section -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-700">Comments</h3>
        <textarea
          v-model="commentText"
          placeholder="Add a comment..."
          class="w-full p-3 border rounded-md shadow-sm text-gray-700 focus:ring-2 focus:ring-green-600 resize-none"
        ></textarea>
        <button
          @click="submitComment"
          class="mt-2 px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition"
        >
          Submit Comment
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRoute } from "vue-router";
import { useQuery, useMutation } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { computed, ref } from "vue";

// Get recipe ID from route
const route = useRoute();
const recipeId = route.params.id;

// Create a ref for the comment input
const commentText = ref("");
const rating = ref(0);

// GraphQL query to fetch single recipe
const GET_RECIPE = gql`
  query GetRecipeById($id: uuid!) {
    recipes_by_pk(id: $id) {
      id
      title
      description
      category
      images
      ingredients {
        name
        quantity
      }
      steps(order_by: { step_number: asc }) {
        step_number
        instruction
      }
      likes {
        user_id
      }
      bookmarks {
        user_id
      }
      ratings {
        user_id
        rating
      }
      comments {
        user_id
        comment
      }
    }
  }
`;

// GraphQL Mutation for Like
const LIKE_RECIPE = gql`
  mutation LikeRecipe($user_id: String!, $recipe_id: uuid!) {
    insert_likes_one(object: { user_id: $user_id, recipe_id: $recipe_id }) {
      id
    }
  }
`;

// GraphQL Mutation for Bookmark
const BOOKMARK_RECIPE = gql`
  mutation BookmarkRecipe($user_id: String!, $recipe_id: uuid!) {
    insert_bookmarks_one(object: { user_id: $user_id, recipe_id: $recipe_id }) {
      id
    }
  }
`;

// GraphQL Mutation for Rating
const RATE_RECIPE = gql`
  mutation RateRecipe($user_id: String!, $recipe_id: uuid!, $rating: Int!) {
    insert_ratings_one(
      object: { user_id: $user_id, recipe_id: $recipe_id, rating: $rating }
    ) {
      id
    }
  }
`;

// GraphQL Mutation for Comment
const SUBMIT_COMMENT = gql`
  mutation SubmitComment(
    $user_id: String!
    $recipe_id: uuid!
    $comment: String!
  ) {
    insert_comments_one(
      object: { user_id: $user_id, recipe_id: $recipe_id, comment: $comment }
    ) {
      id
    }
  }
`;

// Run the query
const { result, loading } = useQuery(GET_RECIPE, { id: recipeId });

// Get recipe and sort steps
const recipe = computed(() => result.value?.recipes_by_pk || null);
const sortedSteps = computed(() => recipe.value?.steps || []);

// Handle Like
const handleLike = async () => {
  const user_id = "example-user-id"; // Replace with actual logged-in user ID
  await useMutation(LIKE_RECIPE, { user_id, recipe_id: recipeId });
};

// Handle Bookmark
const handleBookmark = async () => {
  const user_id = "example-user-id"; // Replace with actual logged-in user ID
  await useMutation(BOOKMARK_RECIPE, { user_id, recipe_id: recipeId });
};

// Handle Rating
const setRating = async (value) => {
  rating.value = value;
  const user_id = "example-user-id"; // Replace with actual logged-in user ID
  await useMutation(RATE_RECIPE, {
    user_id,
    recipe_id: recipeId,
    rating: value,
  });
};

// Handle Comment Submission
const submitComment = async () => {
  const user_id = "example-user-id"; // Replace with actual logged-in user ID
  const comment = commentText.value;
  await useMutation(SUBMIT_COMMENT, { user_id, recipe_id: recipeId, comment });
};
</script>

<style scoped>
/* Custom Styles */
h1 {
  text-transform: capitalize;
}

button:hover {
  transform: scale(1.05);
}

button:active {
  transform: scale(1);
}

textarea {
  resize: vertical;
}

.star .star-rating {
  display: inline-block;
  color: gold;
}
</style>
