<script lang="ts">
	import { get } from 'svelte/store';
	import { loggedInUser } from '../../shared/dataStore';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import type { IUser } from '../../app';
	import Textfield from '@smui/textfield';
	import { _ } from 'svelte-i18n';
	import HelperText from '@smui/textfield/helper-text';
	import Button, { Icon, Label } from '@smui/button';
	import ContentSave from 'svelte-material-icons/ContentSave.svelte';

	let username: string;
	let userData: IUser;

	onMount(async () => {
		username = get(loggedInUser);
		if (username === '') {
			await goto('/login');
		}
	});
	const getUserData = async (username: string): Promise<IUser> => {
		const response = await fetch('http://localhost:8000/api/users', {
			method: 'GET'
		});
		const users: IUser[] = await response.json();
		return users.filter((user) => user.username === username)[0];
	};
</script>

<div>
	{#await getUserData(username) then userData}
		<Textfield bind:value={username} label={$_('Signup.Username')} disabled />
		<span class="email">
			<Textfield
				type="email"
				updateInvalid
				value={userData.mail}
				label={$_('Signup.Email')}
				input$autocomplete="email"
			>
				<HelperText validationMsg slot="helper">
					{$_('Signup.InvalidEmail')}
				</HelperText>
			</Textfield>
		</span>
		<Button>
			<Icon>
				<ContentSave />
			</Icon>
			<Label>
				{$_('Administration.Save')}
			</Label>
		</Button>
	{/await}
</div>
