import {useUserStore} from "@/store/user";
import {useRouteStore} from "@/store/route";
import router, {resetRouter} from "@/router";



export const logout = () => {
  const userStore = useUserStore();
  const routeStore = useRouteStore();

  userStore.clear();
  routeStore.clear();
  resetRouter();

  router.push({name: 'Home'});
}
