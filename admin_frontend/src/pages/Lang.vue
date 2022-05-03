<template>
  <Navigation/>
  <h1>This is lang page for {{ lang }}</h1>
</template>

<script setup lang="ts">
import Navigation from '../components/Navigation.vue';
import {useRoute, onBeforeRouteUpdate} from "vue-router";
import {ref} from "vue";

const route = useRoute();
const lang = ref<string>(route.params.lang as string);

// Update the lang ref when the route changes
// get data from backend
onBeforeRouteUpdate((to, from) => {
  lang.value = to.params.lang as string;
});


</script>

<style scoped>

</style>