<template>
	<div class="modalEditor">
		<div class="modalEditor__wrapper">
			<div class="modalEditor__header">
				<h2 class="modalEditor__title">Редактирование изображения</h2>

				<button @click="close"
				        class="modalEditor__close">
					<span></span>
				</button>
			</div>

			<div class="modalEditor__fields">
				<div class="editPage__row">
					<div class="editPage__img">
						<img :src="preparedURL" alt="">
					</div>

					<label for="modalImageChooser" class="editPage__label">Изменить</label>

					<input
						@change="handleImageChange"
						type="file"
						class="editPage__input modalEditor__image_chooser"
						id="modalImageChooser"
						placeholder="Введите название">
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
import {GridImageWithURL} from "../types/Image";
import {useMessageStore} from "../store/message";
import {computed, ref} from "vue";

const msgStore = useMessageStore();
const newImageFile = ref<File>();
const newImageURL = ref<string>();

const props = defineProps<{
	image: GridImageWithURL
}>()

const emits = defineEmits<{
	(e: "close"): void,
	(e: "update", value: File): void,
}>()

const handleSaveButton = () => {
	if (!newImageFile.value) {
		emits("close");
		return
	}

	emits("update", newImageFile.value);
}

const close = () => {
	msgStore.Reset();
	emits("close");
}

const preparedURL = computed(() => {
	if (newImageURL.value) {
		return newImageURL.value;
	}

	if (props.image.url.startsWith("uploads/")) {
		return `http://tuff-stuff.local/${props.image.url}`;
	}

	return newImageURL.value;
})

const handleImageChange = (e: Event) => {
	const target = e.target as HTMLInputElement;
	if (!target.files) {
		msgStore.SetModalError("Не удалось загрузить изображение");
		return;
	}

	newImageFile.value = target.files[0];

	const reader = new FileReader();
	reader.onload = (e: ProgressEvent) => {
		const dataURL = (e.target as FileReader).result;
		newImageURL.value = dataURL as string;
	};
	reader.readAsDataURL(newImageFile.value as Blob);
}
</script>

<style scoped lang="postcss">
.editPage__row {
	margin: 0 !important;
}

.editPage__label {
	margin: 0;

	cursor: pointer;

	padding: var(--offset_quarter);
	border: 2px solid rgb(var(--light-orange));
	border-radius: var(--offset_eight);
	max-width: min-content;

	transition: var(--mainTransition);
}

.editPage__label:hover {
	background-color: rgb(var(--light-orange));
	color: rgb(var(--white));
}

.modalEditor__image_chooser {
	display: none;
}
</style>