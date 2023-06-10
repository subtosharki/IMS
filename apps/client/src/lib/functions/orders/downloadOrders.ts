import { getOrders } from '$lib/functions/orders/getOrders';
import type { Clone } from '$lib/types';
import {getUserByAPIKey} from "$lib/functions/users/getUserByAPIKey";

export async function downloadOrders(
	apikey: string,
	fetch: typeof window.fetch,
) {
	const user = await getUserByAPIKey(apikey, fetch);
	let orders = await getOrders(apikey, fetch);
	if(!orders) {
		return alert('No orders found')
	}
	if(user.role !== 'admin') {
		orders = orders.filter(order => order.placedBy === user.email)
	}
	if(orders.length === 0) {
		return alert('No orders found')
	}
	const keys = Object.keys(orders[0]);
	let csv = '';
	for (const key of keys) {
		csv += key + ',';
	}
	for (const order of orders) {
		csv += '\n';
		for (const key of keys) {
			if (key === 'clones') {
				csv +=
					(order as any)[key]
						.map((clone: Clone) => clone.name)
						.join(';') + ',';
				continue;
			}
			csv += (order as any)[key] + ',';
		}
	}
	const a = document.createElement('a');
	document.body.append(a);
	const date = new Date();
	const dateString =
		date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate();
	a.download = 'orders-' + dateString + '.csv';
	a.href = URL.createObjectURL(new Blob([csv], { type: 'text/csv' }));
	a.click();
	a.remove();
}
