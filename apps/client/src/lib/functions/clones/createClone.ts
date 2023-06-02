import { API_URL } from '$lib/constants';

export async function createClone(
	name: string,
	date: string,
	quantity: number,
	apikey: string,
	fetch: typeof window.fetch
) {
	const res = await fetch(API_URL + '/clones/create', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			apikey,
		},
		body: JSON.stringify({
			name,
			quantity,
			date,
		}),
	});
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	}
}
