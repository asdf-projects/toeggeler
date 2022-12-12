<div class="game-selection">
    <Select bind:value={selectedGameType} label="{ $_('Play.GameType.Selection') }">
        {#each gameTypes as gameType}
            <Option value={gameType}>{ $_(gameType.label) }</Option>
        {/each}
    </Select>
    <Select bind:value={selectedGameEndType} label="{ $_('Play.GameEndType.Selection') }">
        {#each gameEndTypes as gameEndType}
            <Option value={ gameEndType }>{ $_(gameEndType.label) }</Option>
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
                { $_('Play.NumberOfGoals') } { numberOfGoals }
            </span>
        </FormField>
    {/if}
    <Button class="action-button" href="/game">
        <Icon>
            <Play></Play>
        </Icon>
        <Label>{ $_('Play.StartGame') }</Label>
    </Button>
</div>

<script lang="ts">
    import { _ } from 'svelte-i18n';
    import Select, { Option } from '@smui/select';
    import FormField from '@smui/form-field';
    import Slider from '@smui/slider';
    import Button, { Label, Icon } from '@smui/button';
    import Play from 'svelte-material-icons/Play.svelte';

    const gameTypes = [{
        key: '1vs1',
        label: 'Play.GameType.1vs1'
    }, {
        key: '2vs2',
        label: 'Play.GameType.2vs2'
    }];
    const gameEndTypes = [{
        key: 'TIME',
        label: 'Play.GameEndType.Time'
    }, {
        key: 'RESULT',
        label: 'Play.GameEndType.Result'
    }];

    let selectedGameType = gameTypes[1];
    let selectedGameEndType = gameEndTypes[1];
    let numberOfGoals = 8;
</script>
