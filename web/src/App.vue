<template>
  <v-app id="app" :theme="theme" >
    <component v-if="routerLoaded">
      <router-view/>
    </component>
  </v-app>

</template>

<script setup lang="ts">
import {useThemeStore} from "@/store/theme";
import {computed, onMounted} from "vue";
import {useRoute} from "vue-router";

const route = useRoute();
const routerLoaded = computed(() => {
  return route.name !== undefined;
})

const themeStore = useThemeStore();
const theme = computed(() => {
  return themeStore.theme;
})

onMounted(() => {
  themeStore.setSystemTheme(window.matchMedia('(prefers-color-scheme: dark)').matches);
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
    themeStore.setSystemTheme(e.matches);
  });
})

</script>
