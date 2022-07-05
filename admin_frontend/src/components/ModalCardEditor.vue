<template>
	<div class="modalEditor">

		<div class="modalEditor__wrapper">
			<div class="modalEditor__header">
				<h2 class="modalEditor__title">Редактировать карточку {{ NameByCardType(props.cardType) }}</h2>

				<button @click="close"
				        class="modalEditor__close">
					<span></span>
				</button>
			</div>

			<div class="modalEditor__fields">
				<div class="editPage__row">
					<label for="modalCardTitle" class="editPage__label">Название карточки</label>

					<input
						v-model="props.card.title"
						@input="msgStore.Reset"
						type="text"
						class="editPage__input"
						id="modalCardTitle"
						placeholder="Введите название">
				</div>

				<div class="editPage__row">
					<label for="modalCardContent" class="editPage__label">Текст карточки</label>
					<multi-text v-model="editableContent"
					            :divider="'\r\n'"
					            id="modalCardContent"
					            class="editPage__multi_text"
					></multi-text>
				</div>
			</div>

			<div class="modalEditor__footer">
				<button
					@click="save"
					class="modalEditor__save">
					Сохранить
				</button>

				<div v-if="msgStore.IsModalError" class="modalEditor__error">
					{{ msgStore.ModalErrorMessage }}
				</div>

				<div v-if="msgStore.IsModalSuccess" class="modalEditor__success">
					{{ msgStore.ModalSuccessMessage }}
				</div>
			</div>
		</div>

	</div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useMessageStore} from "../store/message";
import MultiText from "./MultiText.vue";
import {Card} from "../types/Card";
import {CardType} from "../types/CardType";
import NameByCardType from "../util/NameByCardType";

const msgStore = useMessageStore();

const emits = defineEmits<{
	(e: "close"): void,
	(e: "save"): void,
}>()

const props = defineProps<{
	card: Card,
	cardType: CardType,
}>()

const editableContent = computed({
	get: () => {
		return props.card.content.split("\r\n")
	},
	set: (value: string[]) => {
		props.card.content = value.join("\r\n");
	},
});

const close = () => {
	msgStore.Reset();

	emits("close");
}

const save = () => {
	msgStore.Reset();

	emits("save");
}

</script>

<style scoped>
.modalEditor {
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;

	background-color: rgba(var(--gray), var(--alpha50));

	display: flex;
	align-items: center;
	justify-content: center;
}

.modalEditor__wrapper {
	background-color: rgb(var(--white));

	max-width: var(--siteWidth);
	width: 100%;

	margin: 0 auto;

	border-radius: var(--offset_half);
	padding: var(--offset);

	display: grid;
	grid-template-rows: min-content 1fr min-content;
	gap: var(--offset);
}

.modalEditor__header {
	display: flex;
	align-items: center;
	justify-content: space-between;
}

.modalEditor__title {
	margin: 0;
	text-transform: uppercase;
}

.modalEditor__close {
	background-color: rgb(var(--gray-light));
	border: 2px solid rgb(var(--red));
	border-radius: 50%;

	height: var(--offset);
	width: var(--offset);

	position: relative;

	cursor: pointer;

	transition: var(--mainTransition);
}

.modalEditor__close:hover {
	background-color: rgb(var(--red));
	color: rgb(var(--white));
}

.modalEditor__close:hover span::before,
.modalEditor__close:hover span::after {
	background-color: rgb(var(--white));
	color: rgb(var(--red));
}

.modalEditor__close span {
	display: block;
}

.modalEditor__close span::before {
	content: "";
	display: block;

	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%) rotate(45deg);

	width: var(--offset_half);
	height: 2px;

	background-color: rgb(var(--red));
}

.modalEditor__close span::after {
	content: "";
	display: block;

	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%) rotate(-45deg);

	width: var(--offset_half);
	height: 2px;

	background-color: rgb(var(--red));
}

.modalEditor__fields .editPage__row:first-of-type {
	margin-bottom: var(--offset);
}

.modalEditor__footer {
	display: grid;
	grid-template-columns: max-content 1fr;
	align-items: center;
	gap: var(--offset_twice);
}

.modalEditor__save {
	background-color: rgb(var(--white));
	color: rgb(var(--purple));

	border: 2px solid rgb(var(--purple));

	border-radius: var(--offset_quarter);
	padding: var(--offset_half);

	font-weight: bold;

	cursor: pointer;
	transition: var(--mainTransition);
}

.modalEditor__save:hover {
	background-color: rgb(var(--purple));
	color: rgb(var(--white));
}

.modalEditor__error,
.modalEditor__success {
	font-weight: bold;
	border: 2px solid rgb(var(--gray));
	border-radius: var(--offset_quarter);

	padding: var(--offset_half);
}

.modalEditor__error {
	border-color: rgb(var(--red));
	background-color: rgb(var(--red));
	color: rgb(var(--white));
}

.modalEditor__success {
	border-color: rgb(var(--light-orange));
	background-color: rgb(var(--light-orange));
	color: rgb(var(--white));
}
</style>