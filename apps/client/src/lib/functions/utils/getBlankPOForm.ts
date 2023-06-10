export async function getBlankPOForm(fetch: typeof window.fetch) {
	const res = await fetch('/BLANK_PO.png');
	return await res.blob();
}
