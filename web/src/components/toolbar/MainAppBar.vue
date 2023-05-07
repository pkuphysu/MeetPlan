<script setup lang="ts">
import {useUserStore} from "@/store/user";
// import {useDisplay} from "vuetify";
import {computed} from "vue";
import {useThemeStore} from "@/store/theme";
import {loginRedirectUrl} from "@/utils/constants";
import ToolbarUser from "@/components/toolbar/ToolbarUser.vue";
import ToolbarLanguage from "@/components/toolbar/ToolbarLanguage.vue";

const userStore = useUserStore();
const themeStore = useThemeStore();

const hasLogin = computed(() => {
  return userStore.user !== undefined;
})

// const {mdAndUp} = useDisplay();

</script>

<template>
  <v-app-bar density="default" elevation="0">
    <div class="plan-box v-toolbar__content px-lg-4 px-4">
      <!--      <div class="hidden-md-and-down">-->
      <div class="logo" v-if="hasLogin">
        <router-link to="/" class="d-flex">
          <v-img v-if="themeStore.theme === 'dark'" src="@/assets/logo-dark.png" alt="Home" width="125"></v-img>
          <v-img v-else src="@/assets/logo.png" alt="Home" width="125"></v-img>
        </router-link>
      </div>
      <template v-if="hasLogin">
        <router-link to="/dashboard">主页</router-link>
      </template>

      <div class="flex-grow-1"></div>
      <ToolbarLanguage/>
      <template v-if="hasLogin">
        <ToolbarUser/>
      </template>
      <v-btn
        :href="loginRedirectUrl"
        v-if="!hasLogin"
        class="float-right text-white bg-background font-weight-bold"
      >
        统一认证登录
      </v-btn>
    </div>
  </v-app-bar>
</template>

<style scoped>
.bg-background {
  background-color: #8c0000 !important;
  color: #fff !important;
}

.plan-box {
  max-width: 1200px;
  margin: 0 auto;
}

@media screen and (max-width: 1024px) {
  .logo {
    width: auto;
  }
}

</style>
