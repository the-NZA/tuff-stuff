<template>
  <div class="wrapper">
    <LoginForm/>
  </div>
</template>

<script setup lang="ts">
import LoginForm from "../components/LoginForm.vue"
</script>

<style scoped lang="postcss">
.wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
</style>