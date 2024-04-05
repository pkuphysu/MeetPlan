import { $t } from "@/plugins/i18n";

export default [
  {
    path: "/profile",
    name: "UserInfo",
    component: () => import("@/views/user/profile/index.vue"),
    meta: {
      icon: "ep:avatar",
      title: $t("menus.personalCenter")
    }
  }
] satisfies Array<RouteConfigsTable>;
