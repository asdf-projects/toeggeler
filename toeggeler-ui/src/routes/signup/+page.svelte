<script lang="ts">
	import Textfield from '@smui/textfield';
	import { _ } from 'svelte-i18n';
	import HelperText from '@smui/textfield/helper-text';
	import Button, { Icon, Label } from '@smui/button';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import ContentSave from 'svelte-material-icons/ContentSave.svelte';
	import ErrorMessage from '../../shared/ErrorMessage.svelte';

	let username = '';
	let email = '';
	let password = '';

	let invalid: boolean;

	let errorMessage;

	const onClickSave = async () => {
		const newUser: { id } = await addUser();
		$page.url.searchParams.set('highlightedUserId', newUser.id);
		await goto(`/users?${$page.url.searchParams.toString()}`);
	};
	const addUser = async () => {
		const user = { username, mail: email, password };
		return fetch('http://localhost:8000/api/users', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(user)
		})
			.then(async (response) => {
				const isJson = response.headers.get('content-type')?.includes('application/json');
				const data = isJson ? await response.json() : null;

				if (!response.ok) {
					const error =
						(data && data.message) || response.status === 400
							? $_('Error.InvalidInput')
							: $_('Error.General');
					return Promise.reject(error);
				}
				return data;
			})
			.catch((error) => {
				errorMessage = error;
			});
	};
</script>

<form>
	<Textfield bind:value={username} label={$_('Signup.Username')} />
	<div class="email">
		<Textfield
			type="email"
			bind:invalid
			updateInvalid
			bind:value={email}
			label={$_('Signup.Email')}
			input$autocomplete="email"
		>
			<HelperText validationMsg slot="helper">
				{$_('Signup.InvalidEmail')}
			</HelperText>
		</Textfield>
	</div>
	<Textfield type="password" bind:value={password} label={$_('Signup.Password')} />
	<div>
		<Button
			class="action-button"
			on:click={onClickSave}
			disabled={username.length === 0 || invalid || password.length === 0}
		>
			<Icon>
				<ContentSave />
			</Icon>
			<Label>{$_('Signup.Save')}</Label>
		</Button>
	</div>
	<ErrorMessage bind:text={errorMessage} />
</form>
