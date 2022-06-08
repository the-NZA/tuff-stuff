<template>
	<nav class="navigation">
		<router-link :to="{ name: 'Index' }">Главная</router-link>
		<router-link :to="{ name: 'Lang', params: { lang: 'ru' }}" replace>Русская версия</router-link>
		<router-link :to="{ name: 'Lang', params: { lang: 'en' }}" replace>Английская версия</router-link>
		<router-link :to="{ name: 'Options' }">Настройки</router-link>
	</nav>
</template>

<script setup lang="ts">
</script>

<style scoped>
.navigation a:not(:last-child) {
	margin-right: var(--offset);
}

.navigation a {
	color: rgb(var(--gray));
	text-decoration: none;
	font-weight: 600;
	font-size: 1.1rem;

	transition: var(--mainTransition);
}

.navigation a.router-link-active {
	color: rgb(var(--black));
}

.navigation a:hover {
	color: rgb(var(--black));
}
</style>