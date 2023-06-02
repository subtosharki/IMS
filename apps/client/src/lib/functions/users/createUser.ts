import { API_URL } from '$lib/constants';
import type { CreateUserBody } from '$lib/types';

export async function createUser(
	body: CreateUserBody,
	apikey: string,
	fetch: typeof window.fetch
) {
	const res = await fetch(API_URL + '/users/create', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			apikey,
		},
		body: JSON.stringify(body),
	});
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	}
}
