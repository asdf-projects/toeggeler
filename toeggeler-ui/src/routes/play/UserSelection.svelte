<span>
    {#await loadUsers()}
    {:then users}
        <Select
            key={(user) => `${user ? user.id : ''}`}
            bind:value={selectedUser}
            label="{ $_(placeholder) }"
        >
            {#each users as user}
                <Option value={user}>{user.username} ({user.mail})</Option>
            {/each}
        </Select>
    {/await}
</span>

<script lang="ts">

    import { _ } from 'svelte-i18n';
    import Select, {Option} from "@smui/select";

    export interface IUser {
        id: number;
        username: string;
        mail: string;
    }


    export let selectedUser: IUser;
    export let placeholder: string;

    const loadUsers = async (): Promise<IUser[]> => {
        const response = await fetch('http://localhost:8000/api/users', {
            method: 'GET'
        });
        return await response.json();
    };
</script>