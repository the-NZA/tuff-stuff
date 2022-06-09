<template>
	<div class="settings">
		<Loader v-if="isLoading"/>

		<h2 class="settings__title">Настройки сайта</h2>

		<div v-for="opt in options" class="settings__row">
			<label :for="opt.name + `Setting`" class="settings__label">{{ opt.title }}</label>

			<p class="settings__snippet">
				{{ helperText[opt.name].snippet }}
			</p>

			<input v-model="opt.value"
			       @input="reset"
			       :id="opt.name + `Setting`"
			       type="text"
			       class="settings__input"
			       placeholder="Название">
		</div>

		<div class="settings__footer">
			<button
				@click="save"
				class="settings__save">Сохранить
			</button>

			<div v-if="isError" class="settings__error">
				{{ message }}
			</div>

			<div v-if="isSuccess" class="settings__success">
				{{ message }}
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import Loader from './Loader.vue';
import {onBeforeMount} from "vue";
import HTTP from "../util/HTTP";
import {Option} from "../types/Option";
import {Response} from "../types/response";
import {reactive, ref} from "vue";

const helperText = {
	email: {
		snippet: "Введите адрес электронной почты",
		placeholder: "test@test.com"
	},
	phone: {
		snippet: "Номер телефона для обратной связи",
		placeholder: "+79999999999"
	},
	telegram: {
		snippet: "Ник телеграм для обратной связи без @",
		placeholder: "test_nickname"
	},
	whatsapp: {
		snippet: "Номер WhatsApp для обратной связи",
		placeholder: "+79999999999"
	}
};

const options = reactive<Option[]>([]);
const message = ref("");

// State flags
const isLoading = ref(true);
const isError = ref(false);
const isSuccess = ref(false);

// Reset flags and message
const reset = () => {
	isError.value = false;
	isSuccess.value = false;
	message.value = "";
}

// Load data from server
onBeforeMount(async () => {
	try {
		const res = await HTTP.get<Response<Option[]>>("/api/v1/option")
		options.splice(0, options.length, ...res.data.data);
	} catch (err) {
		isError.value = true;
		message.value = "Не удалось данные с сервера";
	}

	isLoading.value = false;
});

// Save data to server
const save = async () => {
	isLoading.value = true;
	isError.value = false;
	isSuccess.value = false;

	try {
		// JSON.stringify() used for sending data to server
		// because axios add "data" property to request body,
		// and it is not allowed in PUT request handler in server
		await HTTP.put<Response<Option[]>>("/api/v1/option", JSON.stringify(options));

		isSuccess.value = true;
		message.value = "Настройки сохранены";
	} catch (err) {
		isError.value = true;
		message.value = "Не удалось сохранить настройки";
	}

	isLoading.value = false;
}

</script>

<style scoped>

.settings {
	display: grid;
	gap: var(--offset);
}

.settings__title {
	margin: 0;
	text-align: center;
}

.settings__row {
	padding: var(--offset_half);
	background-color: rgb(var(--gray-light));

	border: 1px solid rgba(var(--gray), var(--alpha25));
	border-radius: var(--offset_quarter);
}

.settings__label {
	font-size: 1.1rem;
	font-weight: bold;

	color: rgb(var(--black));

	display: block;

	margin-bottom: var(--offset_quarter);
}

.settings__snippet {
	margin: 0 0 var(--offset_half) 0;

	color: rgb(var(--gray));

	font-size: 0.8rem;
	font-weight: 500;
}

.settings__input {
	width: 100%;
	max-width: 50%;

	padding: var(--offset_quarter);
	border: 1px solid rgba(var(--gray), var(--alpha25));
	border-radius: var(--offset_eight);
}

.settings__footer {
	display: grid;
	grid-template-columns: max-content 1fr;
	align-items: start;
	gap: var(--offset);
}

.settings__save {
	min-width: 160px;

	padding: var(--offset_half);

	border: 2px solid rgb(var(--purple));
	border-radius: var(--offset_eight);

	cursor: pointer;
	font-weight: bolder;

	color: rgb(var(--purple));
	background-color: rgb(var(--white));

	transition: var(--mainTransition);
}

.settings__save:hover {
	background-color: rgb(var(--purple));
	color: rgb(var(--white));
}

.settings__error {
	padding: var(--offset_half);
	background-color: rgb(var(--red));
	color: rgb(var(--white));
	font-weight: bold;
	border-radius: var(--offset_eight);
	border: 2px solid rgb(var(--red));
}

.settings__success {
	padding: var(--offset_half);

	background-color: rgb(var(--light-orange));
	color: rgb(var(--white));

	font-weight: bold;
	border-radius: var(--offset_eight);
	border: 2px solid rgb(var(--light-orange));
	text-shadow: 0 0 1px rgba(var(--gray), var(--alpha50));
}
</style>
