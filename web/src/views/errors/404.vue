<template>
  <v-container class="fill-height">
    <v-responsive class="d-flex justify-center align-center text-center">
      <div>
        <Vue3Lottie
          :animation-data="lottie404"
          :loop="true"
          :autoplay="true"
          height="70vh"
        />
        <v-btn flat color="primary" class="mb-4 text-white" to="/">Go Back to Home</v-btn>
      </div>
    </v-responsive>
<!--    <v-card variant="flat"  class="d-flex justify-center align-center text-center">-->
<!--    </v-card>-->
  </v-container>
</template>

<script setup lang="ts">
import lottie404 from "@/assets/lottie/36395-lonely-404.json";
import "vue3-lottie/dist/style.css";
</script>

<style scoped>

</style>
