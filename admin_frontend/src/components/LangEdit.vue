<template>
	<div class="editPage">
		<loader v-if="isLoading"/>

		<h2 class="editPage__title">Редактирование {{ currentLang }} языка</h2>

		<!-- There is rows -->

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
import {onBeforeRouteUpdate, useRoute} from "vue-router";
import {computed, onBeforeMount, reactive, ref} from "vue";
import HTTP from "../util/HTTP";
import {AxiosError} from "axios";
import {Response} from "../types/Response";
import {Homepage} from "../types/Homepage";

// State flags
const isLoading = ref(false);
const isError = ref(false);
const isSuccess = ref(false);

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

	isLoading.value = false;
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