import { API_URL } from '$lib/constants';

export async function deleteUser(id: string, apikey: string, fetch: typeof window.fetch) {
	const res = await fetch(API_URL + '/users/' + id + '/delete', {
		method: 'DELETE',
		headers: {
			apikey
		}
	});
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	} else if (res.status === 404) {
		throw 'User not found';
	}
}
