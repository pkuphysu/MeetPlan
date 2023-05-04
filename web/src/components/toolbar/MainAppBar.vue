<script setup lang="ts">
import {useUserStore} from "@/store/user";
// import {useDisplay} from "vuetify";
import {computed} from "vue";
import {useThemeStore} from "@/store/theme";

const userStore = useUserStore();
const themeStore = useThemeStore();

const logoUrl = computed(() => {
  return themeStore.theme === 'dark' ? '@/assets/logo-dark.png' : '@/assets/logo.png';
})

const hasLogin = computed(() => {
  return userStore.user !== undefined;
})

// const {mdAndUp} = useDisplay();

const loginUrl = `https://auth.phy.pku.edu.cn/oidc/authorize/?response_type=code&scope=openid profile email address pku&client_id=16302204390022&redirect_uri=${import.meta.env.VITE_HOST_URL}`

</script>

<template>
  <v-app-bar density="default" elevation="0">
    <div class="plan-box v-toolbar__content px-lg-4 px-4">
      <!--      <div class="hidden-md-and-down">-->
      <div class="logo" v-if="hasLogin">
        <router-link to="/" class="d-flex">
          <v-img :src="logoUrl" alt="Home" width="130"></v-img>
        </router-link>
      </div>
      <!--      </div>-->
      <div class="flex-grow-1"></div>
      <v-btn :href="loginUrl" v-if="!hasLogin"
        class="float-right text-white bg-background font-weight-bold"
      >统一认证登录
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
