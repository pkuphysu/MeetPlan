// Plugins
import vue from '@vitejs/plugin-vue'
import vuetify, {transformAssetUrls} from 'vite-plugin-vuetify'
import Components from 'unplugin-vue-components/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'

// Utilities
import {defineConfig} from 'vite'
import {fileURLToPath, URL} from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      template: {transformAssetUrls}
    }),
    // https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vite-plugin
    vuetify({
      autoImport: true,
    }),
    AutoImport({
      imports: ['vue', 'pinia', 'vue-router', 'vue-i18n'],
      dirs: ['src/router', 'src/store'],
      resolvers: [
        IconsResolver()
      ],
      dts: "src/types/auto-imports.d.ts"
    }),
    Icons({
      autoInstall: true,
      compiler: 'vue3'
    }),
    Components({
      dts: 'src/types/components.d.ts',
      dirs: ['src/components'],
      extensions: ['vue'],
      resolvers: [
        // 自动按需加载iconify图标库图标
        IconsResolver()
      ]
    }),
  ],
  define: {'process.env': {}},
  resolve: {
    alias: {
      '~': fileURLToPath(new URL('./', import.meta.url)),
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: [
      '.js',
      '.json',
      '.jsx',
      '.mjs',
      '.ts',
      '.tsx',
      '.vue',
    ],
  },
  server: {
    port: 3000,
    open: true,
    host: '0.0.0.0',
    proxy: undefined
  },
  css: {
    preprocessorOptions: {
      sass: {
        additionalData: `@import "./src/styles/variables.scss";`
      }
    }
  }
})
