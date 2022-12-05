<h1>GO GO GO</h1>
{#if !gameEnded}
    <p>Aktueller Spielstand: {scoreTeam1}:{scoreTeam2}</p>
    <LayoutGrid>
        <Cell span={6}>
            {#await team1Ofense then player}
                <Button class="player" on:click={() => scoreGoal(player.id)}>{ player.username }</Button>
            {/await}
        </Cell>
        <Cell span={6}>
            {#await team1Defense then player}
                <Button class="player" on:click={() => scoreGoal(player.id)}>{ player.username }</Button>
            {/await}
        </Cell>
        <Cell span="{12}">

        </Cell>
        <Cell span={6}>
            {#await team2Defense then player}
                <Button class="player" on:click={() => scoreGoal(player.id)}>{ player.username }</Button>
            {/await}
        </Cell>
        <Cell span={6}>
            {#await team2Ofense then player}
                <Button class="player" on:click={() => scoreGoal(player.id)}>{ player.username }</Button>
            {/await}
        </Cell>
    </LayoutGrid>
{:else}
    <p>Spiel beendet. Schlussstand: {scoreTeam1}:{scoreTeam2}</p>
{/if}

<script lang="ts">
    import LayoutGrid, { Cell } from '@smui/layout-grid';
    import Button from '@smui/button';

    export interface ITeam {
        ofense: number;
        defense: number;
    }
    export interface IUser {
        id: number,
        username: string;
        mail: string;
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
    const team1: ITeam = { ofense: 1, defense: 2 };
    const team2: ITeam = { ofense: 3, defense: 4 };
    const getPlayerData = async (id: number): Promise<IUser> => {
        const response = await fetch(`http://localhost:8000/api/users`, {
            method: 'GET'
        });
        const users: IUser[] = await response.json();
        return users.filter(user => user.id === id)[0];
    };
    const team1Ofense = getPlayerData(team1.ofense);
    const team1Defense = getPlayerData(team1.defense);
    const team2Ofense = getPlayerData(team2.ofense);
    const team2Defense = getPlayerData(team2.defense);

    const scoreGoal = (player: number) => {
        const event: IGameEvent = { event: EventType.GOAL, timestamp: Date.now(),  player }
        storeEvent(event);
        if (team1.ofense === player || team1.defense === player) {
            scoreTeam1++;
        }
        if (team2.ofense === player || team2.defense === player) {
            scoreTeam2++;
        }
        if (scoreTeam1 === 8 || scoreTeam2 === 8) {
            storeEvent({
                event: EventType.GAME_END,
                timestamp: Date.now()
            });
            gameEnded = true;
            shareGameResult(currentEvents);
        }
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

<style>
</style>