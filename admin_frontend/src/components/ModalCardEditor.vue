<template>
	<div class="modalEditor">

		<div class="modalEditor__wrapper">
			<div class="modalEditor__header">
				<h2 class="modalEditor__title">Редактировать карточку</h2>

				<button @click="close"
				        class="modalEditor__close">
					Закрыть
				</button>
			</div>

			<div class="modalEditor__fields">
				<div class="editPage__row">
					<label for="aboutTitle" class="editPage__label">Заголовок блока о компании</label>

					<!--					v-model="homepage.content.about_title"-->
					<!--					@input="reset"-->

					<input
						type="text"
						class="editPage__input"
						id="aboutTitle"
						placeholder="Введите отображаемое название">
				</div>
			</div>

			<div class="modalEditor__footer">
				<button
					@click="save"
					class="modalEditor__save">Сохранить
				</button>
			</div>
		</div>

	</div>
</template>

<script setup lang="ts">
import {Card} from "../types/Card";
import {onBeforeMount, ref} from "vue";

const editableCard = ref<Card>(<Card>{});

const emits = defineEmits<{
	(e: "close"): void;
}>()

const props = defineProps<{
	card: Card,
}>()

const close = () => {
	emits("close");
}

const save = () => {
	console.log("saved");
}

onBeforeMount(() => {
	// copy card to local editableCard
	editableCard.value.id = props.card.id;
	editableCard.value.title = props.card.title;
	editableCard.value.content = props.card.content;
	editableCard.value.lang = props.card.lang;
})

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

</style>