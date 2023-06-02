import { API_URL } from '$lib/constants';
import type { Clone } from '$lib/types';

export async function getClones(apikey: string, fetch: typeof window.fetch) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/clones', {
			headers: {
				apikey,
			},
		});
	} catch (e) {
		throw 'Error connecting to server';
	}
	if (!res.ok) throw 'Error getting clones';
	return (await res.json()) as Clone[];
}
