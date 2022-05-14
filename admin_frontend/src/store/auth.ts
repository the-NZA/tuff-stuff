import {defineStore} from "pinia";

export const useAuthStore = defineStore("auth", {
	state: () => ({
		isLoggedIn: false,
	}),
	getters: {
		IsLoggedIn: (state) => state.isLoggedIn,
	},
	actions: {
		SetLogin(status: boolean) {
			this.isLoggedIn = status;
		},
	}
})