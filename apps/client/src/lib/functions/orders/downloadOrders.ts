import { getOrders } from '$lib/functions/orders/getOrders';

export async function downloadOrders(
	document: Document,
	apikey: string,
	fetch: typeof window.fetch
) {
	const json = JSON.stringify(await getOrders(apikey, fetch), null, 2);
	const a = document.createElement('a');
	document.body.append(a);
	const date = new Date();
	const dateString = date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate();
	a.download = 'orders-' + dateString + '.json';
	a.href = URL.createObjectURL(new Blob([json], { type: 'application/json' }));
	a.click();
	a.remove();
}
