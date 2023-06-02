import { API_URL } from '$lib/constants';
import type { Order } from '$lib/types';

export async function placeOrder(
	body: {
		dateRequired: string;
		clones: { quantity: number; name: string; id: string; date: string }[];
		use: string;
		customerName: string;
		notes: string;
	},
	apikey: string,
	fetch: typeof window.fetch
) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/orders/place', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				apikey,
			},
			body: JSON.stringify(body),
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
	return (await res.json())['order'] as Order;
}
