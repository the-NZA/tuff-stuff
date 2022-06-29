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
				            :divider="'\n'"
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
			<div class="editPage__row">
				<label for="shoppingCards" class="editPage__label">
					Карточки услуг персонального шоппинга
				</label>

				<div class="editPage__cards">
					<!-- Edit Cards Component -->
					<edit-cards :cards="shoppingCards"
					            @editCard="handleEditCard"
					            @deleteCard="handleDeleteCard"
					            :card-type="CardType.Shopping"
					/>
				</div>
			</div>
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
			<div class="editPage__row">
				<label for="howItWorksCards" class="editPage__label">
					Карточки как работает сервис персонального шоппинга
				</label>

				<div class="editPage__cards">
					<!-- Edit Cards Component -->
					<edit-cards :cards="howItWorksCards" :card-type="CardType.HowWorks"/>
				</div>
			</div>
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
			<div class="editPage__row">
				<label for="commissionCards" class="editPage__label">
					Карточки услуг комиссионного ресурса
				</label>

				<div class="editPage__cards">
					<!-- Edit Cards Component -->
					<edit-cards :cards="commissionCards" :full-width="true" :card-type="CardType.Commission"/>
				</div>
			</div>
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
			<div class="editPage__row">
				<label for="whyUsCards" class="editPage__label">
					Карточки почему выбирают нас
				</label>

				<div class="editPage__cards">
					<!-- Edit Cards Component -->
					<edit-cards :cards="whyUsCards" :card-type="CardType.WhyUs"/>
				</div>
			</div>
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
import EditCards from './EditCards.vue';
import {onBeforeRouteUpdate, useRoute, useRouter} from "vue-router";
import {computed, onBeforeMount, reactive, ref} from "vue";
import HTTP from "../util/HTTP";
import {AxiosError} from "axios";
import {Response} from "../types/Response";
import {Card} from "../types/Card";
import {Homepage, HomepageContent} from "../types/Homepage";
import {CardType} from "../types/CardType";
import {Delay} from "../util/delay";
import SlugByCardType from "../util/SlugByCardType";

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

const router = useRouter();
const route = useRoute();
const lang = ref<string>(route.params.lang as string);
const message = ref("");
const currentLang = computed<string>(() => lang.value === "ru" ? "русского" : "английского");

const homepage = reactive<Homepage>(<Homepage>{})
const shoppingCards = ref<Card[]>(<Card[]>[])
const howItWorksCards = ref<Card[]>(<Card[]>[])
const commissionCards = ref<Card[]>(<Card[]>[])
const whyUsCards = ref<Card[]>(<Card[]>[])

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

		// Get shopping cards from server
		const shopRes = await HTTP.get <Response<Card[]>>(`/api/v1/shopping-card`, {
			params: {
				lang: lang.value
			}
		});
		shoppingCards.value = shopRes.data.data;

		// Get how it works cards from server
		const howItWorksRes = await HTTP.get <Response<Card[]>>(`/api/v1/how-works-card`, {
			params: {
				lang: lang.value
			}
		});
		howItWorksCards.value = howItWorksRes.data.data;

		// Get commission cards from server
		const commissionRes = await HTTP.get <Response<Card[]>>(`/api/v1/commission-card`, {
			params: {
				lang: lang.value
			}
		});
		commissionCards.value = commissionRes.data.data;

		// Get why us cards from server
		const whyUsRes = await HTTP.get <Response<Card[]>>(`/api/v1/why-us-card`, {
			params: {
				lang: lang.value
			}
		});
		whyUsCards.value = whyUsRes.data.data;

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
	isLoading.value = true;
	reset();

	// Remove all empty lines in about text array
	const preparedText = homepage.content.about_text.filter(line => line.trim().length > 0);

	// Check if homepage content's fields are empty
	if (
		homepage.content.about_title === "" ||
		preparedText.length === 0 ||
		homepage.content.shopping_title === "" ||
		homepage.content.how_works_title === "" ||
		homepage.content.commission_title === "" ||
		homepage.content.why_us_title === ""
	) {
		isError.value = true;
		isLoading.value = false;
		message.value = "Заполните все поля";

		return
	}

	try {
		await HTTP.put<Response<Homepage>>(`/api/v1/homepage/${homepage.id}`, homepage);

		isSuccess.value = true;
		message.value = "Данные успешно сохранены";

		// Reload page
		Delay(() => {
			router.go(0)
		}, 450);

	} catch (e) {
		isError.value = true;
		if ((e as AxiosError).response!.status >= 400 && (e as AxiosError).response!.status < 500) {
			message.value = "Указаны некорректные данные";
		} else {
			message.value = (e as AxiosError).message;
		}
		message.value = (e as AxiosError).message;
		console.log((e as AxiosError).response);
	}

	Delay(() => {
		isLoading.value = false;
	});
}

// handleEditCard is called when the user clicks the edit button on each card
const handleEditCard = async (val: any) => {
	const card = val.card;
	const cardType = val.cardType;

	isLoading.value = true;
	reset();

	try {
		const res = await HTTP.get<Response<Card>>(`/api/v1/${SlugByCardType(cardType)}/${card.id}`);

		console.log(res.data)

		// Open modal with card data

		isLoading.value = false;

	} catch (e) {
		isError.value = true;
		message.value = (e as AxiosError).message;
		console.log(e)
	}

}

// handleDeleteCard is called when the user clicks the delete button on each card
const handleDeleteCard = async (val: any) => {
	const cardID = val.cardID;
	const cardType = val.cardType;

	// Give the user a chance to cancel the action
	if (!confirm("Вы действительно хотите удалить эту карточку?")) return;

	isLoading.value = true;
	reset();

	try {
		await HTTP.delete<Response<Card>>(`/api/v1/${SlugByCardType(cardType)}/${cardID}`);

		isSuccess.value = true;
		message.value = "Данные успешно удалены";

		// Reload page
		Delay(() => {
			router.go(0)
		}, 450);

	} catch (e) {
		isError.value = true;
		console.log(e)
		if ((e as AxiosError).response!.status >= 400 && (e as AxiosError).response!.status < 500) {
			message.value = "Указаны некорректные данные";
		} else {
			message.value = (e as AxiosError).message;
		}
		message.value = (e as AxiosError).message;
		console.log((e as AxiosError).response);
	}

	Delay(() => {
		isLoading.value = false;
	});
}
</script>

<style lang="postcss" scoped>

</style>