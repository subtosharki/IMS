import { ENCRYPTION_KEY } from '$lib/constants';
import { AES, enc } from 'crypto-js';

export function encrypt(data: string) {
	return AES.encrypt(data, ENCRYPTION_KEY).toString();
}

export function decrypt(data: string) {
	return AES.decrypt(data, ENCRYPTION_KEY).toString(enc.Utf8);
}
