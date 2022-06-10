<template>
	<div class="editPage">
		<Loader v-if="isLoading"/>

		<h2 class="editPage__title">Настройки сайта</h2>

		<div v-for="opt in options" class="editPage__row">
			<label :for="opt.name + `Setting`" class="editPage__label">{{ opt.title }}</label>

			<p class="editPage__snippet">
				{{ helperText[opt.name].snippet }}
			</p>

			<input v-model="opt.value"
			       @input="reset"
			       :id="opt.name + `Setting`"
			       type="text"
			       class="editPage__input"
			       placeholder="Название">
		</div>

		<div class="editPage__footer">
			<button
				@click="save"
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
import {onBeforeMount} from "vue";
import HTTP from "../util/HTTP";
import {Option} from "../types/Option";
import {Response} from "../types/Response";
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
		message.value = "Не удалось загрузить данные с сервера";
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

<style lang="postcss" scoped>
</style>
