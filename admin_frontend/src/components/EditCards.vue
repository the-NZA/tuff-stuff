<template>
	<div class="editCards">
		<div class="editCards__cards">
			<template v-if="props.cards">
				<tuff-card v-for="c in props.cards" :card="c" @editCard="handleEditCard"
				           @deleteCard="handleDeleteCard"/>
			</template>
			<p v-else class="editCards__noCards">
				Пока не добавлено ни одной карточки
			</p>
		</div>
		<div class="editCards__footer">
			<button
				@click="createCard"
				class="editCards__createButton">
				Добавить
			</button>
		</div>
	</div>
</template>

<script setup lang="ts">
import TuffCard from "./TuffCard.vue";
import {Card} from "../types/Card";
import {ref} from "vue";

const props = defineProps<{
	cards: Card[]
}>()

const isEdit = ref(false);

const createCard = () => {
	console.log("createCard");
}

const handleEditCard = (card: Card) => {
	isEdit.value = true;

	console.log("edit", card);
}

const handleDeleteCard = (cardID: String) => {

	console.log("delete", cardID);
}
</script>

<style lang="postcss" scoped>
.editCards {
	display: grid;
	grid-template-rows: 1fr max-content;
	gap: var(--offset);
}

.editCards__cards {
	display: grid;
	grid-template-columns: repeat(2, 1fr);
	gap: var(--offset_half);
}

.editCards__createButton {
	cursor: pointer;
	padding: var(--offset_quarter);

	background-color: rgb(var(--white));
	color: rgb(var(--purple));

	border: 2px solid rgb(var(--purple));
	border-radius: var(--offset_eight);

	font-weight: bold;

	text-shadow: none;
	transition: var(--mainTransition);
}

.editCards__createButton:hover {
	background-color: rgb(var(--purple));
	color: rgb(var(--white));
	text-shadow: 0 0 1px rgba(var(--gray), var(--alpha50));
}

.editCards__noCards {
	margin: 0;
}
</style>