<script lang="ts">
	import Textfield from '@smui/textfield';
	import { _ } from 'svelte-i18n';
	import Button, { Icon, Label } from '@smui/button';
	import AccountPlus from 'svelte-material-icons/AccountPlus.svelte';
	import Login from 'svelte-material-icons/Login.svelte';
	import { loggedInUser, sessionToken } from '../../shared/dataStore';
	import { goto } from '$app/navigation';
    import {page} from "$app/stores";
    import ErrorMessage from "../../shared/ErrorMessage.svelte";
    import {getErrorMessage} from "../../shared/utils";

	let username = '';
	let password = '';
	let errorMessage = '';

	export interface ILoginToken {
		token: string;
	}

	const login = async () => {
        errorMessage = '';
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
		} else {
            errorMessage = getErrorMessage(response);
        }
	};

	const storeSessionData = (loginToken: string): boolean => {
		if (!verifyToken(loginToken)) {
			return false;
		}
		sessionToken.set(loginToken);
		const jwtPayloadEncoded = loginToken.split('.')[1];
		const jwtPayload = JSON.parse(atob(jwtPayloadEncoded));
		loggedInUser.set(jwtPayload.username);
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
	<Button class="action-button" on:click={login} disabled={username.length===0||password.length===0}>
		<Icon>
			<Login />
		</Icon>
		<Label>{$_('Login.Login')}</Label>
	</Button>
	<ErrorMessage bind:text={errorMessage} />
</form>
<Button class="action-button" on:click={async () => await goto('/signup')}>
	<Icon>
		<AccountPlus />
	</Icon>
	<Label>{$_('Login.Signup')}</Label>
</Button>
