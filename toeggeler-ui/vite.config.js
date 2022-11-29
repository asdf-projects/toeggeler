import { sveltekit } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
    server: {
        port: 8000,
        proxy: {
            '/api': {
                target: 'http://localhost:8080'
            }
        },
    },
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
};

export default config;
