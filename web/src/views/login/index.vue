<script setup lang="ts">
import { useI18n } from "vue-i18n";
import Motion from "./utils/motion";
import { useRouter } from "vue-router";
import { message } from "@/utils/message";
import { useLayout } from "@/layout/hooks/useLayout";
import { avatar, bg, illustration } from "./utils/static";
import { onMounted, ref, toRaw } from "vue";
import { useTranslationLang } from "@/layout/hooks/useTranslationLang";
import { useDataThemeChange } from "@/layout/hooks/useDataThemeChange";
import { addPathMatch } from "@/router/utils";
import { usePermissionStoreHook } from "@/store/modules/permission";

import dayIcon from "@/assets/svg/day.svg?component";
import darkIcon from "@/assets/svg/dark.svg?component";
import globalization from "@/assets/svg/globalization.svg?component";
import { useUserStoreHook } from "@/store/modules/user";

defineOptions({
  name: "Login"
});
const router = useRouter();
const loading = ref(false);

const { initStorage } = useLayout();
initStorage();

const { t } = useI18n();
const { dataTheme, dataThemeChange } = useDataThemeChange();
dataThemeChange();
const { changeLocale } = useTranslationLang();

const redirectToAuthLogin = () => {
  window.location.href = `https://auth.phy.pku.edu.cn/oidc/authorize/?response_type=code&scope=openid profile email phone address pku&client_id=16302204390022&redirect_uri=${window.location.origin}/#/login`;
};

const onLogin = async (code: string) => {
  loading.value = true;
  // 全部静态路由
  usePermissionStoreHook().handleWholeMenus([]);
  addPathMatch();
  await useUserStoreHook()
    .loginByCode(code)
    .then(() => {
      message("登录成功", { type: "success" });
      useUserStoreHook()
        .getSelfInfo()
        .then(() => {
          message("获取用户信息成功", { type: "success" });
          window.location.href = `${window.location.origin}/#${router.currentRoute.value.fullPath}`;
        })
        .catch(err => {
          message("获取用户信息失败" + err, { type: "error" });
        });
    })
    .catch(err => {
      message("登录失败" + err, { type: "error" });
    });
  loading.value = false;
};

onMounted(() => {
  useUserStoreHook().userInfo?.id && router.push("/welcome");
  let params = new URLSearchParams(window.location.search);
  let code = params.get("code"); // 获取code
  console.log(router.currentRoute.value);
  if (code) {
    console.log("code", code);
    onLogin(code);
  }
});
</script>

<template>
  <div class="select-none">
    <img :src="bg" class="wave" />
    <div class="flex-c absolute right-5 top-3">
      <!-- 主题 -->
      <el-switch
        v-model="dataTheme"
        inline-prompt
        :active-icon="dayIcon"
        :inactive-icon="darkIcon"
        @change="dataThemeChange"
      />
      <!-- 国际化 -->
      <globalization
        class="hover:text-primary hover:!bg-[transparent] w-[20px] h-[20px] ml-1.5 cursor-pointer outline-none duration-300"
        @click="changeLocale"
      />
    </div>
    <div class="login-container">
      <div class="img">
        <component :is="toRaw(illustration)" />
      </div>
      <div class="login-box">
        <div class="login-form">
          <avatar class="avatar" />
          <Motion>
            <h2 class="outline-none">{{ t("login.siteTitle") }}</h2>
          </Motion>
          <Motion :delay="250">
            <el-button
              class="w-full mt-4"
              size="default"
              type="primary"
              :loading="loading"
              @click="redirectToAuthLogin"
            >
              {{ t("login.redirectToAuthLogin") }}
            </el-button>
          </Motion>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url("@/style/login.css");
</style>

<style lang="scss" scoped>
:deep(.el-input-group__append, .el-input-group__prepend) {
  padding: 0;
}

.translation {
  ::v-deep(.el-dropdown-menu__item) {
    padding: 5px 40px;
  }

  .check-zh {
    position: absolute;
    left: 20px;
  }

  .check-en {
    position: absolute;
    left: 20px;
  }
}
</style>
