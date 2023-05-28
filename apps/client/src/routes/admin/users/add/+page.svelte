<script lang="ts">
	import { createUser } from '$lib/functions/users/createUser';
	import { goto } from '$app/navigation';

	export let data: App.PageData;
	let email,
		password,
		admin,
		errMsg = '',
		user = data.user || {
			apikey: ''
		};
	async function handleSubmit() {
		if (!email || !password) {
			setError('Please enter an email and password');
			return;
		}
		await createUser({ email, password, admin }, user.apikey, data.fetch!);
		email = '';
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
