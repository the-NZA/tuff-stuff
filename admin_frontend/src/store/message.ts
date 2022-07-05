import {defineStore} from "pinia";

export const useMessageStore = defineStore("message", {
	state: () => ({
		errorMessage: "",
		isError: false,
		modalErrorMessage: "",
		isModalError: false,
		modalSuccessMessage: "",
		isModalSuccess: false,
		successMessage: "",
		isSuccess: false,
		isShowModal: false,
	}),
	getters: {
		ErrorMessage: (state) => state.errorMessage,
		IsError: (state) => state.isError,
		SuccessMessage: (state) => state.successMessage,
		IsSuccess: (state) => state.isSuccess,
		ModalErrorMessage: (state) => state.modalErrorMessage,
		IsModalError: (state) => state.isModalError,
		ModalSuccessMessage: (state) => state.modalSuccessMessage,
		IsModalSuccess: (state) => state.isModalSuccess,
		IsShowModal: (state) => state.isShowModal,
	},
	actions: {
		SetError(message: string) {
			this.errorMessage = message;
			this.isError = true;
		},
		SetSuccess(message: string) {
			this.successMessage = message;
			this.isSuccess = true;
		},
		SetModalError(message: string) {
			this.modalErrorMessage = message;
			this.isModalError = true;
		},
		SetModalSuccess(message: string) {
			this.modalSuccessMessage = message;
			this.isModalSuccess = true;
		},
		SetShowModal(isShowModal: boolean) {
			this.isShowModal = isShowModal;
		},
		Reset() {
			this.errorMessage = "";
			this.isError = false;
			this.successMessage = "";
			this.isSuccess = false;
			this.modalErrorMessage = "";
			this.isModalError = false;
			this.modalSuccessMessage = "";
			this.isModalSuccess = false;
		}
	}
})