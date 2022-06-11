<template>
	<div class="editPage" :key="route.path">
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
				<multi-text v-if="homepage.content.about_text"
				            v-model="homepage.content.about_text"
				            id="aboutText"
				            class="editPage__multi_text"
				></multi-text>
			</div>
			<!--About Text END-->

			<!--Shopping services-->
			<div class="editPage__row">
				<label for="shoppingServices" class="editPage__label">
					Заголовок блока услуг персонального шоппинга
				</label>
				<input
					v-model="homepage.content.shopping_title"
					@input="reset"
					type="text"
					class="editPage__input"
					id="shoppingServices"
					placeholder="Введите отображаемый заголовок">
			</div>
			<!--Shopping services END-->

			<!--Shopping cards-->
			<!--Shopping cards END-->

			<!--How it works-->
			<div class="editPage__row">
				<label for="howItWorks" class="editPage__label">
					Заголовок блока как работает сервис персонального шоппинга
				</label>
				<input
					v-model="homepage.content.how_works_title"
					@input="reset"
					type="text"
					class="editPage__input"
					id="howItWorks"
					placeholder="Введите отображаемый заголовок">
			</div>
			<!--How it works END-->

			<!--How works cards-->
			<!--How works cards END-->

			<!--Commission services-->
			<div class="editPage__row">
				<label for="commissionServices" class="editPage__label">
					Заголовок блока услуг комиссионного ресурса
				</label>
				<input
					v-model="homepage.content.commission_title"
					@input="reset"
					type="text"
					class="editPage__input"
					id="commissionServices"
					placeholder="Введите отображаемый заголовок">
			</div>
			<!--Commission services END-->

			<!--Commission cards-->
			<!--Commission cards END-->

			<!--Why us-->
			<div class="editPage__row">
				<label for="whyUs" class="editPage__label">
					Заголовок блока почему выбирают нас
				</label>
				<input
					v-model="homepage.content.why_us_title"
					@input="reset"
					type="text"
					class="editPage__input"
					id="whyUs"
					placeholder="Введите отображаемый заголовок">
			</div>
			<!--Why us END-->

			<!--Why us cards-->
			<!--Why us cards END-->

			<!--Image Grid-->
			<!--Image Grid END-->
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
import MultiText from './MultiText.vue';
import {onBeforeRouteUpdate, useRoute} from "vue-router";
import {computed, onBeforeMount, reactive, ref} from "vue";
import HTTP from "../util/HTTP";
import {AxiosError} from "axios";
import {Response} from "../types/Response";
import {Homepage, HomepageContent} from "../types/Homepage";
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

	// reset homepage
	homepage.content = <HomepageContent>{};

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