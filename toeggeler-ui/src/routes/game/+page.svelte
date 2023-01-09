<h1>GO GO GO</h1>
{#if !gameEnded}
    <p>{ $_('Game.Score')} {scoreTeam1}:{scoreTeam2}</p>
    <LayoutGrid>
        {#await playerData then players}
            {#each players as player, index}
                <Cell span={6}>
                    {#if (mousedownTimer?.duration >= 500 && mousedownTimer?.buttonIndex === index)}
                        <div class="overlay">
                            <Button on:click={() => scoreGoal(player.id, EventType.GOAL)}>
                                <Icon>
                                    <Soccer></Soccer>
                                </Icon>
                            </Button>
                            <Button on:click={() => scoreGoal(player.id, EventType.FOETELI)}>
                                <Icon>
                                    <CameraWireless></CameraWireless>
                                </Icon>
                            </Button>
                            <Button on:click={() => scoreGoal(player.id, EventType.OWN_GOAL)}>
                                <Icon>
                                    <SkipBackward></SkipBackward>
                                </Icon>
                            </Button>
                        </div>
                    {:else}
                        <Button
                                class="player"
                                on:click={() => scoreGoal(player.id)}
                                on:mousedown={() => mouseDownTimerStart(index) }
                                on:mouseup={() => mousedownTimerStop() }
                        >{ player.username }</Button>
                    {/if}
                </Cell>
                {#if index===1}
                    <Cell span={12}>
                        <SoccerField width="70%" height="50px"></SoccerField>
                    </Cell>
                {/if}
            {/each}
        {/await}
    </LayoutGrid>
{:else}
    <p>{ $_('Game.EndOfGame') } { scoreTeam1 }:{ scoreTeam2 }</p>
{/if}

<script lang="ts">
    import { _ } from 'svelte-i18n';
    import LayoutGrid, { Cell } from '@smui/layout-grid';
    import Button, { Icon } from '@smui/button';
    import CameraWireless from 'svelte-material-icons/CameraWireless.svelte';
    import SkipBackward from 'svelte-material-icons/SkipBackward.svelte';
    import Soccer from 'svelte-material-icons/Soccer.svelte';
    import SoccerField from 'svelte-material-icons/SoccerField.svelte';
    import {page} from "$app/stores";
    import type {IUser} from "../../app";

    export interface ITeam {
        offense: number;
        defense: number;
    }
    enum EventType {
        GAME_START = 'GAME_START',
        GAME_END = 'GAME_END',
        GOAL = 'GOAL',
        OWN_GOAL = 'OWN_GOAL',
        FOETELI = 'FOETELI'
    }
    export interface IGameEvent {
        event: EventType;
        timestamp: number;
        team1?: ITeam;
        team2?: ITeam;
        player?: number;
    }

    let scoreTeam1 = 0;
    let scoreTeam2 = 0;
    const currentEvents = [];
    let gameEnded = false;
    let mousedownTimer: { start: number; duration: number; buttonIndex: number; };
    const team1: ITeam = JSON.parse($page.url.searchParams.get('team1'));
    const team2: ITeam = JSON.parse($page.url.searchParams.get('team2'));

    const getPlayerData = async (id: number): Promise<IUser> => {
        const response = await fetch(`http://localhost:8000/api/users`, {
            method: 'GET'
        });
        const users: IUser[] = await response.json();
        return users.filter(user => user.id === id)[0];
    };

    const playerData = Promise.all([
        getPlayerData(team1.offense),
        getPlayerData(team1.defense),
        getPlayerData(team2.defense),
        getPlayerData(team2.offense)
    ]);

    const mouseDownTimerStart = (buttonIndex: number) => {
        mousedownTimer = { start: Date.now(), duration: 0, buttonIndex: buttonIndex };
    };
    const mousedownTimerStop = () => {
        mousedownTimer.duration = Date.now() - mousedownTimer.start;
    }
    const updateScore = (player: number, eventType: EventType) => {
        if (eventType === EventType.OWN_GOAL) {
            if (player === team1.offense || player === team1.defense) {
                scoreTeam2++;
            } else {
                scoreTeam1++;
            }
        } else {
            if (player === team1.offense || player === team1.defense) {
                scoreTeam1++;
            } else {
                scoreTeam2++;
            }
        }
    };
    const scoreGoal = (player: number, eventType: EventType = EventType.GOAL) => {
        const event: IGameEvent = { event: eventType, timestamp: Date.now(),  player }
        storeEvent(event);
        updateScore(player, eventType);
        if (scoreTeam1 === 8 || scoreTeam2 === 8) {
            storeEvent({
                event: EventType.GAME_END,
                timestamp: Date.now()
            });
            gameEnded = true;
            shareGameResult(currentEvents);
        }
        mouseDownTimerStart(undefined);
    };
    const storeEvent = (event: IGameEvent) => {
        console.log(JSON.stringify(event));
        currentEvents.push(event);
    };

    const shareGameResult = async (events: IGameEvent[]) => {
        const response = await fetch('http://localhost:8000/api/games', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(events)
        });
        await response.json();
    };

    storeEvent({
        event: EventType.GAME_START,
        timestamp: Date.now(),
        team1,
        team2
    });
</script>
