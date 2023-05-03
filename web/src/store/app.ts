// Utilities
import { defineStore } from 'pinia'

interface State{

}

export const useAppStore = defineStore('app', {
  state: (): State => ({
    alertMessages: [],
    user: JSON.parse(
      sessionStorage.getItem('user') || localStorage.getItem('user') || '{}'
    )
  }),
  actions: {

  }
})
