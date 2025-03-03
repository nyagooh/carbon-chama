/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{html,js}",
  ],
  theme: {
    extend: {
      colors: {
        'dark-purple': '#340D5C',
        'light-purple': '#8A2BE2',
        'yellow': '#FFD700',
        'medium-purple': '#9B59B6',
        'pinkish-purple': '#E74C3C',
      },
    },
  },
  plugins: [],
}