// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
export interface IUser {
    id: number;
    username: string;
    mail: string;
    password?: string;
}

export interface ITeam {
    offense: number;
    defense: number;
}

export interface IStatistic {
    foetelis: number;
    goals: number;
    losses: number;
    ownGoals: number;
    playerId: number;
    rating: number;
    wins: number;
}

declare namespace App {
	// interface Locals {}
	// interface PageData {}
	// interface Error {}
	// interface Platform {}
}
