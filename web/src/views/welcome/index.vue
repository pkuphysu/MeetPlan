<script setup lang="ts">
import { useUserStoreHook } from "@/store/modules/user";
import WelcomeTeacher from "@/views/welcome/teacher/index.vue";
import WelcomeStudent from "@/views/welcome/student/index.vue";

defineOptions({
  name: "Welcome"
});

const userStore = useUserStoreHook();
</script>

<template>
  <WelcomeTeacher v-if="userStore.userInfo.isTeacher" />
  <WelcomeStudent v-else />
</template>
