import { init, register } from 'svelte-i18n';

const defaultLocale = 'de';

register('de', () => import('./locales/de.json'));

init({
	fallbackLocale: defaultLocale,
	initialLocale: defaultLocale
});
