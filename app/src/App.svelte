<script lang="ts">
	import "./app.scss";
	import Login from "./lib/login.svelte";
	import Main from "./lib/main.svelte";

	let user = null;
	let token = null;
</script>

<main>
	{#if !user}
		<Login />
	{:else}
		<Main />
	{/if}
</main>
