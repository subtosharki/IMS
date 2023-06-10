import { API_URL } from '$lib/constants';

export async function login(
	email: string,
	password: string,
	fetch: typeof window.fetch
) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/auth/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ email, password }),
		});
	} catch (e) {
		throw 'Error connecting to server';
	}
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 401) {
		throw 'Invalid credentials';
	} else if (res.status === 500) {
		throw 'Internal server error';
	}
	return (await res.json())['apikey'] as string;
}
