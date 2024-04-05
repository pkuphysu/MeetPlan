import { $t } from "@/plugins/i18n";

export default [
  {
    path: "/users",
    redirect: "/user/list",
    meta: {
      icon: "ri:user-settings-line",
      title: $t("menus.userManagement")
    },
    children: [
      {
        path: "/user/list",
        name: "UserList",
        component: () => import("@/views/admin/user/index.vue"),
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
        component: () => import("@/views/admin/department/index.vue"),
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
        component: () => import("@/views/admin/major/index.vue"),
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
        component: () => import("@/views/admin/grade/index.vue"),
        meta: {
          icon: "ph:users",
          title: $t("menus.gradeList"),
          roles: ["admin"],
          keepAlive: true
        }
      }
    ]
  },
  {
    path: "/system",
    redirect: "/system/option",
    meta: {
      icon: "ri:user-settings-line",
      title: $t("menus.optionList")
    },
    children: [
      {
        path: "/system/option",
        name: "OptionList",
        component: () => import("@/views/admin/option/index.vue"),
        meta: {
          icon: "ri:user-settings-line",
          title: $t("menus.optionList"),
          roles: ["admin"],
          keepAlive: true
        }
      }
    ]
  }
] satisfies Array<RouteConfigsTable>;
