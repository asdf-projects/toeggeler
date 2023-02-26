<script lang="ts">
	import AlertCircle from 'svelte-material-icons/AlertCircle.svelte';
	import { _ } from 'svelte-i18n';
	export let text;
</script>

{#if text}
	<div class="error-message">
		<AlertCircle />
		<p>{$_(text)}</p>
	</div>
{/if}

<style>
	:global(.error-message) {
		display: flex;
		width: 30%;
		align-items: center;
		color: darkred;
	}
	:global(.error-message p) {
		margin-left: 10px;
	}
</style>
