<template>
	<div class="imgEdtr">
		<grid-card v-for="img in props.images" :image="img" @edit="editGridCard"/>
	</div>


	<transition name="slide">
		<ModalImageEditor v-if="msgStore.IsShowGridModal"
		                  :image="editableImage"
		                  @update="updateGridImage"
		                  @close="resetModal"/>
	</transition>
</template>

<script setup lang="ts">
import {GridImage, GridImageWithURL, Image} from "../types/Image";
import GridCard from "./GridCard.vue";
import ModalImageEditor from "./ModalImageEditor.vue";
import {useMessageStore} from "../store/message";
import {ref} from "vue";
import HTTP from "../util/HTTP";
import {Response} from "../types/Response";
import {Delay} from "../util/delay";

const msgStore = useMessageStore();

const editableImage = ref<GridImageWithURL>(<GridImageWithURL>{});

const props = defineProps<{
	images: GridImageWithURL[];
}>()

const editGridCard = (image: GridImageWithURL) => {
	editableImage.value = image;
	msgStore.SetShowGridModal(true);
}

const resetModal = () => {
	console.log("resetModal");
	msgStore.SetShowGridModal(false);
}

const updateGridImage = (imageFile: File) => {
	uploadImage(imageFile);
}

const uploadImage = async (imageFile: File) => {
	const formData = new FormData();
	formData.append("tuff_image", imageFile);

	try {
		const imgRes = await HTTP.post<Response<Image>>(`/api/v1/image`, formData, {
			headers: {
				"Content-Type": "multipart/form-data",
			},
		});

		const gridRes = await HTTP.put<Response<GridImage>>(`/api/v1/image-grid/${editableImage.value.id}`, {
			id: editableImage.value.id,
			image_id: imgRes.data.data.id,
		} as GridImage);

		editableImage.value.url = imgRes.data.data.url;

		msgStore.SetModalSuccess("Изображение успешно обновлено");

		Delay(() => {
			msgStore.Reset()
			msgStore.SetShowGridModal(false);
		}, 1000);

	} catch (e: any) {
		// If axios error
		if (e.response) {
			console.log(e.response.data);
			msgStore.SetModalError(e.response.data.error);
		} else {
			console.log(e);
			msgStore.SetModalError("Не удалось загрузить изображение");
		}
	}
}
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