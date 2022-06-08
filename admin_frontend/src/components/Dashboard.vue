<script setup lang="ts">
</script>

<template>
	<div class="dashboard">
		<div class="dashboard__cards dashcards">
			<div class="dashcard">
				<router-link :to="{ name: 'Index' }">Главная</router-link>
			</div>
			<div class="dashcard">
				<router-link :to="{ name: 'Lang', params: { lang: 'ru' }}" replace>Русская версия</router-link>
			</div>
			<div class="dashcard">
				<router-link :to="{ name: 'Lang', params: { lang: 'en' }}" replace>Английская версия</router-link>
			</div>
			<div class="dashcard">
				<router-link :to="{ name: 'Options' }">Настройки</router-link>
			</div>
		</div>
	</div>
</template>

<style scoped>
.dashboard__cards {
	display: grid;
	grid-gap: var(--offset);

	grid-template-columns: repeat(2, minmax(200px, 1fr));
	grid-template-rows: repeat(2, minmax(300px, 1fr));
}

.dashcard {
	display: flex;
	justify-content: center;
	align-items: center;

	border-radius: var(--offset_half);
	box-shadow: var(--neoShadow);

	/* background gradient */
	background-color: rgba(var(--light-orange), var(--alpha50));
	background-image: linear-gradient(to right, rgb(242, 112, 156), rgb(255, 148, 114));

}

.dashcard a {
	color: rgb(var(--white));
	font-size: 1.3rem;
	font-weight: bold;
	text-decoration: none;

	transition: var(--mainTransition);
}

.dashcard a:hover {
	color: rgb(var(--purple));
}

</style>