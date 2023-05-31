import { goto } from '$app/navigation';
import type { PageLoad } from './$types';
import { browser } from '$app/environment';
import { getUserByAPIKey } from '$lib/functions/users/getUserByAPIKey';
import { ENCRYPTION_KEY } from '$lib/constants';
import { AES, enc } from 'crypto-js';

export const load = (async ({ fetch }) => {
	if (browser && localStorage.getItem('apikey')) {
		try {
			const user = await getUserByAPIKey(
					AES.decrypt(localStorage.getItem('apikey') as string, ENCRYPTION_KEY).toString(enc.Utf8),
					fetch
				)
			if (!user) {
				localStorage.removeItem('apikey');
			} else {
				await goto('/dashboard');
			}
		} catch (e) {
			localStorage.removeItem('apikey');
		}
	}
	return { fetch };
}) satisfies PageLoad;
