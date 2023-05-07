import {useUserStore} from "@/store/user";
import {useRouteStore} from "@/store/route";
import {resetRouter} from "@/router";
import {useRouter} from "vue-router";



export const logout = () => {
  const userStore = useUserStore();
  const routeStore = useRouteStore();
  const router = useRouter();

  userStore.clear();
  routeStore.clear();
  resetRouter();

  router.push({name: 'Home'});
}
