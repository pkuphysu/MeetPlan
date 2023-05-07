<template>
  <MainAppBar/>
  <v-main>
    <slot/>
  </v-main>
  <Footer/>
</template>

<script setup lang="ts">
import MainAppBar from "@/components/toolbar/MainAppBar.vue";
import Footer from "@/components/Footer.vue";
</script>
