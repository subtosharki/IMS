<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { placeOrder } from '$lib/functions/orders/placeOrder';
	import { getMonthName } from '$lib/functions/getMonthName';
	import type { Clone, Order, OrderClone, User } from '$lib/types';
	import { getClones } from '$lib/functions/clones/getClones';
	import { getCustomerNames } from '$lib/functions/customers/getCustomerNames';

	export let data;
	let clones: Clone[] = [],
		user = data.user || ({ apikey: '' } as User),
		customerName = '',
		use = '',
		dateRequired = '',
		placedOrder = '' as Order,
		notes = '',
			customers = [] as string[];

	onMount(async () => {
		clones = await getClones(user.apikey, data.fetch!);
		customers = await getCustomerNames(user.apikey, data.fetch!);
	});

	async function handleOrder() {
		if (use === 'Choose an option') {
			alert('Please select a use');
			return;
		}
		const selectedClones = clones.filter(
			(clone: OrderClone) => clone.selectedQuantity > 0
		) as OrderClone[];
		if (selectedClones.length === 0) {
			alert('Please select at least one clone');
			return;
		}

		const order = {
			customerName,
			use,
			dateRequired,
			clones: selectedClones.map((clone) => ({
				id: clone.id,
				cloneId: clone.cloneId,
				quantity: clone.selectedQuantity,
				name: clone.name,
				date: clone.date
			})),
			notes
		};
		placedOrder = await placeOrder(order, user.apikey, data.fetch!);

		customerName = '';
		use = '';
		notes = '';
		dateRequired = '';
		clones = clones.map((clone: Clone) => ({
			...clone,
			selectedQuantity: 0
		}));
	}
</script>

<svelte:head>
	<title>Place an order</title>
	<meta name="description" content="Order form" />
</svelte:head>

<main>
	<button on:click={async () => await goto('./')}>Back</button>
	{#if placedOrder}
		<h2>Order Placed</h2>
		<p>Order ID: {placedOrder.orderNumber}</p>
		<p>Customer Name: {placedOrder.customerName}</p>
		<p>Use: {placedOrder.use}</p>
		<p>Date Required: {placedOrder.dateRequired}</p>
		<p>Clones:</p>
		<ul>
			{#each placedOrder.clones as clone}
				<li>{clone.name} - {clone.quantity}</li>
			{/each}
		</ul>
	{/if}
	<h1>Place an order</h1>
	<form on:submit|preventDefault={async () => await handleOrder()}>
		<label for="name">Customer Purchasing
			<br>
			<input list="name" bind:value={customerName} />
		</label>
		<datalist id="name">
			{#each customers as customer}
				<option label={customer}>{customer}</option>
			{/each}
		</datalist>


		<label for="use">Med or AU</label>
		<select id="use" bind:value={use}>
			<option label="Choose an option">Choose an option</option>
			<option label="Medical">Medical</option>
			<option label="Adult Use">Adult Use</option>
		</select>

		<label for="readytime">Date Required</label>
		<input type="date" id="readytime" name="readytime" required bind:value={dateRequired} />

		<label for="clones">Select Clones</label>
		<table id="clones">
			<thead>
				<tr>
					<th>Clone Name</th>
					<th>Month</th>
					<th>Year</th>
					<th>Quantity</th>
					<th>Select</th>
				</tr>
			</thead>
			<tbody>
				{#each clones as clone}
					<tr>
						<td>{clone.name}</td>
						<td>{getMonthName(Number(clone.date.split('/')[0]))}</td>
						<td>{clone.date.split('/')[1]}</td>
						<td>{clone.quantity}</td>
						<td>
							<input
								type="number"
								min="0"
								max={clone.quantity}
								id="clone-{clone.id}"
								placeholder="0"
								bind:value={clone.selectedQuantity}
								on:change={async() => {
									if(clone.selectedQuantity === 0) {
										clones = await getClones(user.apikey, data.fetch);
									} else {
									clones = clones.filter((c) => {
										if (getMonthName(Number(c.date.split('/')[0])) === getMonthName(Number(clone.date.split('/')[0])) && c.date.split('/')[1] === clone.date.split('/')[1]) {
											return clone;
										}
									});
									}
								}}
							/>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
		<br>

		<label for="notes">Extra Notes</label>
		<textarea id="notes" bind:value={notes} cols="62" rows="5"/>

		<button type="submit" on:click={() => (window.location.href = '#top')}>Place Order</button>
	</form>
</main>

<style>
	main {
		text-align: center;
	}

	form {
		display: inline-block;
		text-align: left;
	}

	label {
		display: block;
	}

	input {
		margin-bottom: 1rem;
	}

	button {
		margin-top: 1rem;
		display: block;
	}

	table {
		margin-top: 1rem;
		border-collapse: collapse;
	}

	th,
	td {
		padding: 0.5rem;
		border: 1px solid black;
	}

	tr:nth-child(even) {
		background-color: #dddddd;
	}

	select {
		margin-bottom: 1rem;
	}

	input[type='number'] {
		width: 4rem;
	}

	input[type='date'] {
		width: 10rem;
	}

	input[type='text'] {
		width: 20rem;
	}
</style>
