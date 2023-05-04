<template>
  <v-app id="app" :theme="theme">
    <component :is="currentLayout" v-if="routerLoaded">
      <router-view/>
    </component>
  </v-app>

</template>

<script setup lang="ts">
import {useThemeStore} from "@/store/theme";
import {computed, onBeforeMount, onMounted} from "vue";
import {useRoute} from "vue-router";
import {useUserStore} from "@/store/user";
import Default from "@/layouts/Default.vue";
import router, {registerDynamicRoutes} from "@/router";

const route = useRoute();
const routerLoaded = computed(() => {
  return route.name !== undefined;
})

const layouts: { [key: string]: any } = {
  default: Default,
}

type layoutName = 'default';

const currentLayout = computed(() => {
  const meta: { [key: string]: any } = route.meta || {};
  const layoutName = meta.layout as layoutName;
  if (!layoutName) {
    return Default;
  }
  return layouts[layoutName];
})


const themeStore = useThemeStore();
const theme = computed(() => {
  return themeStore.theme;
})

const userStore = useUserStore();

onBeforeMount(()=>{
  if (userStore.user){
    console.log(userStore.user)
    if (registerDynamicRoutes(userStore.isTeacher, userStore.isAdmin)){
      router.replace(route);
    }
  }
})

onMounted(() => {
  themeStore.setSystemTheme(window.matchMedia('(prefers-color-scheme: dark)').matches);
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
    themeStore.setSystemTheme(e.matches);
  });
})

</script>
