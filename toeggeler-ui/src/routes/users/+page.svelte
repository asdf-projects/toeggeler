<div class="users-with-stats">
    <div class="user-list">
        <List
            class="user-list"
            twoLine
            avatarList
            singleSelection
        >
        {#await loadUsers()}
        {:then registeredUsers}
            {#each registeredUsers as user}
                <Item>
                    <Graphic style="background-image: url(https://place-hold.it/40x40?text={user.id}&fontsize=16);"></Graphic>
                    <Text>
                        <PrimaryText>{user.username}</PrimaryText>
                        <SecondaryText>{user.mail}</SecondaryText>
                    </Text>
                    <Meta>
                        <Button class="row-button" on:click={ () => { selectedUser = user; }}>
                            <MagnifyPlus></MagnifyPlus>
                        </Button>
                    </Meta>
                </Item>
            {/each}
        {/await}
        </List>
    </div>
    <div class="user-stats">
        {#if selectedUser}
            {#await loadUserStatistics(selectedUser.id)}
            {:then stats}
                <h3>{$_('Users.StatsFor')} {selectedUser.username}</h3>
                <List>
                    <Item>
                        <Text>
                            {$_('Users.Wins')}: {stats.wins} / {stats.losses} ({getWinLossRatio(stats.wins, stats.losses)}%)
                        </Text>
                    </Item>
                    <Item>
                        <Text>
                            {$_('Users.Goals')}: {stats.goals}
                        </Text>
                    </Item>
                    <Item>
                        <Text>
                            {$_('Users.Foetelis')}: {stats.foetelis}
                        </Text>
                    </Item>
                    <Item>
                        <Text>
                            {$_('Users.OwnGoals')}: {stats.ownGoals}
                        </Text>
                    </Item>
                    <Item>
                        <Text>{$_('Users.Rating')}: {stats.rating}</Text>
                    </Item>
                </List>
            {/await}
        {/if}
    </div>
</div>
<Button class="action-button" on:click={ () => goto('/signup') }>
    <AccountPlus></AccountPlus>
</Button>

<script lang="ts">
    import List, {
        Item,
        Graphic,
        Meta,
        Text,
        PrimaryText,
        SecondaryText,
    } from '@smui/list';
    import Button from '@smui/button';
    import AccountPlus from 'svelte-material-icons/AccountPlus.svelte';
    import MagnifyPlus from 'svelte-material-icons/MagnifyPlus.svelte';
    import type {IStatistic, IUser} from "../../app";
    import {goto} from "$app/navigation";
    import {_} from "svelte-i18n";

    let selectedUser;

    const loadUsers = async (): Promise<IUser[]> => {
        const response = await fetch('http://localhost:8000/api/users', {
            method: 'GET'
        });
        return await response.json();
    };

    const getWinLossRatio = (numberOfWins: number, numberOfLosses: number): number => {
        if (numberOfLosses === 0) {
            return 100;
        }
        return Math.round((numberOfWins/(numberOfWins + numberOfLosses))*100);
    }

    const loadUserStatistics = async (userId: string): Promise<IStatistic> => {
        const response = await fetch(`http://localhost:8000/api/stats/${userId}`, {
            method: 'GET'
        });
        return await response.json();
    };
</script>

<style>
    :global(.action-button > svg), :global(.row-button > svg) {
        height: 80%;
        width: 80%;
    }
    :global(.row-button > svg){
        height: 65%;
        width: 65%;
    }
    :global(div.users-with-stats) {
        display:flex;
        gap: 50px;
    }
</style>
