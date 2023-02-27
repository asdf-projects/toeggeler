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
	import ErrorMessage from '../../shared/ErrorMessage.svelte';
	import { getErrorMessage } from '../../shared/utils';

	let username: string;
	let userData: IUser;

	let dirty: boolean;
	let invalid: boolean;

	let errorMessage: string;

	onMount(async () => {
		username = get(loggedInUser);
		if (username === '') {
			await goto('/login', { replaceState: false, state: { name: '/administration' } });
		}
		userData = await getUserData(username);
	});
	const getUserData = async (username: string): Promise<IUser> => {
		const response = await fetch('http://localhost:8000/api/users', {
			method: 'GET'
		});
		const users: IUser[] = await response.json();
		return users.filter((user) => user.username === username)[0];
	};

	const updateUser = async (userData: IUser): Promise<IUser> => {
		errorMessage = '';
		return fetch(`http://localhost:8000/api/users/${userData.id}`, {
			method: 'PUT',
			body: JSON.stringify(userData)
		})
			.then(async (response) => {
				const data = await response.json();

				if (!response.ok) {
					const error = getErrorMessage(data);
					return Promise.reject(error);
				}
				dirty = false;
				return data;
			})
			.catch((error) => {
				errorMessage = error;
			});
	};
</script>

<div>
	{#if userData}
		<Textfield bind:value={username} label={$_('Signup.Username')} disabled />
		<div class="email">
			<Textfield
				type="email"
				bind:dirty
				bind:invalid
				updateInvalid
				bind:value={userData.mail}
				label={$_('Signup.Email')}
				input$autocomplete="email"
			>
				<HelperText validationMsg slot="helper">
					{$_('Signup.InvalidEmail')}
				</HelperText>
			</Textfield>
		</div>
		<Button on:click={updateUser(userData)} disabled={!dirty || invalid}>
			<Icon>
				<ContentSave />
			</Icon>
			<Label>
				{$_('Administration.Save')}
			</Label>
		</Button>
		<ErrorMessage bind:text={errorMessage} />
	{/if}
</div>
