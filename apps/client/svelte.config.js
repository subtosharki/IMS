import adapter from '@sveltejs/adapter-node';
import { vitePreprocess } from '@sveltejs/kit/vite';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter(),
		files: {
			assets: 'static',
			serviceWorker: 'src/service-worker.js',
		},
		serviceWorker: {
			register: false,
		}
	},
};

export default config;
