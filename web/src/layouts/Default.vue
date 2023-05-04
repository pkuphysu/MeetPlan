<template>
  <MainAppBar/>
  <v-main>
    <slot/>
  </v-main>
  <v-footer>
    <v-row justify="center" no-gutters>
      <v-col class="text-center mt-0" cols="12">
        <span>{{ new Date().getFullYear() }} - <strong>北京大学物理学院</strong></span>
      </v-col>
    </v-row>
  </v-footer>
</template>

<script setup lang="ts">
import MainAppBar from "@/components/toolbar/MainAppBar.vue";
</script>
