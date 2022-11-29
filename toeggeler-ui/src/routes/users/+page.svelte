<h1>Das sind unsere registrierten Benutzer</h1>
<div>
    <List
        class="user-list"
        twoLine
        avatarList
        singleSelection
    >
        {#each registeredUsers as user}
            <Item>
                <Graphic style="background-image: url(https://place-hold.it/40x40?text={user.id}&fontsize=16);"></Graphic>
                <Text>
                    <PrimaryText>{user.username}</PrimaryText>
                    <SecondaryText>{user.mail}</SecondaryText>
                </Text>
                <Meta>
                    <Button class="row-button" on:click={() => { console.log('view user detail')}}>
                        <MagnifyPlus></MagnifyPlus>
                    </Button>
                    <Button class="row-button" on:click={() => prepareUserForUpdate(user)}>
                        <AccountEdit></AccountEdit>
                    </Button>
                    <Button class="row-button" on:click={() => deleteUser(user)}>
                        <AccountRemove></AccountRemove>
                    </Button>
                </Meta>
            </Item>
        {/each}
    </List>
    {#if editMode === false }
        <Button class="action-button" on:click={() => { editMode = true; }}>
            <AccountPlus></AccountPlus>
        </Button>
    {:else}
        <div class="user-edit">
            <Textfield bind:value={username} label="Benutzername">
            </Textfield>
            <span class="email">
                <Textfield
                    type="email"
                    updateInvalid
                    bind:value={email}
                    label="E-Mail"
                    input$autocomplete="email"
                >
                    <HelperText validationMsg slot="helper">
                        Ungültige E-Mail Adresse
                    </HelperText>
                </Textfield>
            </span>
            {#if isUpdate === false}
                <Textfield type="password" bind:value={password} label="Passwort">
                </Textfield>
                <Button class="action-button" on:click={() => addUser()}>
                    <ContentSave></ContentSave>
                </Button>
            {:else }
                <Button class="action-button" on:click={() => updateUser()}>
                    <ContentSave></ContentSave>
                </Button>
            {/if}
            <Button class="action-button" on:click={() => resetForm()}>
                <Cancel></Cancel>
            </Button>
        </div>
    {/if}
</div>

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
    import Textfield from '@smui/textfield';
    import AccountPlus from 'svelte-material-icons/AccountPlus.svelte';
    import AccountEdit from 'svelte-material-icons/AccountEdit.svelte';
    import AccountRemove from 'svelte-material-icons/AccountRemove.svelte';
    import MagnifyPlus from 'svelte-material-icons/MagnifyPlus.svelte';
    import ContentSave from 'svelte-material-icons/ContentSave.svelte';
    import Cancel from 'svelte-material-icons/Cancel.svelte';
    import HelperText from '@smui/textfield/helper-text';

    export interface IUser {
        id: number;
        username: string;
        mail: string;
        password: string;
    }

    const loadUsers = async (): Promise<IUser[]> => {
        const response = await fetch('http://localhost:8000/api/users', {
            method: 'GET'
        });
        return await response.json();
    };

    const addUser = async () => {
        const user = { username, mail: email, password };
        const response = await fetch('http://localhost:8000/api/users', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user)
        });
        const newUser = await response.json();
        registeredUsers = [...registeredUsers, newUser];
    };

    const updateUser = async () => {
        const user = { id, username, mail: email };
        const response = await fetch(`http://localhost:8000/api/users/${user.username}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user)
        });
        return await response.json();
    };

    const deleteUser = async (user: IUser): Promise<void> => {
        await fetch(`http://localhost:8000/api/users/${user.username}`, {
            method: 'DELETE'
        });
        registeredUsers = registeredUsers.filter(registeredUser => registeredUser.id !== user.id);
    };

    const prepareUserForUpdate = (user: IUser) => {
        editMode = true;
        isUpdate = true;
        id = user.id;
        username = user.username;
        email = user.mail;
    };

    const resetForm = () => {
        editMode = false;
        isUpdate = false;
    };

    let registeredUsers = [];
    loadUsers().then(users => {
        registeredUsers = users;
    });

    let id: number | null = null;
    let username: string | null = null;
    let email: string | null = null;
    let password: string | null = null;

    let editMode = false;
    let isUpdate = false;
</script>

<style>
    .user-edit {
        display: flex;
        justify-content: flex-start;
        gap: 10px;
    }
    .email {
        flex-direction: column;
    }
    :global(.action-button > svg), :global(.row-button > svg) {
        height: 80%;
        width: 80%;
    }
    :global(.row-button > svg){
        height: 65%;
        width: 65%;
    }
</style>
