import { API_URL } from '$lib/constants';
import type { User } from '$lib/types';

export async function getUsers(apikey: string, fetch: typeof window.fetch) {
	const res = await fetch(API_URL + '/users', {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			apikey,
		},
	});
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	}
	return (await res.json()) as User[];
}
