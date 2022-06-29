<template>
	<div class="editCards">
		<div class="editCards__cards" :class="{ 'editCards__cards--full-width': props.fullWidth }">
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

		<transition>
			<ModalCardEditor v-if="isShowModal" :card="editableCard" @close="resetModal"/>
		</transition>
	</div>
</template>

<script setup lang="ts">
import TuffCard from "./TuffCard.vue";
import ModalCardEditor from "./ModalCardEditor.vue";
import {Card} from "../types/Card";
import {CardType} from "../types/CardType";
import {ref} from "vue";

const isShowModal = ref(false);
const editableCard = ref<Card>(<Card>{});

const props = defineProps<{
	cards: Card[],
	fullWidth?: boolean,
	cardType: CardType,
}>()

const emits = defineEmits<{
	(e: "createCard", value: {
		card: Card,
		cardType: CardType,
	}): void;
	(e: "editCard", value: {
		card: Card,
		cardType: CardType,
	}): void;
	(e: "deleteCard", value: {
		cardID: String,
		cardType: CardType,
	}): void;
}>()

const createCard = () => {
	isShowModal.value = true;

	// emits("createCard", {
	// 	card: {
	// 		id: "",
	// 		title: "",
	// 		content: "",
	// 		lang: "",
	// 	},
	// 	cardType: props.cardType,
	// });
}

const resetModal = () => {
	isShowModal.value = false;
	editableCard.value = <Card>{};
}

const handleEditCard = (card: Card) => {
	emits("editCard", {
		card: card,
		cardType: props.cardType,
	});
}

const handleDeleteCard = (cardID: String) => {
	emits("deleteCard", {
		cardID: cardID,
		cardType: props.cardType,
	});
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

.editCards__cards--full-width {
	grid-template-columns: 1fr;
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