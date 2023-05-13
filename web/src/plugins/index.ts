/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Plugins
import { loadFonts } from './webfontloader'
import vuetify from './vuetify'
import router from '../router'
import pinia from '../store'
import i18n from "./i18n";
import Vue3Lottie from 'vue3-lottie'

// Types
import type { App } from 'vue'

export function registerPlugins (app: App) {
  loadFonts()
  app
    .use(vuetify)
    .use(router)
    .use(pinia)
    .use(i18n)
    .use(Vue3Lottie, { name: "LottieAnimation" })
}
