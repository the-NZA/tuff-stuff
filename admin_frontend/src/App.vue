<script setup lang="ts">
import {onBeforeMount} from "vue";
import {useAuthStore} from "./store/auth";
import {useRouter} from "vue-router";

const authStore = useAuthStore();
const router = useRouter();

onBeforeMount(() => {
})

</script>

<template>
	<router-view/>
</template>

<style lang="postcss">
:root {
	/* colors */
	--black: 0, 0, 0;
	--text-black: 1, 22, 30;
	--white: 255, 255, 255;
	--gray: 100, 100, 100;
	--gray-light: 250, 250, 250;
	--red: 255, 0, 0;
	--purple: 128, 0, 128;

	/* opacity */
	--alpha100: 1;
	--alpha50: 0.5;
	--alpha25: 0.25;

	/* offsets */
	--offset: 32px;
	--offset_half: calc(var(--offset) / 2);
	--offset_third: calc(var(--offset) / 3);
	--offset_quarter: calc(var(--offset) / 4);
	--offset_eight: calc(var(--offset) / 8);
	--offset_twice: calc(var(--offset) * 2);
	--offset_triply: calc(var(--offset) * 3);

	/* sizes */
	--siteWidth: 1200px;
	--loginFormWidth: calc(var(--siteWidth) / 3);

	/* transitions */
	--mainTransition: all 150ms ease-in-out;

	/* shadows */

}

* {
	box-sizing: border-box;
}

body {
	height: 100%;

	background-color: rgb(var(--gray-light));
}

#app {
	min-height: 100vh;

	font-family: system-ui;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: rgb(var(--text-black));

	margin: 0;
	padding: 0;
}
</style>
