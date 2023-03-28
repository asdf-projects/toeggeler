<script lang="ts">
	import { get } from 'svelte/store';
    import {loggedInUserId, sessionToken} from '../../shared/dataStore';
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

	let userId: number;
	let userData: IUser;

	let dirty: boolean;
	let invalid: boolean;

	let errorMessage: string;

	onMount(async () => {
		userId = get(loggedInUserId);
		if (userId === -1) {
			await goto('/login');
		}
		userData = await getUserData(userId);
	});

	const getUserData = async (userId: number): Promise<IUser> => {
		const response = await fetch(`/api/users/${userId}`, {
			method: 'GET'
		});
		return await response.json();
	};

	const updateUser = async (userData: IUser): Promise<IUser> => {
		errorMessage = '';
		return fetch(`/api/users/${userData.id}`, {
			method: 'PUT',
            headers: {
                Authorization: `Bearer ${get(sessionToken)}`
            },
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
		<Textfield bind:value={userData.username} label={$_('Signup.Username')} disabled />
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
