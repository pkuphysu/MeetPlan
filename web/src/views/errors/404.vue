<template>
  <v-card variant="flat" height="85vh" class="d-flex justify-center align-center text-center">
    <div>
      <Vue3Lottie
        :animation-data="lottie404"
        :loop="true"
        :autoplay="true"
        height="45vh"
      />
      <v-btn flat color="primary" class="mb-4 text-white" to="/">Go Back to Home</v-btn>
    </div>
  </v-card>
</template>

<script setup lang="ts">
import {Vue3Lottie} from "vue3-lottie";
import lottie404 from "@/assets/lottie/36395-lonely-404.json";
import "vue3-lottie/dist/style.css";
</script>

<style scoped>

</style>
