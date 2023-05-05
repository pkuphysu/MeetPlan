<script setup lang="ts">
import {ref} from 'vue';
import {useRoute} from "vue-router";
import {getSelf, login, LoginParams} from "@/api/user";
import {useUserStore} from "@/store/user";
import router from "@/router";
import LoginLoader from "@/components/animations/LoginLoader.vue";
import {loginRedirectUrl} from "@/utils/constants";

const userStore = useUserStore();
const route = useRoute();

const text = ref('')

const redirectHome = () => {
  router.push({path: '/dashboard'})
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
      redirectHome();
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
    redirectHome();
  }).catch((err) => {
    console.log(err)
  })
} else {
  text.value = '未登录，跳转中。。。'
  window.location.href = loginRedirectUrl;
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
