import { browser } from '$app/environment';
import type { PageLoad } from './$types';
import { goto } from '$app/navigation';

export const load = (async () => {
	if (browser) {
		if (localStorage.getItem('apikey')) {
			localStorage.removeItem('apikey');
			await goto('/');
		} else await goto('/');
	}
}) satisfies PageLoad;
