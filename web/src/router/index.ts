// Composables
import {createRouter, createWebHistory, Router, RouteRecordRaw} from 'vue-router'

export const static_routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Guest.vue'),
    meta: {
      title: '首页',
      layout: 'default',
      needAuth: false,
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/Login.vue'),
    meta: {
      title: '登录',
      layout: 'default',
      needAuth: false,
    }
  },
  {
    path: '/:catchAll(.*)',
    name: 'Redirect',
    redirect: to => {return { name: 'Login', query: { redirect: to.path }}},
  }
]

export const dynamic_routes: Array<RouteRecordRaw> = [
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('@/views/student/Home.vue'),
    meta: {
      title: '主页',
      role: ['student'],
      needAuth: true,
    }
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('@/views/teacher/Home.vue'),
    meta: {
      title: '主页',
      role: ['teacher'],
      needAuth: true,
    }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/teacher/Profile.vue'),
    meta: {
      role: ['teacher'],
      needAuth: true,
    }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/student/Profile.vue'),
    meta: {
      role: ['student'],
      needAuth: true,
    }
  },
  // {
  //   path: '/meetplan',
  //   name: 'MeetPlan',
  //   component: () => import('@/views/meetplan/Profile.vue'),
  //   meta: {
  //     role: ['teacher', 'student'],
  //   }
  // },
  // {
  //   path: '/meetplanorder',
  //   name: 'MeetPlanOrder',
  //   component: () => import('@/views/meetplan/Order.vue'),
  //   meta: {
  //     role: ['teacher', 'student'],
  //   }
  // }
  {
    path: "/:pathMatch(.*)*",
    name: "Error",
    component: () => import( "@/views/errors/404.vue"),
    meta: {
      title: '404',
      layout: 'default',
      needAuth: true,
    }
  },
]

const initRouter = (): Router =>{
  return createRouter({
    history: createWebHistory(),
    routes: static_routes,
    scrollBehavior(to, from, savedPosition) {
      if (savedPosition) {
        return savedPosition
      }
      return {top: 0}
    }
  })
}

const router = initRouter()

export function resetRouter () {
  const newRouter = initRouter()
  // @ts-ignore bef
  router.matcher = newRouter.matcher
}

export default router
