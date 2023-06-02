import { API_URL } from '$lib/constants';
import type { Order } from '$lib/types';

export async function getOrders(apikey: string, fetch: typeof window.fetch) {
	const res = await fetch(API_URL + '/orders', {
		method: 'GET',
		headers: {
			apikey,
		},
	});
	if (!res.ok) throw 'Error getting orders';
	return (await res.json()) as Order[];
}
