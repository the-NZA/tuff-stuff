<template>
	<div class="imgEdtr">
		<grid-card v-for="img in props.images" :image="img"/>
	</div>
</template>

<script setup lang="ts">
import {GridImage} from "../types/Image";
import GridCard from "./GridCard.vue";

const props = defineProps<{
	images: GridImage[];
}>()
</script>

<style lang="postcss" scoped>
.imgEdtr {
	display: grid;
	grid-template-columns: repeat(3, 350px);
	grid-template-rows: repeat(3, 350px);
	gap: var(--offset_half);

	align-items: center;
	justify-content: center;
}
</style>