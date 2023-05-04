// Composables
import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Guest',
    component: () => import('@/views/Guest.vue'),
    meta: {
      title: '首页',
      layout: 'default',
    }
  },
  {
    path: "/:pathMatch(.*)*",
    name: "Error",
    component: () => import( "@/views/errors/404.vue"),
    meta: {
      title: '404',
      layout: 'default',
    }
  },
]

const dynamicRoutes: Array<RouteRecordRaw> = [
  {
    path: '/home',
    name: 'Overview',
    component: () => import('@/views/Home.vue'),
    meta: {
      role: ['teacher', 'student', 'admin'],
    }
  },
  // {
  //   path: '/profile',
  //   name: 'Profile',
  //   component: () => import('@/views/profile/Profile.vue'),
  //   meta: {
  //     role: ['teacher', 'student'],
  //   }
  // },
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
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }
    return {top: 0}
  }
})

const getTitle = (title?: string) => {
  const siteName = import.meta.env.VITE_APP_TITLE;
  // const siteNameEn = import.meta.env.VITE_APP_TITLE_EN;
  if (title) {
    return `${title} - ${siteName}`
  }
  return `${import.meta.env.VITE_APP_NAME}`
}

router.beforeEach((to) => {
  console.log(to)
  const meta: { [key: string]: any } = to.meta || {};
  document.title = getTitle(meta.title);
})

export const registerDynamicRoutes = (isTeacher: boolean, isAdmin: boolean) => {
  const routes = dynamicRoutes.filter((route) => {
    if (router.hasRoute(route.name as string)){
      return false;
    }
    const meta: {[key:string]: any} = route.meta || {};
    if (!meta.role) {
      return true;
    }
    const roles = meta.role as Array<string>;
    if (roles.includes('teacher') && isTeacher) {
      return true;
    }
    if (roles.includes('student') && !isTeacher) {
      return true;
    }
    return roles.includes('admin') && isAdmin;

  });
  console.log(routes)
  if (routes.length === 0) {
    return false;
  }
  routes.forEach((route) => {
    router.addRoute(route);
  });
  return true;
}

export default router
