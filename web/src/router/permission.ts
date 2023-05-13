import router, {dynamic_routes} from '@/router/index';
import {useRouteStore} from "@/store/route";
import {useUserStore} from "@/store/user";

router.beforeEach((to, from, next) => {
  console.log('to', to)
  console.log('from', from)
  console.log('next', next)
  const meta: { [key: string]: any } = to.meta || {};
  document.title = getTitle(meta.title);

  const routeStore = useRouteStore();
  const userStore = useUserStore();

  if (routeStore._add_routes) {
    if (to.path === '/login') {
      next(to.query.redirect? {path: to.query.redirect as string}:{path: '/dashboard'})
    } else {
      next()
    }
  } else {
    // 没加过动态路由
    // 1. 登录完成
    // 2. 刷新页面
    console.log(userStore.user)
    if (userStore.user) {
      const add_routes = registerDynamicRoutes(userStore.user.is_teacher, userStore.user.is_admin)
      routeStore.setRoutes(add_routes)
      console.log(router.getRoutes())
      console.log(to)
      next(to)
    } else {
      if (router.hasRoute(to.name as string) && !(to.meta.needAuth as boolean)) {
        next()
      }else {
        next(`/login?redirect=${to.path}`)
      }
    }
  }
})

export const registerDynamicRoutes = (isTeacher: boolean, isAdmin: boolean) => {
  const routes = dynamic_routes.filter((route) => {
    const meta: { [key: string]: any } = route.meta || {};
    if (meta.needAuth?.value === false) {
      return true;
    }
    const roles = meta.role as Array<string>;
    if (!roles) {
      return true;
    }
    if (isTeacher && roles.includes('teacher')) {
      return true;
    }
    if (!isTeacher && roles.includes('student')) {
      return true;
    }
    return isAdmin && roles.includes('admin');
  });

  routes.forEach((route) => {
    router.addRoute(route);
  });
  router.removeRoute('Redirect');
  return routes;
}

const getTitle = (title?: string) => {
  const siteName = import.meta.env.VITE_APP_TITLE;
  // const siteNameEn = import.meta.env.VITE_APP_TITLE_EN;
  if (title) {
    return `${title} - ${siteName}`
  }
  return `${import.meta.env.VITE_APP_NAME}`
}
