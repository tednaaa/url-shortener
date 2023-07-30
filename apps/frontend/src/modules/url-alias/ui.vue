<script setup lang="ts">
import { onMounted, ref, watch } from "vue";

import { getUrls, createUrl } from "./api/api";
import { Url } from "./api/types";

const isError = ref<boolean>(false);
const urls = ref<Url[]>([]);
const aliasInput = ref<string>("");
const urlInput = ref<string>("");

watch([() => aliasInput.value, () => urlInput.value], () => {
	isError.value = false;
});

onMounted(async () => {
	const { data } = await getUrls();

	if (data.urls) urls.value = data.urls;
});

const createAlias = async () => {
	if (!aliasInput.value || !urlInput.value) return;

	const { data } = await createUrl(urlInput.value, aliasInput.value);

	if (data.status === "OK") {
		urls.value.push({
			alias: aliasInput.value,
			url: urlInput.value,
		});
		return;
	}

	isError.value = true;
};
</script>

<template>
	<div class="inputs">
		<input
			class="input"
			v-model="aliasInput"
			placeholder="Alias..."
			type="text"
		/>
		<input class="input" v-model="urlInput" placeholder="Url..." type="text" />
	</div>
	<button
		@click="createAlias"
		class="button"
		:class="{ 'button--error': isError }"
	>
		Create
	</button>
	<ul class="list" v-if="urls.length">
		<li class="list-item" v-for="{ alias, url } in urls">
			<span class="list-item__alias">{{ alias }}</span>
			<a class="list-item__url" target="_blank" :href="url">Redirect</a>
		</li>
	</ul>
</template>

<style scoped lang="scss">
.inputs {
	display: flex;
	gap: 10px;
	margin-bottom: 15px;
}

.input {
	padding: 10px;
	border: 2px solid gold;
	border-radius: 8px;
	color: white;

	&:focus-visible {
		background-color: #334574;
		border: 2px solid #0e0e0e;
	}
}

.button {
	padding: 10px 30px;
	background-color: #334574;
	border-radius: 8px;
	color: white;
	transition: 0.3s;

	margin-left: auto;
	margin-bottom: 25px;

	&:hover {
		background-color: #395cb6;
	}
}

.button--error {
	background: indianred;

	&:hover {
		background: indianred;
	}
}

.list-item {
	display: flex;
	gap: 10px;
	justify-content: center;
	align-items: center;

	&:not(:last-child) {
		margin-bottom: 15px;
	}
}

.list-item__alias {
	width: 50%;
	background: goldenrod;
	color: #0e0e0e;
	padding: 10px;
	border-radius: 8px;
}

.list-item__url {
	width: 50%;
	background: deepskyblue;
	color: #0e0e0e;
	padding: 10px 20px;

	border-top-right-radius: 8px;
	border-bottom-right-radius: 8px;
	border-radius: 8px;

	transition: 0.3s;

	&:hover {
		background: gold;
	}
}
</style>
