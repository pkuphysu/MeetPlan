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
  },
  {
    path: "/users",
    redirect: "/user/list",
    name: "UserList",
    meta: {
      icon: "ri:user-settings-line",
      title: $t("menus.userManagement")
    },
    children: [
      {
        path: "/user/list",
        name: "UserList",
        component: () => import("@/views/user/list/index.vue"),
        meta: {
          icon: "ri:user-settings-line",
          title: $t("menus.userList"),
          roles: ["admin"],
          keepAlive: true
        }
      },
      {
        path: "/user/department",
        name: "DepartmentList",
        component: () => import("@/views/user/department/index.vue"),
        meta: {
          icon: "mingcute:department-line",
          title: $t("menus.departmentList"),
          roles: ["admin"],
          keepAlive: true
        }
      },
      {
        path: "/user/major",
        name: "MajorList",
        component: () => import("@/views/user/major/index.vue"),
        meta: {
          icon: "icon-park-outline:degree-hat",
          title: $t("menus.majorList"),
          roles: ["admin"],
          keepAlive: true
        }
      },
      {
        path: "/user/grade",
        name: "GradeList",
        component: () => import("@/views/user/grade/index.vue"),
        meta: {
          icon: "ph:users",
          title: $t("menus.gradeList"),
          roles: ["admin"],
          keepAlive: true
        }
      }
    ]
  }
] satisfies Array<RouteConfigsTable>;
