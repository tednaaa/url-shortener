import { createApp } from "vue";
import App from "./app.vue";

import "./scss/main.scss";

import { router } from "@/pages/router";

const app = createApp(App);

app.use(router);

app.mount("#app");
