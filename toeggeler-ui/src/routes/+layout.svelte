<nav>
    <TabBar {tabs} let:tab bind:active>
        <Tab {tab} href="{tab.target}">
            <Icon>
                {#if tab.icon === 'SoccerField'}
                    <SoccerField></SoccerField>
                {/if}
                {#if tab.icon === 'AccountCircle'}
                    <AccountCircle></AccountCircle>
                {/if}
                {#if tab.icon === 'Trophy'}
                    <Trophy></Trophy>
                {/if}
                {#if tab.icon === 'AccountGroup'}
                    <AccountGroup></AccountGroup>
                {/if}
            </Icon>
            <Label>{tab.label}</Label>
        </Tab>
    </TabBar>
</nav>
<slot></slot>

<script lang="ts">
    import Tab, { Icon, Label } from '@smui/tab';
    import TabBar from '@smui/tab-bar';
    import SoccerField from 'svelte-material-icons/SoccerField.svelte';
    import Trophy from 'svelte-material-icons/Trophy.svelte';
    import AccountGroup from 'svelte-material-icons/AccountGroup.svelte';
    import AccountCircle from 'svelte-material-icons/AccountCircle.svelte';

    let tabs = [
        {
            icon: 'SoccerField',
            label: 'Spielen',
            target: '/play'
        },
        {
            icon: 'Trophy',
            label: 'Rangliste',
            target: '/scoreboard'
        },
        {
            icon: 'AccountGroup',
            label: 'Benutzer',
            target: '/users'
        },
        {
            icon: 'AccountCircle',
            label: 'Benutzer-Administration',
            target: '/administration'
        }
    ];
    let active;
</script>