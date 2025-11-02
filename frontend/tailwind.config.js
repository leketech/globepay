/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        "primary": "#0A4D68",
        "background-light": "#F9F9F9",
        "background-dark": "#1A1A1A",
        "secondary": "#00C853",
        "text-light": "#212121",
        "text-dark": "#f4f6f8",
        "neutral-light": "#F4F6F8",
        "neutral-dark": "#2a3b44",
        "border-light": "#dbe2e6",
        "border-dark": "#3a4c56",
        "destructive": "#DE350B",
        "success": "#00875A",
      },
      fontFamily: {
        "display": ["Manrope", "sans-serif"]
      },
      borderRadius: {
        "DEFAULT": "0.25rem",
        "lg": "0.5rem",
        "xl": "0.75rem",
        "full": "9999px"
      },
      screens: {
        'sm': '640px',
        'md': '768px',
        'lg': '1024px',
        'xl': '1280px',
        '2xl': '1536px',
      }
    },
  },
  plugins: [],
  corePlugins: {
    preflight: true,
  }
}