<template>
  <div class="max-w-2xl mx-auto mt-10 bg-white p-6 rounded-xl shadow-md">
    <h2 class="text-2xl font-bold mb-6">Create a Recipe</h2>
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <input v-model="title" placeholder="Title" required class="input" />
      <textarea
        v-model="description"
        placeholder="Description"
        class="input"
      ></textarea>
      <input v-model="category" placeholder="Category" class="input" />

      <!-- Image Upload UI -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2"
          >Upload Images</label
        >
        <div
          class="border-2 border-dashed border-gray-300 rounded-lg p-4 text-center cursor-pointer hover:border-green-500 transition"
        >
          <input
            type="file"
            accept="image/*"
            multiple
            @change="handleImageUpload"
            class="hidden"
            ref="fileInput"
          />
          <div @click="$refs.fileInput.click()" class="text-gray-500">
            📁 Click or drag to upload images
          </div>
        </div>
        <div class="mt-4 grid grid-cols-2 gap-4">
          <div
            v-for="(image, index) in previewImages"
            :key="index"
            class="border rounded overflow-hidden hover:shadow-lg transition"
          >
            <img :src="image" class="w-full h-32 object-cover" alt="Preview" />
          </div>
        </div>
      </div>

      <!-- Ingredients -->
      <div>
        <h3 class="font-semibold">Ingredients</h3>
        <div
          v-for="(ing, index) in ingredients"
          :key="index"
          class="flex gap-2 mt-2"
        >
          <input
            v-model="ing.name"
            placeholder="Name (e.g., Flour)"
            class="input"
          />
          <input
            v-model="ing.quantity"
            placeholder="Quantity (e.g., 2 cups)"
            class="input"
          />
          <button type="button" @click="ingredients.splice(index, 1)">
            ❌
          </button>
        </div>
        <button
          type="button"
          @click="ingredients.push({ name: '', quantity: '' })"
          class="text-green-600 mt-2"
        >
          + Add Ingredient
        </button>
      </div>

      <!-- Steps -->
      <div>
        <h3 class="font-semibold">Steps</h3>
        <div
          v-for="(step, index) in steps"
          :key="index"
          class="flex gap-2 mt-2"
        >
          <input
            v-model="step.instruction"
            placeholder="Step description"
            class="input"
          />
          <button type="button" @click="steps.splice(index, 1)">❌</button>
        </div>
        <button
          type="button"
          @click="steps.push({ instruction: '' })"
          class="text-green-600 mt-2"
        >
          + Add Step
        </button>
      </div>

      <button
        type="submit"
        class="w-full bg-green-600 text-white py-2 rounded-md hover:bg-green-700"
      >
        Create
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import { useRouter } from "vue-router";
import gql from "graphql-tag";

const title = ref("");
const description = ref("");
const category = ref("");
const images = ref([]);
const previewImages = ref([]);
const ingredients = ref([{ name: "", quantity: "" }]);
const steps = ref([{ instruction: "" }]);
const router = useRouter();

const handleImageUpload = (event) => {
  images.value = Array.from(event.target.files);
  previewImages.value = images.value.map((file) => URL.createObjectURL(file));
};

const INSERT_RECIPE = gql`
  mutation InsertRecipe($object: recipes_insert_input!) {
    insert_recipes_one(object: $object) {
      id
      title
    }
  }
`;

const { mutate } = useMutation(INSERT_RECIPE);

const handleSubmit = async () => {
  const userId = JSON.parse(atob(localStorage.getItem("token").split(".")[1]))[
    "https://hasura.io/jwt/claims"
  ]["x-hasura-user-id"];

  try {
    const imageUrls = images.value.map(
      (file) => "https://dummy.url/" + file.name
    );

    await mutate({
      object: {
        title: title.value,
        description: description.value,
        category: category.value,
        user_id: userId,
        images: imageUrls,
        thumbnail: imageUrls[0] || null,
        ingredients: {
          data: ingredients.value.map((i) => ({
            name: i.name,
            quantity: i.quantity,
          })),
        },
        steps: {
          data: steps.value.map((s, idx) => ({
            instruction: s.instruction,
            step_number: idx + 1,
          })),
        },
      },
    });

    alert("✅ Recipe created!");
    router.push("/categories");
  } catch (err) {
    console.error("Error:", err);
    alert("❌ Failed to create recipe");
  }
};
</script>

<style scoped>
.input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 8px;
}
</style>
