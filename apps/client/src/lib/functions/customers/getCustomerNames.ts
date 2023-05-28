import { API_URL } from '$lib/constants';

export async function getCustomerNames(apikey: string, fetch: typeof window.fetch) {
	const res = await fetch(API_URL + '/customers/names', {
		method: 'GET',
		headers: {
			apikey
		}
	});
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	}
	return (await res.json()) as string[];
}
