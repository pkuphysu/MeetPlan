// Utilities
import {defineStore} from 'pinia'

interface State {
  darkTheme?: boolean
  systemTheme?: boolean
}

export const useThemeStore = defineStore( {
  id: 'theme',
  state: (): State => ({
    darkTheme: undefined,
    systemTheme: undefined
  }),
  actions: {
    setDarkTheme(value?: boolean) {
      this.darkTheme = value
    },
    setSystemTheme(value?: boolean) {
      this.darkTheme = value
    }
  },
  getters: {
    theme(): string{
      if (this.darkTheme === undefined) {
        return this.systemTheme ? 'dark' : 'light'
      }
      return this.darkTheme ? 'dark' : 'light'
    }
  },
  persist: {
    enabled: true,
    strategies: [
      {
        storage: localStorage,
      }
    ]
  }
})
