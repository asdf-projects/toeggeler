<script lang="ts">
	import { _ } from 'svelte-i18n';
	import Select, { Option } from '@smui/select';
	import type { IUser } from '../../app';

	export let selectedUser: number;
	export let placeholder: string;

	const loadUsers = async (): Promise<IUser[]> => {
		const response = await fetch('/api/users', {
			method: 'GET'
		});
		return await response.json();
	};
</script>

<span>
	{#await loadUsers() then users}
		<Select
			key={(user) => `${user ? user.id : ''}`}
			bind:value={selectedUser}
			label={$_(placeholder)}
		>
			{#each users as user}
				<Option value={user.id}>{user.username} ({user.mail})</Option>
			{/each}
		</Select>
	{/await}
</span>
