<template>
	<div class="modalEditor">

		<div class="modalEditor__wrapper">
			<div class="modalEditor__header">
				<h2 class="modalEditor__title">
					{{ editorTitle }}
				</h2>

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
					@click="handleSaveButton"
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
	(e: "update"): void,
}>()

const props = defineProps<{
	card: Card,
	cardType: CardType,
}>()

const editorTitle = computed(() => {
	return (props.card.id ? "Редактирование карточки " : "Создание карточки ") + NameByCardType(props.cardType);
});

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

const update = () => {
	msgStore.Reset();

	emits("update");
}

const handleSaveButton = () => {
	if (props.card.id === "") {
		save()
		return
	}

	update();
}

</script>

<style scoped>
</style>