// Utilities
import {defineStore} from 'pinia'

interface State {
  language: string
}

export const useLanguageStore = defineStore( {
  id: 'language',
  state: (): State => ({
    language: 'zh',
  }),
  actions: {
    setLanguage(language: string): void {
      this.language = language
    }
  },
  getters: {},
  persist: {
    enabled: true,
    strategies: [
      {
        storage: localStorage,
      }
    ]
  }
})
