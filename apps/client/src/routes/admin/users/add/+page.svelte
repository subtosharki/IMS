<script lang="ts">
	import { createUser } from '$lib/functions/users/createUser';
	import { goto } from '$app/navigation';
	import type {User} from "$lib/types";
	import {getUserByAPIKey} from "$lib/functions/users/getUserByAPIKey";
	import {decrypt} from "$lib/functions/crpyt";
	import {onMount} from "svelte";

	export let data;
	let email,
			firstName, lastName,
		password,
		admin,
		errMsg = '',
		user =  {
			apikey: ''
		} as User;
	onMount(async () => {
		user = await getUserByAPIKey(decrypt(data.apikey), data.fetch)
	});

	async function handleSubmit() {
		if (!email || !password) {
			setError('Please enter an email and password');
			return;
		}
		await createUser({ email, password, firstName, lastName, admin }, user.apikey, data.fetch!);
		email = '';
		firstName = '';
		lastName = '';
		password = '';
		admin = false;
		errMsg = 'User Created';
	}

	const setError = (err) => {
		errMsg = err;
		setTimeout(() => {
			errMsg = '';
		}, 5000);
	};
</script>

<svelte:head>
	<title>Add User</title>
	<meta name="description" content="Add User Page" />
</svelte:head>
<main>
	<h1>Add Users</h1>
	<button on:click={async () => await goto('./')}>Back</button>
	<p>Add Users Page</p>
	<h2>{errMsg}</h2>

	<form on:submit|preventDefault={async () => await handleSubmit()}>
		<label for="email">Email</label>
		<input type="email" id="email" name="email" bind:value={email} />

		<label for="firstName">First Name</label>
		<input type="text" id="firstName" name="firstName" bind:value={firstName} />
		<label for="lastName">Last Name</label>
		<input type="text" id="lastName" name="lastName" bind:value={lastName} />
		<label for="password">Password</label>
		<input type="password" id="password" name="password" bind:value={password} />
		<label for="admin">Admin</label>
		<input type="checkbox" id="admin" name="admin" bind:checked={admin} />
		<button type="submit">Add User</button>
	</form>
</main>

<style>
	main {
		text-align: center;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	input {
		margin: 0.5rem;
	}

	button {
		background-color: #4caf50;
		border: none;
		color: white;
		padding: 15px 32px;
		margin: 4px 2px;
		cursor: pointer;
	}
</style>
