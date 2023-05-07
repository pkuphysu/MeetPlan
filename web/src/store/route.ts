// Utilities
import {defineStore} from 'pinia'
import {static_routes} from "@/router";
import {RouteRecordRaw} from "vue-router";

interface State {
  _routes: Array<RouteRecordRaw>,
  _add_routes?: Array<RouteRecordRaw>
}

export const useRouteStore = defineStore('route', {
  state: (): State => ({
    _routes: static_routes,
    _add_routes: undefined
  }),
  actions: {
    clear(){
      this._routes = static_routes
      this._add_routes = undefined
    },
    setRoutes(routes: Array<RouteRecordRaw>) {
      this._routes = this._routes.concat(routes)
      this._add_routes = routes
    }
  },
  getters: {},
})
