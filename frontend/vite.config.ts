import { defineConfig } from 'vite'
import reactRefresh from '@vitejs/plugin-react-refresh'
import { babel } from '@rollup/plugin-babel';
import replace from '@rollup/plugin-replace';
import macros from "vite-plugin-babel-macros"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
      macros(),
      reactRefresh(),
  ],
  mode: 'development',
    optimizeDeps: {
      exclude: ['react-switch']
    },
  build: {
    minify: false,
    rollupOptions: {
        external: ['react-switch'],
      plugins: [
          babel(),
          replace({
            'process.env.NODE_ENV': JSON.stringify('production'),
          }),
      ]
    }
  }
})
