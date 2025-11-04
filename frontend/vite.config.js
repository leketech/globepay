import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// Determine the backend URL based on environment
const backendUrl = process.env.VITE_API_URL || 'http://localhost:8080';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 3000,
    proxy: {
      '/api': {
        target: backendUrl,
        changeOrigin: true,
        secure: false,
      }
    }
  },
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: true,
  },
  resolve: {
    alias: {
      '@': '/src',
    },
  },
});