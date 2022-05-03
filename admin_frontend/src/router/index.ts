import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";

import Index from "../pages/Index.vue";
import Login from "../pages/Login.vue";
import Lang from "../pages/Lang.vue";
import Options from "../pages/Options.vue";

const routes: RouteRecordRaw[] = [
	{
		path: "/",
		name: "Index",
		component: Index,
		meta: {
			title: "Index",
		},
	},
	{
		path: "/lang/:lang",
		name: "Lang",
		component: Lang,
		meta: {
			title: "Lang",
		},
	},
	{
		path: "/options",
		name: "Options",
		component: Options,
		meta: {
			title: "Options",
		},
	},
	{
		path: "/login",
		name: "Login",
		component: Login,
		meta: {
			title: "Login",
		},
	}
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;