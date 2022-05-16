<template>
	<div class="login_wrapper">
		<LoginForm/>
	</div>
</template>

<script setup lang="ts">
import LoginForm from "../components/LoginForm.vue"
</script>

<style scoped lang="postcss">
.login_wrapper {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	height: 100vh;
}
</style>