import { API_URL } from '$lib/constants';

export async function updateCustomerNotes(
	customerId: string,
	notes: string,
	apikey: string,
	fetch: typeof window.fetch
) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/customers/' + customerId + '/notes', {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json',
				apikey,
			},
			body: JSON.stringify({
				notes,
			}),
		});
	} catch (e) {
		throw 'Error connecting to server';
	}
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	} else if (res.status === 404) {
		throw 'Order not found';
	}
}
