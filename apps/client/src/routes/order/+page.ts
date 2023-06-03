import { goto } from '$app/navigation';
import type { PageLoad } from './$types';
import { browser } from '$app/environment';

export const load = (async ({ fetch }) => {
	let apikey;
	if (browser) {
		apikey = localStorage.getItem('apikey');
		if (!apikey) return goto('/');
	}
	return { apikey, fetch };
}) satisfies PageLoad;
