<script lang="ts">
	import { get } from 'svelte/store';
	import { loggedInUser, sessionToken } from '../shared/dataStore';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { _ } from 'svelte-i18n';

	let userName = '';

	onMount(async () => {
		const storedJwtToken = get(sessionToken);
		if (storedJwtToken === '') {
			await goto('/login');
		}
		userName = get(loggedInUser);
	});
</script>

<h1>{$_('Home.Welcome')} {userName}</h1>
