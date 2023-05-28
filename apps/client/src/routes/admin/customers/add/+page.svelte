<script lang="ts">
	import { goto } from '$app/navigation';
	import { createCustomer } from '$lib/functions/customers/createCustomer';

	export let data: App.PageData;
	let email,
		name,
		subscriptionId,
		phone,
		city,
		state,
		zip,
		address,
		msg = '',
		user = data.user || {
			apikey: ''
		};
	async function handleSubmit() {
		if (!email || !name || !subscriptionId || !phone || !address || !city || !zip || !state) {
			setError('Please fill out all fields');
			return;
		}

		await createCustomer(
			{ email, name, subscriptionId, phone, address, city, zip, state },
			user.apikey,
			data.fetch!
		);
		email = '';
		name = '';
		subscriptionId = '';
		phone = '';
		address = '';
		msg = 'Customer Created';
	}

	function setError(err)  {
		msg = err;
		setTimeout(() => {
			msg = '';
		}, 5000);
	}
</script>

<svelte:head>
	<title>Add Customer</title>
	<meta name="description" content="Add Customer Page" />
</svelte:head>
<main>
	<h1>Add Customers</h1>
	<button on:click={async () => await goto('./')}>Back</button>
	<p>Add Customers Page</p>
	<h2>{msg}</h2>

	<form on:submit|preventDefault={async () => await handleSubmit()}>
		<label for="name">Email</label>
		<input type="email" id="email" placeholder="Email" bind:value={email} />
		<label for="name">Name</label>
		<input type="text" id="name" placeholder="Name" bind:value={name} />
		<label for="subId">Subscription ID</label>
		<input type="text" id="subId" placeholder="Subscription ID" bind:value={subscriptionId} />
		<label for="phone">Phone</label>
		<input type="tel" id="phone" placeholder="Phone" bind:value={phone} />
		<label for="address">Address</label>
		<input type="text" id="address" placeholder="Address" bind:value={address} />
		<label for="address">City</label>
		<input type="text" id="city" placeholder="Address" bind:value={city} />
		<label for="address">State</label>
		<input type="text" id="state" placeholder="Address" bind:value={state} />
		<label for="address">Zip</label>
		<input type="text" id="zip" placeholder="Address" bind:value={zip} />
		<button type="submit">Submit</button>
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
