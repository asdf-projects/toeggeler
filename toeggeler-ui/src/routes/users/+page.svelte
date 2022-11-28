<h1>Das sind unsere registrierten Benutzer</h1>
<div>
    <List
            class="user-list"
            twoLine
            avatarList
            singleSelection
            bind:selectedIndex={selectionIndex}
    >
        {#each registeredUsers as user}
            <Item>
                <Graphic style="background-image: url(https://place-hold.it/40x40?text={user.id}&fontsize=16);"></Graphic>
                <Text>
                    <PrimaryText>{user.username}</PrimaryText>
                    <SecondaryText>{user.mail}</SecondaryText>
                </Text>
                <Meta>
                    <Button on:click={() => { console.log('view user detail')}}>
                        <MagnifyPlus></MagnifyPlus>
                    </Button>
                    <Button on:click={() => prepareUserForUpdate(user)}>
                        <AccountEdit></AccountEdit>
                    </Button>
                    <Button on:click={() => deleteUser(user)}>
                        <AccountRemove></AccountRemove>
                    </Button>
                </Meta>
            </Item>
        {/each}
    </List>
    {#if editMode === false }
        <Button on:click={() => { editMode = true; }}>
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
                <Button on:click={addUser}>
                    <ContentSave></ContentSave>
                </Button>
            {:else }
                <Button on:click={updateUser()}>
                    <ContentSave></ContentSave>
                </Button>
            {/if}
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
    import HelperText from '@smui/textfield/helper-text';

    export interface IUser {
        id: number;
        username: string;
        mail: string;
        password: string;
    }

    const loadUsers = async (): Promise<IUser[]> => {
        const response = await fetch('http://localhost:8080/api/users', {
            method: 'GET'
        });
        return await response.json();
    };

    const addUser = async (): Promise<IUser> => {
        const user = { username, mail: email, password };
        const response = await fetch('http://localhost:8080/api/users', {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user)
        });
        return await response.json();
    };

    const prepareUserForUpdate = (user: IUser): void => {
        editMode = true;
        isUpdate = true;
        id = user.id;
        username = user.username;
        email = user.mail;
    };

    const updateUser = async (): Promise<IUser> => {
        const user = { id, username, mail: email };
        const response = await fetch(`http://localhost:8080/api/users/${user.username}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user)
        });
        return await response.json();
    };

    const deleteUser = async (user: IUser): Promise<void> => {
        const response = await fetch(`http://localhost:8080/api/users/${user.username}`, {
            method: 'DELETE'
        });
        return await response.json();
    };

    let registeredUsers = [];
    loadUsers().then(users => {
        registeredUsers = users;
    });

    let selectionIndex: number | undefined = undefined;
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
</style>
