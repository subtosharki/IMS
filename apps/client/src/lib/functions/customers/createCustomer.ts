import type { CreateCustomerBody } from '$lib/types';
import { API_URL } from '$lib/constants';

export async function createCustomer(
	body: CreateCustomerBody,
	apikey: string,
	fetch: typeof window.fetch
) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/customers/create', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				apikey
			},
			body: JSON.stringify(body)
		});
	} catch (e) {
		throw 'Error connecting to server';
	}
	if (res.status === 400) {
		throw 'Bad request';
	} else if (res.status === 500) {
		throw 'Internal server error';
	} else if (res.status === 404) {
		throw 'Clone not found';
	}
}
