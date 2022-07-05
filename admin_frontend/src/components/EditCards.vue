<template>
	<div class="editCards">
		<div class="editCards__cards" :class="{ 'editCards__cards--full-width': props.fullWidth }">
			<template v-if="props.cards">
				<tuff-card v-for="c in props.cards" :card="c" @editCard="updateCard"
				           @deleteCard="handleDeleteCard"/>
			</template>
			<p v-else class="editCards__noCards">
				Пока не добавлено ни одной карточки
			</p>
		</div>
		<div class="editCards__footer">
			<button
				@click="addCard"
				class="editCards__createButton">
				Добавить
			</button>
		</div>

		<transition name="slide">
			<ModalCardEditor v-if="isShowModal"
			                 :card="editableCard"
			                 :card-type="props.cardType"
			                 @save="createCard"
			                 @update="handleEditCard"
			                 @close="resetModal"/>
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

const setModalState = (state: boolean) => {
	isShowModal.value = state;
};

const props = defineProps<{
	cards?: Card[],
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

// Set some functions visible as template refs
defineExpose({
	setModalState,
})

const addCard = () => {
	setModalState(true);
	editableCard.value = <Card>{
		id: "",
		title: "",
		content: "",
		lang: "",
	};
}

const createCard = () => {
	emits("createCard", {
		card: editableCard.value,
		cardType: props.cardType,
	});
}

const updateCard = (card: Card) => {
	setModalState(true);
	editableCard.value = {...card};
}

const resetModal = () => {
	setModalState(false);
	editableCard.value = <Card>{};
}

const handleEditCard = () => {
	emits("editCard", {
		card: editableCard.value,
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