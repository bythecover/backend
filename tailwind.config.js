/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["!./node_modules/", "**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('daisyui'),
  ],
  daisyui: {
    themes: ['light'],
  }
}
