/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './node_modules/nextra-theme-docs/**/*.js',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
  darkMode: 'class',
}
