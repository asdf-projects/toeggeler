import { writable } from 'svelte/store';

export const sessionToken = writable('');
export const loggedInUser = writable('');
