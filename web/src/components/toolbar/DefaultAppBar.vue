<script setup lang="ts">
import {useUserStore} from "@/store/user";
// import {useDisplay} from "vuetify";
import {computed} from "vue";
import {useThemeStore} from "@/store/theme";
import {loginRedirectUrl} from "@/utils/constants";
import ToolbarLanguage from "@/components/toolbar/ToolbarLanguage.vue";
import ToolbarUser from "@/components/toolbar/ToolbarUser.vue";

const userStore = useUserStore();
const themeStore = useThemeStore();

const hasLogin = computed(() => {
  return userStore.user !== undefined;
})

// const {mdAndUp} = useDisplay();

</script>

<template>
  <v-app-bar density="default" elevation="0">
    <div class="maxWidth v-toolbar__content px-lg-4 px-4">
      <!--      <div class="hidden-md-and-down">-->
      <div class="logo" v-if="hasLogin">
        <router-link to="/" class="d-flex">
          <v-img v-if="themeStore.theme === 'dark'" src="@/assets/logo-dark.png" alt="Home" width="125"></v-img>
          <v-img v-else src="@/assets/logo.png" alt="Home" width="125"></v-img>
        </router-link>
      </div>

      <v-row v-if="hasLogin" class="pl-4 hidden-sm-and-down ">
        <router-link :to="{name: 'Dashboard'}" class="text-decoration-none flex-0-0">
          <v-btn density="default" size="large" style="color: #2a3547">
            首页
          </v-btn>
        </router-link>
        <router-link :to="{name: 'Profile'}" class="text-decoration-none flex-0-0">
          <v-btn density="default" size="large" style="color: #2a3547">
            我
          </v-btn>
        </router-link>
      </v-row>

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

@media screen and (max-width: 1024px) {
  .logo {
    width: auto;
  }
}
</style>
