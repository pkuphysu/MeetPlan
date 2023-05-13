<script setup lang="ts">
import "vue3-lottie/dist/style.css";
import loginLoader from "@/assets/lottie/97018-window10-style-login-loader.json";

interface Props {
  size?: number
}
const props = withDefaults(defineProps<Props>(), {
  size: 500
})
</script>

<template>
    <Vue3Lottie
      :animation-data="loginLoader"
      :loop="true"
      :autoplay="true"
      :height="props.size"
      :width="props.size"
    />
</template>

<style scoped>

</style>
