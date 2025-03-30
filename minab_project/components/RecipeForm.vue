<template>
  <div class="max-w-2xl mx-auto p-6">
    <h1 class="text-2xl font-bold mb-4 text-green-700">🍳 Create a New Recipe</h1>

    <form @submit.prevent="submitForm" class="space-y-4">
      <input
        v-model="form.title"
        type="text"
        placeholder="Recipe Title"
        class="w-full border border-gray-300 rounded px-4 py-2"
        required
      />

      <textarea
        v-model="form.description"
        placeholder="Recipe Description"
        class="w-full border border-gray-300 rounded px-4 py-2"
        required
      ></textarea>

      <input
        v-model.number="form.prep_time"
        type="number"
        placeholder="Preparation Time (mins)"
        class="w-full border border-gray-300 rounded px-4 py-2"
        required
      />

      <input
        v-model="form.image_url"
        type="text"
        placeholder="Image URL"
        class="w-full border border-gray-300 rounded px-4 py-2"
      />

      <button
        type="submit"
        class="bg-green-600 hover:bg-green-700 text-white px-6 py-2 rounded-full font-semibold transition"
        :disabled="loading"
      >
        {{ loading ? 'Submitting...' : 'Submit Recipe' }}
      </button>

      <p v-if="successMessage" class="text-green-600 font-semibold mt-2">✅ {{ successMessage }}</p>
      <p v-if="errorMessage" class="text-red-600 font-semibold mt-2">❌ {{ errorMessage }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import gql from 'graphql-tag'
import { useMutation } from '@vue/apollo-composable'

const form = ref({
  title: '',
  description: '',
  prep_time: null,
  image_url: '',
})

const successMessage = ref('')
const errorMessage = ref('')

const CREATE_RECIPE = gql`
  mutation InsertRecipe($title: String!, $description: String!, $prep_time: Int!, $image_url: String) {
    insert_recipes_one(object: {
      title: $title,
      description: $description,
      prep_time: $prep_time,
      image_url: $image_url
    }) {
      id
    }
  }
`

const { mutate, loading } = useMutation(CREATE_RECIPE)

const submitForm = async () => {
  successMessage.value = ''
  errorMessage.value = ''

  try {
    await mutate({
      title: form.value.title,
      description: form.value.description,
      prep_time: form.value.prep_time,
      image_url: form.value.image_url || null,
    })
    successMessage.value = 'Recipe submitted successfully!'
    form.value = { title: '', description: '', prep_time: null, image_url: '' }
  } catch (err) {
    errorMessage.value = 'Failed to submit recipe.'
    console.error(err)
  }
}
</script>

<style scoped>
input, textarea {
  transition: all 0.3s ease;
}
input:focus, textarea:focus {
  border-color: #34d399;
  outline: none;
  box-shadow: 0 0 0 2px rgba(34, 197, 94, 0.2);
}
</style>