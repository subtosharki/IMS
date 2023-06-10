import { getBlankPOForm } from '$lib/functions/utils/getBlankPOForm';

export async function fillPOForm(fetch: typeof window.fetch) {
	const poBlob = await getBlankPOForm(fetch);
}
