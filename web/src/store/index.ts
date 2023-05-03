// Utilities
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persist'

export default createPinia().use(piniaPersist)
