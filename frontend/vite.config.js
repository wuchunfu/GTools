import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  publicPath: './',
  server: {
    host: 'localhost',
    cors: true,
    open: false,
    hmr: true,
    watch: {
      usePolling: true
    }
  }
})
