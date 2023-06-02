import { API_URL } from '$lib/constants';

export async function setCloneQuantity(
	id: string,
	quantity: number,
	apikey: string,
	fetch: typeof window.fetch
) {
	const res = await fetch(API_URL + `/clones/${id}/quantity/${quantity}`, {
		method: 'PUT',
		headers: {
			apikey,
		},
	});
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	}
}
