<template>
  <div class="max-w-3xl mx-auto p-6">
    <h1 class="text-3xl font-bold mb-6 text-green-700">Create a New Recipe</h1>

    <form @submit.prevent="handleSubmit" class="space-y-4">
      <input v-model="form.title" type="text" placeholder="Title" class="form-input" required />

      <textarea v-model="form.description" placeholder="Description" class="form-textarea" required></textarea>

      <input v-model="form.image_url" type="text" placeholder="Image URL" class="form-input" required />

      <input v-model.number="form.prep_time" type="number" placeholder="Prep Time (in mins)" class="form-input" required />

      <input v-model="form.chef" type="text" placeholder="Chef Name" class="form-input" required />

      <div>
        <label class="font-semibold">Ingredients:</label>
        <div v-for="(ingredient, index) in form.ingredients" :key="index" class="flex gap-2 mb-2">
          <input v-model="form.ingredients[index]" type="text" placeholder="Ingredient" class="form-input flex-1" />
          <button type="button" @click="form.ingredients.splice(index, 1)" class="text-red-500">Remove</button>
        </div>
        <button type="button" @click="form.ingredients.push('')" class="text-sm text-green-600 hover:underline">+ Add Ingredient</button>
      </div>

      <div>
        <label class="font-semibold">Steps:</label>
        <div v-for="(step, index) in form.steps" :key="index" class="flex gap-2 mb-2">
          <input v-model="form.steps[index]" type="text" placeholder="Step" class="form-input flex-1" />
          <button type="button" @click="form.steps.splice(index, 1)" class="text-red-500">Remove</button>
        </div>
        <button type="button" @click="form.steps.push('')" class="text-sm text-green-600 hover:underline">+ Add Step</button>
      </div>

      <button type="submit" class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded font-semibold">
        Submit Recipe
      </button>
    </form>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { navigateTo } from '#app'

const form = reactive({
  title: '',
  description: '',
  image_url: '',
  prep_time: 0,
  chef: '',
  ingredients: [''],
  steps: ['']
})

const CREATE_RECIPE = gql`
  mutation InsertRecipe($object: recipes_insert_input!) {
    insert_recipe_one(object: $object) {
      id
      title
    }
  }
`

const { mutate: insertRecipe, onDone, onError } = useMutation(CREATE_RECIPE)

function handleSubmit() {
  insertRecipe({
    object: {
      ...form
    }
  })
}

onDone(() => {
  alert('✅ Recipe created!')
  navigateTo('/')
})

onError((err) => {
  console.error('Error creating recipe:', err.message)
  alert('❌ Something went wrong')
})
</script>

<style scoped>
.form-input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 0.375rem;
}

.form-textarea {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 0.375rem;
  min-height: 100px;
}
</style>
