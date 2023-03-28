<script lang="ts">
	import Textfield from '@smui/textfield';
	import { _ } from 'svelte-i18n';
	import Button, { Icon, Label } from '@smui/button';
	import AccountPlus from 'svelte-material-icons/AccountPlus.svelte';
	import Login from 'svelte-material-icons/Login.svelte';
    import {loggedInUser, loggedInUserId, sessionToken} from '../../shared/dataStore';
    import {afterNavigate, goto} from '$app/navigation';
	import ErrorMessage from '../../shared/ErrorMessage.svelte';
	import { getErrorMessage } from '../../shared/utils';

	let username = '';
	let password = '';
	let errorMessage = '';

    let previousPage = '/';

	export interface ILoginToken {
		token: string;
	}

    afterNavigate(({from}) => {
        if (from?.url?.pathname !== '/login') {
            previousPage = from?.url?.pathname || '/';
        }
    })

	const login = async () => {
		errorMessage = '';
		const loginData = { username, password };
		const response = await fetch('/api/authenticate', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(loginData)
		});
		if (response.ok) {
			const loginResponse: ILoginToken = await response.json();
			if (storeSessionData(loginResponse.token)) {
				await goto(previousPage);
			}
		} else {
            const errorResponse = await response.json()
			errorMessage = getErrorMessage(errorResponse);
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
        loggedInUserId.set(jwtPayload.id);
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
	<Button
		class="action-button"
		on:click={login}
		disabled={username.length === 0 || password.length === 0}
	>
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
