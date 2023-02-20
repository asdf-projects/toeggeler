<script lang="ts">
	import '$lib/i18n';
	import { _, isLoading } from 'svelte-i18n';
	import Tab, { Icon, Label } from '@smui/tab';
	import TabBar from '@smui/tab-bar';
	import SoccerField from 'svelte-material-icons/SoccerField.svelte';
	import Trophy from 'svelte-material-icons/Trophy.svelte';
	import AccountGroup from 'svelte-material-icons/AccountGroup.svelte';
	import AccountCircle from 'svelte-material-icons/AccountCircle.svelte';
	import LogoutVariant from 'svelte-material-icons/LogoutVariant.svelte';

	let tabs = [
		{
			icon: 'SoccerField',
			label: 'Menu.Play',
			target: '/play'
		},
		{
			icon: 'Trophy',
			label: 'Menu.Stats',
			target: '/stats'
		},
		{
			icon: 'AccountGroup',
			label: 'Menu.Users',
			target: '/users'
		},
		{
			icon: 'AccountCircle',
			label: 'Menu.Administration',
			target: '/administration'
		},
        {
            icon: 'LogoutVariant',
            label: 'Menu.Logout',
            target: '/logout'
        }
	];
	let active;
</script>

{#if $isLoading}
	<p>Wait for Translations...</p>
{:else}
	<nav>
		<TabBar {tabs} let:tab bind:active>
			<Tab {tab} href={tab.target}>
				<Icon>
					{#if tab.icon === 'SoccerField'}
						<SoccerField />
					{/if}
					{#if tab.icon === 'AccountCircle'}
						<AccountCircle />
					{/if}
					{#if tab.icon === 'Trophy'}
						<Trophy />
					{/if}
					{#if tab.icon === 'AccountGroup'}
						<AccountGroup />
					{/if}
                    {#if tab.icon === 'LogoutVariant'}
                        <LogoutVariant />
                    {/if}
				</Icon>
				<Label>{$_(tab.label)}</Label>
			</Tab>
		</TabBar>
	</nav>
	<slot />
{/if}
