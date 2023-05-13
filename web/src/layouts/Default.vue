<template>
  <DefaultAppBar/>
  <VMain>
    <VContainer fluid class="maxWidth page-wrapper pb-sm-15 pb-10">
      <slot/>
    </VContainer>
  </VMain>
  <DefaultFooter/>
</template>

<script setup lang="ts">
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
