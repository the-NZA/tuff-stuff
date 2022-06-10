<template>
	<div class="editPage">
		<loader v-if="isLoading || !isDataLoaded"/>

		<h2 class="editPage__title">Редактирование {{ currentLang }} языка</h2>

		<template v-if="isDataLoaded">

			<!--About Title-->
			<div class="editPage__row">
				<label for="aboutTitle" class="editPage__label">Заголовок блока о компании</label>

				<input
					v-model="homepage.content.about_title"
					@input="reset"
					type="text"
					class="editPage__input"
					id="aboutTitle"
					placeholder="Введите отображаемое название">
			</div>
			<!--About Title END-->

			<!--About Text-->
			<div class="editPage__row">
				<label for="aboutText" class="editPage__label">Текст блока о компании</label>
				<multi-text v-model="homepage.content.about_text"
				            id="aboutText"
				            class="editPage__multi_text"
				></multi-text>
			</div>
			<!--About Text END-->
		</template>


		<div class="editPage__footer">
			<button
				@click="saveHomepage"
				class="editPage__save">Сохранить
			</button>

			<div v-if="isError" class="editPage__error">
				{{ message }}
			</div>

			<div v-if="isSuccess" class="editPage__success">
				{{ message }}
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import Loader from './Loader.vue';
import MultiLineText from './MultiLineText.vue';
import MultiText from './MultiText.vue';
import {onBeforeRouteUpdate, useRoute} from "vue-router";
import {computed, onBeforeMount, reactive, ref} from "vue";
import HTTP from "../util/HTTP";
import {AxiosError} from "axios";
import {Response} from "../types/Response";
import {Homepage} from "../types/Homepage";
import {Delay} from "../util/delay";

// State flags
const isLoading = ref(false);
const isDataLoaded = ref(false);
const isError = ref(false);
const isSuccess = ref(false);

// Reset flags and message
const reset = () => {
	isError.value = false;
	isSuccess.value = false;
	message.value = "";
}

const route = useRoute();
const lang = ref<string>(route.params.lang as string);
const message = ref("");
const currentLang = computed<string>(() => lang.value === "ru" ? "русского" : "английского");

const homepage = reactive<Homepage>(<Homepage>{})

// loadData is called before the component is mounted or updated
const loadData = async () => {
	isLoading.value = true;
	isError.value = false;
	isSuccess.value = false;
	message.value = "";

	try {
		// Get data from server
		const res = await HTTP.get<Response<Homepage>>(`/api/v1/homepage`, {
			params: {
				lang: lang.value
			}
		});

		homepage.id = res.data.data.id;
		homepage.lang = res.data.data.lang;
		homepage.content = res.data.data.content;
	} catch (e) {
		isError.value = true;
		message.value = (e as AxiosError).message;
		console.log(e)
	}

	isDataLoaded.value = true;

	Delay(() => {
		isLoading.value = false;
	});
}

// Load data on component first mount
onBeforeMount(async () => {
	await loadData();
});

// Update the lang ref when the route changes
// get data from backend
onBeforeRouteUpdate(async (to, _) => {
	lang.value = to.params.lang as string;

	await loadData();
});

// Save homepage data to server
const saveHomepage = async () => {
	console.log("saved");
}

</script>

<style lang="postcss" scoped>

</style>