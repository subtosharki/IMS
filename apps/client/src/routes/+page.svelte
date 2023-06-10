<script lang="ts">
	import { login } from '$lib/functions/utils/auth';
	import { goto } from '$app/navigation';
	import {encrypt} from "$lib/functions/utils/crpyt.js";

	export let data;
	let email,
		password,
		errMsg = '';

	function setError(err)  {
		errMsg = err;
		setTimeout(() => {
			errMsg = '';
		}, 5000);
	}

	async function handleLogin() {
		if (!email || !password) {
			setError('Please enter an email and password');
			return;
		}
		try {
			const apikey = await login(email, password, data.fetch);
			localStorage.setItem('apikey', encrypt(apikey));
			await goto('/dashboard');
		} catch (err) {
			setError(err);
		}
	}
</script>

<svelte:head>
	<title>Login</title>
	<meta name="description" content="Login Page" />
</svelte:head>

<main>
	<h1>Login</h1>
	<h2>{errMsg}</h2>
	<form on:submit|preventDefault={async () => await handleLogin()}>
		<label for="email">Email</label>
		<input type="email" id="email" name="email" bind:value={email} />
		<label for="password">Password</label>
		<input type="password" id="password" name="password" bind:value={password} />
		<button type="submit">Login</button>
	</form>
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
</style>
