<script lang="ts">
	import { goto } from '$app/navigation';
	import { createClone } from '$lib/functions/clones/createClone';
	import type { User } from '$lib/types';

	export let data: App.PageData;
	let name,
		quantity,
		msg = '',
		user =
			data.user ||
			({
				apikey: ''
			} as User),
		date;
	async function handleSubmit() {
		if (!name || !quantity || !date) {
			setError('Please enter a name, quantity and date');
			return;
		}
		date = date.split('-');
		date = date[1] + '/' + date[0];
		await createClone(name, date, quantity, user.apikey, data.fetch!);
		name = '';
		date = '';
		quantity = '';
		msg = 'Clone Added';
	}

	const setError = (err) => {
		msg = err;
		setTimeout(() => {
			msg = '';
		}, 5000);
	};
</script>

<svelte:head>
	<title>Add Clones</title>
	<meta name="description" content="Add Clones Page" />
</svelte:head>
<main>
	<h1>Add Clones</h1>
	<button on:click={async () => await goto('./')}>Back</button>
	<p>Add Clone Page</p>
	<h2>{msg}</h2>

	<form on:submit|preventDefault={async () => await handleSubmit()}>
		<label for="name">Clone Name</label>
		<input type="text" id="name" name="name" bind:value={name} />

		<label for="name">Month Available</label>
		<input type="date" id="month" name="name" bind:value={date} />

		<label for="quantity">Quantity</label>
		<input type="number" id="quantity" name="quantity" bind:value={quantity} />
		<button type="submit">Add Clone</button>
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
