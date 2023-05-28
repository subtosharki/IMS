import { API_URL } from '$lib/constants';
import type { OrderStatus } from '$lib/types';

export async function setOrderStatus(
	orderId: string,
	status: OrderStatus,
	apikey: string,
	fetch: typeof window.fetch
) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/orders/' + orderId + '/status', {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json',
				apikey
			},
			body: JSON.stringify({
				status
			})
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
