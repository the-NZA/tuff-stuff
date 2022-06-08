<script setup lang="ts">
import Navigation from "./Navigation.vue"
</script>

<template>
	<div class="page">
		<header class="page__header">
			<div class="page__logo">
				<h1 class="page__title">TuffStuff</h1>
			</div>

			<div class="page__nav">
				<Navigation/>
			</div>
		</header>
		<div class="page__body">
			<div class="wrapper">
				<slot></slot>
			</div>
		</div>
	</div>
</template>

<style scoped>
.page {
	display: grid;
	grid-template-rows: min-content 1fr;
	min-height: 100vh;
	gap: var(--offset);
}

.page__header {
	width: 100%;
	max-width: var(--siteWidth);

	margin: 0 auto;
	padding: var(--offset_half) 0;

	border-bottom: 1px solid rgb(var(--purple));

	display: grid;
	grid-template-columns: min-content 1fr;

	align-items: center;
	gap: var(--offset_twice);
}

.page__title {
	margin: 0;
	text-transform: uppercase;
	font-weight: 800;
}

.page__nav {
	padding: var(--offset_half);
}

.page__body {
	padding: var(--offset);
}

.wrapper {
	max-width: var(--siteWidth);
	margin: 0 auto;
}
</style>