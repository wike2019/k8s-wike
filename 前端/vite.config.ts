import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
export default defineConfig({
  resolve: {
    alias: {
      '@': '/src'
    }
  },
  plugins: [vue()],
  server: {
    host:"0.0.0.0",
    port:2002,
    proxy: {
      // 字符串简写写法
      '/v1': 'http://127.0.0.1:7772/v1',
    }
  },
  //base: '//csdn.52wike.com/test/',

})
