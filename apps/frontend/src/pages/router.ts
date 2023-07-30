import { createRouter, createWebHistory } from "vue-router";

import { routes } from "@/shared/routes";

import { HomePage } from "./home";
import { NotFoundPage } from "./not-found";

export const router = createRouter({
	history: createWebHistory(),
	routes: [
		{
			path: routes.HOME,
			component: HomePage,
		},
		{
			path: "/:pathMatch(.*)*",
			component: NotFoundPage,
		},
	],
});
