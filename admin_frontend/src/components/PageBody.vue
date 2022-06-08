<template>
	<div class="pageBody">
		<slot></slot>
	</div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.pageBody {
	background-color: rgb(var(--white));
	padding: var(--offset);
	border-radius: var(--offset_half);
	box-shadow: var(--neoShadow);
}
</style>