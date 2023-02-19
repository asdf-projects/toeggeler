<script lang="ts">
	import Textfield from '@smui/textfield';
	import { _ } from 'svelte-i18n';
	import HelperText from '@smui/textfield/helper-text';
	import Button, { Label } from '@smui/button';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let username: string | null = null;
	let email: string | null = null;
	let password: string | null = null;

	const onClickSave = async () => {
		const newUser: { id } = await addUser();
		$page.url.searchParams.set('highlightedUserId', newUser.id);
		goto(`/users?${$page.url.searchParams.toString()}`);
	};
	const addUser = async () => {
		const user = { username, mail: email, password };
		const response = await fetch('http://localhost:8000/api/users', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(user)
		});
		return await response.json();
	};
</script>

<form>
	<Textfield bind:value={username} label={$_('Signup.Username')} />
	<span class="email">
		<Textfield
			type="email"
			updateInvalid
			bind:value={email}
			label={$_('Signup.Email')}
			input$autocomplete="email"
		>
			<HelperText validationMsg slot="helper">
				{$_('Signup.InvalidEmail')}
			</HelperText>
		</Textfield>
	</span>
	<Textfield type="password" bind:value={password} label={$_('Signup.Password')} />
	<Button class="action-button" on:click={onClickSave}>
		<Label>{$_('Signup.Save')}</Label>
	</Button>
</form>
