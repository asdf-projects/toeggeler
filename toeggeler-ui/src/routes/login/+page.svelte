<script lang="ts">
	import Textfield from '@smui/textfield';
	import { _ } from 'svelte-i18n';
	import Button, { Icon, Label } from '@smui/button';
	import AccountPlus from 'svelte-material-icons/AccountPlus.svelte';
	import Login from 'svelte-material-icons/Login.svelte';
	import { loggedInUser, sessionToken } from '../../shared/dataStore';
	import { goto } from '$app/navigation';

	let username = '';
	let password = '';
	let errorMessage = '';

	export interface ILoginToken {
		token: string;
	}

	const login = async () => {
		const loginData = { username, password };
		const response = await fetch('http://localhost:8000/api/authenticate', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(loginData)
		});
		if (response.ok) {
			const loginResponse: ILoginToken = await response.json();
			if (storeSessionData(loginResponse.token)) {
				await goto('/');
			}
		} else if (response.status === 401) {
			errorMessage = $_('Login.LoginFailed');
			return;
		}
		errorMessage = $_('Login.GeneralError');
	};

	const storeSessionData = (loginToken: string): boolean => {
		if (!verifyToken(loginToken)) {
			return false;
		}
		sessionToken.update(() => loginToken);
		const jwtPayloadEncoded = loginToken.split('.')[1];
		const jwtPayload = JSON.parse(atob(jwtPayloadEncoded));
		loggedInUser.update(() => jwtPayload.username);
		return true;
	};

	const verifyToken = (jwtToken: string): boolean => {
		const tokenParts = jwtToken.split('.');
		return tokenParts.length === 3;
	};
</script>

<form>
	<Textfield bind:value={username} label={$_('Login.Username')} />
	<Textfield type="password" bind:value={password} label={$_('Login.Password')} />
	<Button class="action-button" on:click={login}>
		<Icon>
			<Login />
		</Icon>
		<Label>{$_('Login.Login')}</Label>
	</Button>
	{#if errorMessage.length > 0}
		<p>{errorMessage}</p>
	{/if}
</form>
<Button class="action-button" on:click={async () => await goto('/signup')}>
	<Icon>
		<AccountPlus />
	</Icon>
	<Label>{$_('Login.Signup')}</Label>
</Button>
