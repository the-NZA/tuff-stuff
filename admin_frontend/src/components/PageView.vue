<script setup lang="ts">
import Navigation from "./Navigation.vue"
</script>

<template>
	<div class="page">
		<div class="page__nav">
			<Navigation/>
		</div>
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

.page__nav {
	padding: var(--offset);
}

.page__body {
	padding: var(--offset);
}

.wrapper {
	max-width: var(--siteWidth);
	margin: 0 auto;
}
</style>