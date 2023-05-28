import { API_URL } from '$lib/constants';

export async function editCloneName(
	cloneId: string,
	name: string,
	apikey: string,
	fetch: typeof window.fetch
) {
	let res: Response;
	try {
		res = await fetch(API_URL + '/clones/' + cloneId + '/name', {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json',
				apikey
			},
			body: JSON.stringify({ name })
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
