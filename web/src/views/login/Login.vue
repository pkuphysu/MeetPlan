<script setup lang="ts">
import {ref} from 'vue';
import {useRoute} from "vue-router";
import {getSelf, login, LoginParams, User} from "@/api/user";
import {useUserStore} from "@/store/user";
import router from "@/router";
import LoginLoader from "@/components/animations/LoginLoader.vue";
import {registerDynamicRoutes} from "@/router/permission";
import {loginRedirectUrl} from "@/utils/utils";

const userStore = useUserStore();
const route = useRoute();

const text = ref('')

const getRedirectTarget = () =>{
  return sessionStorage.getItem('redirect')
}

if (route.query.redirect) {
  sessionStorage.setItem('redirect', route.query.redirect as string)
}

const redirectHome = (user: User) => {
  registerDynamicRoutes(user.is_teacher, user.is_admin)
  console.log(route)
  router.push( getRedirectTarget()? {path: getRedirectTarget()} : {path: '/dashboard'})
  sessionStorage.removeItem('redirect')
}

if (route.query['code']) {
  text.value = '登录中。。。'
  console.log(route.query['code'])
  var params: LoginParams = {
    code: route.query['code'].toString(),
  }
  login(params).then((res) => {
    text.value = 'code校验通过，获取用户信息中。。。'
    userStore.setJwt(res);
    getSelf().then((res) => {
      text.value = '获取用户信息成功，跳转中。。。'
      userStore.setUser(res);
      redirectHome(res);
    }).catch((err) => {
      userStore.clear();
      console.log(err)
    })
  }).catch((err) => {
    console.log(err)
  })
} else if (userStore.jwt) {
  text.value = '已登录，获取用户信息中。。。'
  getSelf().then((res) => {
    userStore.setUser(res);
    redirectHome(res);
  }).catch((err) => {
    console.log(err)
  })
} else {
  text.value = '未登录，跳转中。。。'
  window.location.href = loginRedirectUrl();
}

</script>

<template>
  <v-container class="text-center" height="60vh">
    <v-responsive height="60vh">
      <LoginLoader/>
    </v-responsive>
    <div class="py-6 text-body-1 font-weight-regular mb-n1">{{ text }}</div>

  </v-container>
</template>

<style scoped>

</style>
