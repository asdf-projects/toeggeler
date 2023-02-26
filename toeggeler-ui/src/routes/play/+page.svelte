<script lang="ts">
	import { _ } from 'svelte-i18n';
	import Select, { Option } from '@smui/select';
	import FormField from '@smui/form-field';
	import Slider from '@smui/slider';
	import Button, { Label, Icon } from '@smui/button';
	import Play from 'svelte-material-icons/Play.svelte';
	import UserSelection from './UserSelection.svelte';
	import { goto } from '$app/navigation';
	import type { ITeam } from '../../app';
	import { page } from '$app/stores';

	const gameTypes = [
		{
			key: '1vs1',
			label: 'Play.GameType.1vs1'
		},
		{
			key: '2vs2',
			label: 'Play.GameType.2vs2'
		}
	];
	const gameEndTypes = [
		{
			key: 'TIME',
			label: 'Play.GameEndType.Time'
		},
		{
			key: 'RESULT',
			label: 'Play.GameEndType.Result'
		}
	];

	let selectedGameType = gameTypes[1];
	let selectedGameEndType = gameEndTypes[1];
	let numberOfGoals = 8;
	let team1: ITeam = { offense: undefined as number, defense: undefined as number };
	let team2: ITeam = { offense: undefined as number, defense: undefined as number };

	$: isValidGame =
		team1.offense &&
		team1.defense &&
		team2.offense &&
		team2.defense &&
		[...new Set([team1.offense, team1.defense, team2.offense, team2.defense])].length === 4;

	const startGame = async () => {
		$page.url.searchParams.set('team1', JSON.stringify(team1));
		$page.url.searchParams.set('team2', JSON.stringify(team2));

		await goto(`/game?${$page.url.searchParams.toString()}`);
	};
</script>

<div class="game-selection">
	<h2>{$_('Play.GameSettings')}</h2>
	<Select bind:value={selectedGameType} label={$_('Play.GameType.Selection')} disabled>
		{#each gameTypes as gameType}
			<Option value={gameType}>{$_(gameType.label)}</Option>
		{/each}
	</Select>
	<Select bind:value={selectedGameEndType} label={$_('Play.GameEndType.Selection')} disabled>
		{#each gameEndTypes as gameEndType}
			<Option value={gameEndType}>{$_(gameEndType.label)}</Option>
		{/each}
	</Select>
	{#if selectedGameEndType?.key === 'RESULT'}
		<FormField align="end" style="display: flex;">
			<Slider
				bind:value={numberOfGoals}
				min={0}
				max={10}
				step={1}
				discrete
				input$aria-label="Slider to select the number of Goals to win"
				style="flex-grow: 1;"
				disabled
			/>
			<span slot="label" style="padding-right: 12px; width: max-content; display: block;">
				{$_('Play.NumberOfGoals')}
				{numberOfGoals}
			</span>
		</FormField>
	{/if}

	<h2>{$_('Play.UserSelection.Selection')}</h2>
	<h3>{$_('Play.UserSelection.Team1')}</h3>
	<UserSelection bind:selectedUser={team1.offense} placeholder="Play.UserSelection.Offense" />
	<UserSelection bind:selectedUser={team1.defense} placeholder="Play.UserSelection.Defense" />
	<h3>{$_('Play.UserSelection.Team2')}</h3>
	<UserSelection bind:selectedUser={team2.offense} placeholder="Play.UserSelection.Offense" />
	<UserSelection bind:selectedUser={team2.defense} placeholder="Play.UserSelection.Defense" />

	<div>
        <Button class="action-button" on:click={startGame} disabled={!isValidGame}>
            <Icon>
                <Play />
            </Icon>
            <Label>{$_('Play.StartGame')}</Label>
        </Button>
        {#if !isValidGame}
			<p>{$_('Play.UserSelection.ErrorMessage')}</p>
		{/if}
	</div>
</div>
