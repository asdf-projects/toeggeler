<script lang="ts">
    import {onMount} from "svelte";
    import {loggedInUser, sessionToken} from "../../shared/dataStore";

    onMount(async () => {
        sessionToken.set('');
        loggedInUser.set('');
    });
</script>