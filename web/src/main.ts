/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

import App from './App.vue'
import { createApp } from 'vue'
import { registerPlugins } from '@/plugins'
import './router/permission'
import './styles/main.scss'

const app = createApp(App)
registerPlugins(app)
app.mount('#app')
