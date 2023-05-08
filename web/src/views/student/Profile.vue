<script setup lang="ts">
import Profile from "@/components/profile/Profile.vue";
import {useUserStore} from "@/store/user";

const userStore = useUserStore();

</script>

<template>
  <v-row>
    <Profile :userid="userStore.user!.id"></Profile>
  </v-row>
</template>

<style scoped>

</style>
