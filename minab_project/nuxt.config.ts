import tailwindcss from "@tailwindcss/vite";
export default defineNuxtConfig({
  plugins: ['~/plugins/apollo-client.ts'],
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  vite: {
    plugins: [
      tailwindcss(),
    ],
  },
});