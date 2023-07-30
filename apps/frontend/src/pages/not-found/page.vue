<script setup lang="ts">
import { routes } from "@/shared/routes";

import Container from "@/shared/ui/container.vue";

import astronautsImage from "./assets/astronauts.svg";
</script>

<template>
	<div class="wrapper">
		<Container>
			<div class="content">
				<h1 class="title">
					OOOps! <br />
					Page Not Found
				</h1>
				<p class="text">
					Sorry about that! Please visit our homepage to get where you need to
					go.
				</p>
				<router-link class="back-link" :to="routes.HOME">
					Back to homepage
				</router-link>
			</div>
			<img class="image" :src="astronautsImage" alt="404 not found" />
		</Container>
	</div>
</template>

<style scoped lang="scss">
.wrapper {
	padding: 30px 0;
	display: flex;
	align-items: center;
	min-height: 100vh;
	background-color: var(--color-dark-blue);
}

.container {
	display: flex;
	align-items: center;
	justify-content: space-between;

	@media (max-width: 1060px) {
		align-items: center;
		flex-direction: column;
	}
}

.content {
	display: flex;
	flex-direction: column;
	align-items: center;
	margin-right: 30px;

	@media (max-width: 1060px) {
		margin-right: 0;
		margin-bottom: 30px;
	}
}

.title {
	font-size: 46px;
	line-height: normal;
	margin-bottom: 18px;
	color: #fff;
	z-index: 1;
	text-align: center;
}

.text {
	font-size: 21px;
	margin-bottom: 18px;
	max-width: 380px;
	text-align: center;
	color: rgba(255, 255, 255, 0.2);
}

.back-link {
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

.image {
	max-width: 480px;
}
</style>
