<h1>Spielmodus</h1>
<div class="game-selection">
    <Select bind:value={selectedGameType} label="Spielmodus">
        {#each gameTypes as gameType}
            <Option value={gameType}>{gameType.label}</Option>
        {/each}
    </Select>
    <Select bind:value={selectedGameEndType} label="Spielende">
        {#each gameEndTypes as gameEndType}
            <Option value={gameEndType}>{gameEndType.label}</Option>
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
            >
            </Slider>
            <span
                slot="label"
                style="padding-right: 12px; width: max-content; display: block;"
            >
                Anzahl Tore zum Sieg: {numberOfGoals}
            </span>
        </FormField>
    {/if}
    <Button class="action-button" href="/game">
        <Icon>
            <Play></Play>
        </Icon>
        <Label>Spiel Starten</Label>
    </Button>
</div>

<script lang="ts">
    import Select, { Option } from '@smui/select';
    import FormField from '@smui/form-field';
    import Slider from '@smui/slider';
    import Button, { Label, Icon } from '@smui/button';
    import Play from 'svelte-material-icons/Play.svelte';

    const gameTypes = [{
        key: '1vs1',
        label: '1 gegen 1'
    }, {
        key: '2vs2',
        label: '2 gegen 2'
    }];
    const gameEndTypes = [{
        key: 'TIME',
        label: 'Nach Ablauf der Zeit'
    }, {
        key: 'RESULT',
        label: 'Nach Anzahl Tore'
    }];

    let selectedGameType = gameTypes[1];
    let selectedGameEndType = gameEndTypes[1];
    let numberOfGoals = 8;
</script>
