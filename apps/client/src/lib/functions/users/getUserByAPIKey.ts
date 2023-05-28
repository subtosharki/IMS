import { API_URL } from '$lib/constants';
import type { User } from '$lib/types';

export async function getUserByAPIKey(apikey: string, fetch: typeof window.fetch) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/users/' + apikey);
	} catch (e) {
		throw 'Error connecting to server';
	}
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	} else if (res.status === 404) {
		throw 'User not found';
	}
	return (await res.json()) as User;
}
