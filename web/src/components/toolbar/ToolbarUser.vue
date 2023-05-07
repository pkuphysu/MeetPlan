<script setup lang="ts">
import {useRouter} from "vue-router";
import {logout} from "@/utils/auth";
import {useUserStore} from "@/store/user";
import {computed} from "vue";
import {Md5} from "ts-md5"

const userStore = useUserStore();
const router = useRouter();

const currentUser = computed(() => {
  return userStore.user!
})

const userAvatar = computed(() => {
  if (currentUser.value.avatar) {
    return currentUser.value.avatar
  }
  return `https://cravatar.cn/avatar/${Md5.hashStr(currentUser.value.email)}.png`
});

const handleLogout = () => {
  logout();
  console.log("---");
  console.log(router);
};
</script>

<template>
  <v-menu
    :close-on-content-click="false"
    open-on-hover
    location="bottom right"
    transition="slide-y-transition"
  >
    <!-- ---------------------------------------------- -->
    <!-- Activator Btn -->
    <!-- ---------------------------------------------- -->
    <template v-slot:activator="{ props }">
      <v-btn class="mx-2" icon v-bind="props">
        <v-avatar size="40">
          <v-img :src="userAvatar"></v-img>
        </v-avatar>
      </v-btn>
    </template>
    <v-card max-width="300">
      <v-list lines="three" density="compact">
        <!-- ---------------------------------------------- -->
        <!-- Profile Area -->
        <!-- ---------------------------------------------- -->
        <v-list-item :to="{name: 'Profile'}">
          <template v-slot:prepend>
            <v-avatar size="40">
              <v-img :src="userAvatar"></v-img>
            </v-avatar>
          </template>

          <v-list-item-title class="font-weight-bold text-primary">
            {{ currentUser.name }}
          </v-list-item-title>
          <v-list-item-subtitle>
            {{ currentUser.email }}
          </v-list-item-subtitle>
        </v-list-item>
      </v-list>
      <v-divider/>
      <!-- ---------------------------------------------- -->
      <!-- Menu Area -->
      <!-- ---------------------------------------------- -->
      <v-list variant="flat" elevation="0" :lines="false" density="compact">
        <v-list-item color="primary" to="/profile" link density="compact">
          <template v-slot:prepend>
            <v-avatar size="30"><v-icon>mdi-account-box-outline</v-icon></v-avatar>
          </template>
          <div>
            <v-list-item-subtitle class="text-body-2">Profile Details</v-list-item-subtitle>
          </div>
        </v-list-item>
        <v-list-item color="primary" link density="compact">
          <template v-slot:prepend>
            <v-avatar size="30"><v-icon>mdi-help-circle-outline</v-icon></v-avatar>
          </template>
          <div>
            <v-list-item-subtitle class="text-body-2">Contact Admin</v-list-item-subtitle>
          </div>
        </v-list-item>
      </v-list>
      <v-divider/>
      <!-- ---------------------------------------------- -->
      <!-- Logout Area -->
      <!-- ---------------------------------------------- -->
      <v-list variant="flat" elevation="0" :lines="false" density="compact">
        <v-list-item
          color="primary"
          link
          @click="handleLogout"
          density="compact"
        >
          <template v-slot:prepend>
            <v-avatar size="30">
              <v-icon>mdi-logout</v-icon>
            </v-avatar>
          </template>

          <div>
            <v-list-item-subtitle class="text-body-2">
              Logout
            </v-list-item-subtitle>
          </div>
        </v-list-item>
      </v-list>
    </v-card>
  </v-menu>
</template>

<style scoped>

</style>

