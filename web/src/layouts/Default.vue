<template>
  <MainAppBar/>
  <VMain>
    <VContainer fluid class="maxWidth page-wrapper pb-sm-15 pb-10">
      <slot/>
    </VContainer>
  </VMain>
  <Footer/>
</template>

<script setup lang="ts">
import MainAppBar from "@/components/toolbar/DefaultAppBar.vue";
import Footer from "@/components/Footer.vue";
</script>

<style scoped>
.page-wrapper {
  min-height: calc(100vh - 100px);
  padding: 24px;
}

@media screen and (max-width: 767px) {
  .page-wrapper {
    padding:20px 10px
  }
}
</style>
