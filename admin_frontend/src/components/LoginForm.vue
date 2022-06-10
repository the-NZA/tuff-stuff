<script setup lang="ts">
import {Response} from "../types/Response";
import HTTP from "../util/HTTP"
import {ref} from "vue"
import axios, {AxiosError} from "axios";
import {useRouter} from "vue-router";
import {StatusCodes as code} from "http-status-codes";
import Loader from "../components/Loader.vue";

const router = useRouter()
const username = ref("")
const password = ref("")
const loginError = ref("")
const isLoading = ref(false)

const resetError = () => {
	loginError.value = ""
}

const login = async () => {
	loginError.value = ""
	isLoading.value = true

	if (username.value === "" || password.value === "") {
		loginError.value = "Заполните все поля"
		isLoading.value = false
		return
	}

	try {
		const res = await HTTP.post<Response<{ login: string, token: string }>>(
			"/auth/v1/login",
			{
				login: username.value,
				password: password.value
			},
			{
				headers: {
					"Content-Type": "application/json",
					"Accept": "application/json",
				}
			}
		)

		console.log(res.data.data)
		isLoading.value = false

		await router.push({name: "Index"})

	} catch (err) {
		isLoading.value = false

		if (axios.isAxiosError(err) && (err as AxiosError).response) {
			const e = err as AxiosError<Response<{ login: string, token: string }>>

			switch (e.response!.data.code) {
				case code.BAD_REQUEST:
					loginError.value = "Проверьте логин и пароль"
					break
				default:
					loginError.value = e.response!.data.error
					break
			}
		} else {
			console.error(`Ошибка: ${err}`)
			await router.push("/")
		}
	}

};
</script>

<template>
	<Loader v-if="isLoading"/>

	<div class="login">
		<Transition name="fade">
			<h1 class="login__title">TUFF STUFF</h1>
		</Transition>

		<form class="login__form loginForm" @submit.prevent="login">
			<div class="loginForm__group">
				<label class="loginForm__label" for="username">Логин</label>
				<input @input="resetError"
				       v-model="username"
				       class="loginForm__field"
				       id="username"
				       type="text"
				       placeholder="Введите логин">
			</div>

			<div class="loginForm__group">
				<label class="loginForm__label" for="password">Пароль</label>
				<input @input="resetError"
				       v-model="password"
				       class="loginForm__field"
				       id="password"
				       type="password"
				       placeholder="Введите пароль">
			</div>

			<button class="loginForm__button">Вход</button>
		</form>

		<Transition name="fade">
			<div v-if="loginError" class="login__error">
				{{ loginError }}
			</div>
		</Transition>
	</div>
</template>

<style scoped lang="postcss">
.login {
	position: relative;

	min-width: var(--loginFormWidth);
	max-width: var(--loginFormWidth);

	padding: var(--offset);

	background-color: rgb(var(--white));
	border-radius: var(--offset_quarter);

	box-shadow: 0 0 var(--offset_half) 2px rgba(var(--gray), 0.5);
}

.login__error {
	position: absolute;
	bottom: calc(-1 * var(--offset_triply));
	left: 0;

	border-radius: var(--offset_eight);
	background-color: rgb(var(--red));
	color: rgb(var(--white));
	font-weight: bold;

	padding: var(--offset_half);
	max-width: var(--loginFormWidth);
	width: 100%;
}

.login__title {
	font-size: 2.5rem;
	font-weight: 900;
	margin: 0 0 var(--offset) 0;
}

.login__form {
	display: flex;
	flex-direction: column;
	align-items: center;
}

.loginForm__group {
	display: flex;
	flex-direction: column;
	width: 100%;

	margin-bottom: var(--offset);
}

.loginForm__label {
	text-align: left;
	margin-bottom: var(--offset_quarter);
	font-weight: bold;
}

.loginForm__field {
	padding: var(--offset_quarter);
	border: 2px solid rgb(var(--gray));
	border-radius: var(--offset_eight);
}

.loginForm__button {
	width: 100%;
	padding: var(--offset_third);
	border: 2px solid rgb(var(--purple));
	border-radius: var(--offset_eight);

	font-weight: 800;
	cursor: pointer;
	color: rgb(var(--purple));
	background-color: rgb(var(--white));

	transition: var(--mainTransition);
}

.loginForm__button:hover {
	background-color: rgb(var(--purple));
	color: rgb(var(--white));
}
</style>