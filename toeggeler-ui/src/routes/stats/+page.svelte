<script lang="ts">
	import { _ } from 'svelte-i18n';
	import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
	import type { IStatistic } from '../../app';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { loggedInUser } from '../../shared/dataStore';

	let username;

	onMount(async () => {
		username = get(loggedInUser);
	});

	const getUserStats = async (): Promise<IStatistic[] & { winLossRatio: number }> => {
		const response = await fetch('/api/stats', {
			method: 'GET'
		});
		const userStatistics = await response.json();
		return userStatistics.map((userStatistic) => ({
			...userStatistic,
			winLossRatio: getWinLossRatio(userStatistic.wins, userStatistic.losses)
		}));
	};

	const getWinLossRatio = (numberOfWins: number, numberOfLosses: number): number => {
		if (numberOfLosses === 0) {
			return 100;
		}
		return Math.round((numberOfWins / (numberOfWins + numberOfLosses)) * 100);
	};

	const getUsername = async (userId): Promise<string> => {
		const response = await fetch(`/api/users/${userId}`, {
			method: 'GET'
		});
		const userDetail = await response.json();
		return userDetail.username;
	};
</script>

<div>
	<h2>{$_('Stats.Stats')}</h2>
	{#await getUserStats() then statistics}
		<DataTable sortable>
			<Head>
				<Row>
					<Cell>{$_('Stats.Username')}</Cell>
					<Cell>{$_('Stats.Games')}</Cell>
					<Cell>{$_('Stats.WinLossRatio')}</Cell>
					<Cell>{$_('Stats.Goals')}</Cell>
					<Cell>{$_('Stats.Foeteli')}</Cell>
					<Cell>{$_('Stats.OwnGoals')}</Cell>
					<Cell>{$_('Stats.Rating')}</Cell>
				</Row>
			</Head>
			<Body>
				{#each statistics as statistic}
					{#await getUsername(statistic.playerId) then usernameOfPlayer}
						<Row selected={usernameOfPlayer == username}>
							<Cell>{usernameOfPlayer}</Cell>
							<Cell>{statistic.wins + statistic.losses}</Cell>
							<Cell>{statistic.winLossRatio}%</Cell>
							<Cell>{statistic.goals}</Cell>
							<Cell>{statistic.foetelis}</Cell>
							<Cell>{statistic.ownGoals}</Cell>
							<Cell>{statistic.rating}</Cell>
						</Row>
					{/await}
				{/each}
			</Body>
		</DataTable>
	{/await}
</div>

<style>
	:global([selected='true']) {
		background-color: palegreen;
	}
</style>
