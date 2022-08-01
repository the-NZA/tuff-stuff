<template>
	<div class="gridCard">
		<button @click="handleEdit" class="gridCard__edit">Редактировать</button>

		<div class="gridCard__image">
			<img :src="imgURL" alt="">
		</div>
	</div>
</template>

<script setup lang="ts">
import {GridImageWithURL} from "../types/Image";
import {computed} from "vue";

const props = defineProps<{
	image: GridImageWithURL
}>()

const emits = defineEmits<{
	(e: "edit", value: GridImageWithURL
	): void;
}>()

const imgURL = computed(() => {
	return `http://tuff-stuff.local/${props.image.url}`;
})

const handleEdit = () => {
	emits("edit", props.image);
}
</script>

<style lang="postcss" scoped>

.gridCard {
	position: relative;
}

.gridCard__edit {
	position: absolute;
	bottom: var(--offset_quarter);
	left: var(--offset_quarter);

	font-size: 0.8rem;
	cursor: pointer;
	padding: var(--offset_quarter);
	background-color: rgb(var(--gray-light));
	border: 2px solid rgb(var(--light-orange));
	border-radius: var(--offset_eight);
	color: rgb(var(--light-orange));
	font-weight: 700;
	transition: var(--mainTransition);
}

.gridCard__edit:hover {
	background-color: rgb(var(--light-orange));
	color: rgb(var(--white));
}

.gridCard__image {
	background-color: rgb(var(--light-orange));
	max-width: 100%;
	max-height: 100%;
	min-height: 345px;
}

.gridCard__image img {
	display: block;
	width: 350px;
	height: 350px;

	object-fit: cover;
}

</style>