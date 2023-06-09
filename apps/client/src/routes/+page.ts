import type { PageLoad } from '../../.svelte-kit/types/src/routes/$types';

export const load = (async ({ fetch }) => {
	return { fetch };
}) satisfies PageLoad;
