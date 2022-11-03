import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: 'localhost',
    cors: true,
    open: true,
    hmr: true,
    watch: {
      usePolling: true
    }
  }
})
