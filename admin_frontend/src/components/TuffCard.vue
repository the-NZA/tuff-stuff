<template>
	<div class="card">
		<h3 class="card__title">{{ props.card.title }}</h3>

		<div class="card__content">
			<p v-for="line in prepareContent()" :key="line">
				{{ line }}
			</p>
		</div>

		<div class="card__footer">
			<button class="card__button card__button--edit" @click="editCard">Редактировать</button>
			<button class="card__button card__button--delete" @click="deleteCard">Удалить</button>
		</div>
	</div>
</template>

<script setup lang="ts">
import {Card} from "../types/Card";
import {computed} from "vue";

const props = defineProps<{
	card: Card
}>()

const emits = defineEmits<{
	(e: "deleteCard", value: String): void;
	(e: "editCard", value: Card): void;
}>()

const editCard = (e: Event) => {
	emits("editCard", props.card);
}

const deleteCard = (e: Event) => {
	emits("deleteCard", props.card.id);
}

const prepareContent = () => {
	return props.card.content.split("\r\n")
}
</script>

<style lang="postcss" scoped>
.card {
	border: 1px solid rgba(var(--gray), var(--alpha25));
	border-radius: var(--offset_eight);
	padding: var(--offset_half);

	line-height: 1.5;

	text-transform: uppercase;

	display: grid;
	grid-template-rows: min-content 1fr min-content;
	gap: var(--offset_half);
}

.card__title {
	font-size: 1.2rem;
	font-weight: 700;

	margin: 0;
}

.card__content p {
	margin: 0 0 var(--offset_half) 0;
}

.card__content p:last-child {
	margin: 0;
}

.card__footer {
}

.card__button {
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

.card__button:not(:last-child) {
	margin-right: var(--offset_half);
}

.card__button--delete {
	border-color: rgb(var(--red));
	color: rgb(var(--red));
}

.card__button--delete:hover {
	background-color: rgb(var(--red));
	color: rgb(var(--white));
}

.card__button--edit:hover {
	background-color: rgb(var(--light-orange));
	color: rgb(var(--white));
}
</style>
