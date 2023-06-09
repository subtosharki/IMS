import { API_URL } from '$lib/constants';

export async function deleteCustomer(
	id: string,
	apikey: string,
	fetch: typeof window.fetch
) {
	const res = await fetch(API_URL + '/customers/' + id + '/delete', {
		method: 'DELETE',
		headers: {
			'Content-Type': 'application/json',
			apikey,
		},
		body: JSON.stringify({ id }),
	});
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	} else if (res.status === 404) {
		throw 'Clone not found';
	}
}
