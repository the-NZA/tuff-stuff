<template>
  <nav class="navigation">
    <router-link :to="{ name: 'Index' }">Home</router-link>
    <router-link :to="{ name: 'Lang', params: { lang: 'ru' }}" replace>Русская версия</router-link>
    <router-link :to="{ name: 'Lang', params: { lang: 'en' }}" replace>Английская версия</router-link>
    <router-link :to="{ name: 'Options' }">Настройки</router-link>
  </nav>
</template>

<script setup lang="ts">
</script>

<style scoped>

</style>