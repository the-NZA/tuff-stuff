<template>
	<div class="settings">
		<h2 class="settings__title">Настройки сайта</h2>

		<div class="settings__row">
			<label for="emailSetting" class="settings__label">Название</label>

			<p class="settings__snippet">
				Lorem ipsum dolor sit amet, consectetur adipisicing elit. In, voluptatibus?
			</p>

			<input type="text" id="emailSetting" class="settings__input" placeholder="Название">
		</div>

		<div class="settings__row">
			<label for="phoneSetting" class="settings__label">Название</label>

			<p class="settings__snippet">
				Lorem ipsum dolor sit amet, consectetur adipisicing elit. In, voluptatibus?
			</p>

			<input type="text" id="phoneSetting" class="settings__input" placeholder="Название">
		</div>

		<div class="settings__row">
			<label for="telegramSetting" class="settings__label">Название</label>

			<p class="settings__snippet">
				Lorem ipsum dolor sit amet, consectetur adipisicing elit. In, voluptatibus?
			</p>

			<input type="text" id="telegramSetting" class="settings__input" placeholder="Название">
		</div>

		<div class="settings__footer">
			<button class="settings__save">Сохранить</button>

			<div v-if="false" class="settings__error">
				Lorem ipsum dolor sit amet, consectetur adipisicing elit. Cum dolores error esse harum ipsam
				necessitatibus nisi provident quibusdam? Cumque debitis magnam modi molestiae mollitia quod, reiciendis
				reprehenderit rerum sint temporibus.
			</div>

			<div v-if="true" class="settings__success">
				Lorem ipsum dolor sit amet, consectetur adipisicing elit. Cum dolores error esse harum ipsam
				necessitatibus nisi provident quibusdam? Cumque debitis magnam modi molestiae mollitia quod, reiciendis
				reprehenderit rerum sint temporibus.
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
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
}

.settings__success {
	padding: var(--offset_half);

	background-color: rgb(var(--light-orange));
	color: rgb(var(--white));

	font-weight: bold;
	border-radius: var(--offset_eight);
	text-shadow: 0 0 1px rgba(var(--gray), var(--alpha50));
}


</style>
