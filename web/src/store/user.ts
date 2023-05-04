// Utilities
import {defineStore} from 'pinia'

interface State {
  _user?: {
    id: number
    pku_id: string
    name: string
    is_teacher: boolean
    is_admin: boolean
  },
  _jwt?: string
}

export const useUserStore = defineStore('user', {
  state: (): State => ({
    _user: undefined,
    _jwt: undefined
  }),
  actions: {
    setUser(user: State['_user']) {
      this._user = user
    },
    setJwt(jwt: State['_jwt']) {
      this._jwt = jwt
    },
    clear(){
      this._user = undefined
      this._jwt = undefined
    }
  },
  getters: {
    user(): State['_user'] {
      return this._user
    },
    isTeacher(): boolean {
      return this._user?.is_teacher ?? false
    },
    isAdmin(): boolean {
      return this._user?.is_admin ?? false
    },
    jwt(): string {
      return this._jwt ?? ''
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
