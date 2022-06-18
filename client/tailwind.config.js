/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      minWidth: {
        '111': '8.5rem'
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
