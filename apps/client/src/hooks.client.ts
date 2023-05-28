import { handleErrorWithSentry, Replay } from '@sentry/sveltekit';
import { init } from '@sentry/sveltekit';

init({
	dsn: 'https://2bd0dfa2458a4e75a9451029b0768541@o4505252803510272.ingest.sentry.io/4505252805935104',
	tracesSampleRate: 1.0,
	replaysSessionSampleRate: 0.1,
	replaysOnErrorSampleRate: 1.0,
	integrations: [new Replay()]
});

export const handleError = handleErrorWithSentry();
