import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";
import {useAuthStore} from "../store/auth";
import {useMessageStore} from "../store/message";
import {withBaseSlug} from "../util/withBaseSlug";

import HTTP from "../util/http";
import {Response} from "../types/Response";
import axios, {AxiosError} from "axios";

import Index from "../pages/Index.vue";
import Login from "../pages/Login.vue";
import Lang from "../pages/Lang.vue";
import Options from "../pages/Options.vue";

const routes: RouteRecordRaw[] = [
	{
		path: withBaseSlug("/"),
		name: "Index",
		component: Index,
		meta: {
			title: "Index",
			public: false,
		},
	},
	{
		path: withBaseSlug("/lang/:lang"),
		name: "Lang",
		component: Lang,
		meta: {
			title: "Lang",
			public: false,
		},
	},
	{
		path: withBaseSlug("/options"),
		name: "Options",
		component: Options,
		meta: {
			title: "Options",
			public: false,
		},
	},
	{
		path: withBaseSlug("/login"),
		name: "Login",
		component: Login,
		meta: {
			title: "Login",
			public: true,
		},
	}
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

router.beforeEach(async (to, from, next) => {
	const authStore = useAuthStore()
	const messageStore = useMessageStore()

	// Reset all messages on route change
	messageStore.Reset();

	// If public route, just continue
	if (to.matched.some(record => record.meta.public)) {
		next()
		return
	}

	// Try to ping server
	try {
		await HTTP.get<Response<string>>("/api/v1")

		authStore.SetLogin(true)

		// Check if to route is Login
		if (to.name === "Login") {
			next({name: "Index"})
			return
		}

		next()
		return

	} catch (err) {
		// Check axios error
		if (axios.isAxiosError(err) && (err as AxiosError).response) {
			authStore.SetLogin(false)
			next({name: "Login"})
		} else {
			next("/")
		}
	}
})

export default router;