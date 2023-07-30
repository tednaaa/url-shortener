<template>
	<div class="container">
		<slot></slot>
	</div>
</template>

<style scoped lang="scss">
.container {
	max-width: calc(var(--container-width) + 30px);
	width: 100%;
	padding: 0 15px;
	margin: 0 auto;
}
</style>
