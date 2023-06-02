import { goto } from '$app/navigation';
import type { PageLoad } from './$types';
import { getUserByAPIKey } from '$lib/functions/users/getUserByAPIKey';
import { browser } from '$app/environment';
import { AES, enc } from 'crypto-js';
import { ENCRYPTION_KEY } from '$lib/constants';

export const load = (async ({ fetch }) => {
	if (browser) {
		if (!localStorage.getItem('apikey')) await goto('/');
		try {
			const user = await getUserByAPIKey(
				AES.decrypt(
					localStorage.getItem('apikey') as string,
					ENCRYPTION_KEY
				).toString(enc.Utf8),
				fetch
			);
			return { user, fetch };
		} catch (e) {
			localStorage.removeItem('apikey');
			await goto('/');
		}
	}
}) satisfies PageLoad;
