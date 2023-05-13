// Utilities
import {defineStore} from 'pinia'

interface State {
  _user?: User,
  _jwt?: string
}

export const useUserStore = defineStore('user', {
  state: (): State => ({
    _user: undefined,
    _jwt: undefined
  }),
  actions: {
    setUser(user: User) {
      this._user = user
    },
    setJwt(jwt: string) {
      this._jwt = jwt
    },
    clear(){
      this._user = undefined
      this._jwt = undefined
    }
  },
  getters: {
    user(): User | undefined {
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
