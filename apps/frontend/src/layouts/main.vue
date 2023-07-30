<script setup lang="ts">
import Container from "@/shared/ui/container.vue";

defineProps<{
	title: string;
}>();
</script>

<template>
	<div class="wrapper">
		<Container>
			<h1 class="title">{{ title }}</h1>

			<slot></slot>
		</Container>
	</div>
</template>

<style scoped lang="scss">
.wrapper {
	padding: 30px 0;
	min-height: 100vh;
	color: #fff;
	background-color: var(--color-dark-blue);
}

.title {
	font-size: 30px;
	line-height: 36px;
	margin-bottom: 30px;

	@media (max-width: 576px) {
		font-size: 26px;
		line-height: 30px;
	}
}
</style>
