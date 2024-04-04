import { defineStore } from "pinia";
import { store } from "@/store";
import type { userType } from "./types";
import { routerArrays } from "@/layout/types";
import { resetRouter, router } from "@/router";
import type { LoginResult, UserInfo } from "@/api/user";
import { getSelfInfo, login, refreshToken } from "@/api/user";
import { useMultiTagsStoreHook } from "@/store/modules/multiTags";
import { removeToken, setToken } from "@/utils/auth";
import { storageLocal } from "@pureadmin/utils";
import { responsiveStorageNameSpace } from "@/config";

export const useUserStore = defineStore({
  id: "pure-user",
  state: (): userType => ({
    userInfo: storageLocal().getItem<UserInfo>(
      `${responsiveStorageNameSpace()}userInfo`
    ),
    // 页面级别权限
    roles:
      storageLocal().getItem<string[]>(
        `${responsiveStorageNameSpace()}roles`
      ) ?? []
  }),
  actions: {
    async getSelfInfo() {
      return new Promise<UserInfo>((resolve, reject) => {
        getSelfInfo()
          .then(data => {
            if (data && data.data) {
              this.userInfo = data.data;
              storageLocal().setItem(
                `${responsiveStorageNameSpace()}userInfo`,
                data.data
              );
              let roles = [];
              if (data.data.isAdmin) {
                roles.push("admin");
              }
              if (data.data.isTeacher) {
                roles.push("teacher");
              } else {
                roles.push("student");
              }
              this.roles = roles;
              storageLocal().setItem(
                `${responsiveStorageNameSpace()}roles`,
                roles
              );
              resolve(data.data);
            } else {
              reject(data);
            }
          })
          .catch(error => {
            reject(error);
          });
      });
    },
    /** 登入 */
    async loginByCode(code: string) {
      return new Promise<LoginResult>((resolve, reject) => {
        login(code)
          .then(data => {
            if (data && data.data) {
              setToken(data.data);
              getSelfInfo();
              resolve(data);
            } else {
              reject(data);
            }
          })
          .catch(error => {
            reject(error);
          });
      });
    },
    /** 前端登出（不调用接口） */
    logOut() {
      this.roles = [];
      this.userInfo = undefined;
      storageLocal().removeItem(`${responsiveStorageNameSpace()}userInfo`);
      storageLocal().removeItem(`${responsiveStorageNameSpace()}roles`);
      removeToken();
      useMultiTagsStoreHook().handleTags("equal", [...routerArrays]);
      resetRouter();
      router.push("/login");
    },
    /** 刷新`token` */
    async handRefreshToken(token: string) {
      return new Promise<LoginResult>((resolve, reject) => {
        refreshToken(token)
          .then(data => {
            if (data && data.data) {
              setToken(data.data);
              resolve(data);
            } else {
              reject(data);
            }
          })
          .catch(error => {
            reject(error);
          });
      });
    }
  }
});

export function useUserStoreHook() {
  return useUserStore(store);
}
