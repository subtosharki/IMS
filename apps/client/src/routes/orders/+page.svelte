<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { getOrders } from '$lib/functions/orders/getOrders';
	import { setOrderStatus } from '$lib/functions/orders/setOrderStatus';
	import { updateOrderNotes } from '$lib/functions/orders/updateOrderNotes';
	import { downloadOrders } from '$lib/functions/orders/downloadOrders';
	import type { Order, User } from '$lib/types';
	import {getUserByAPIKey} from "$lib/functions/users/getUserByAPIKey";
	import {decrypt} from "$lib/functions/utils/crpyt";
	import { getMonthName } from '$lib/functions/utils/getMonthName';

	export let data
	let orders: Order[] = [],
			user = {
				email: '',
				role: ''
			} as User

	async function loadOrders() {
		orders = await getOrders(user.apikey, data.fetch!);
		if(!orders) return;
	}
	onMount(async () => {
		user = await getUserByAPIKey(decrypt(data.apikey), data.fetch)
		await loadOrders();
	});
</script>

<svelte:head>
	<title>Orders</title>
	<meta name="description" content="View Orders" />
</svelte:head>

<main>
	<button on:click={async () => await goto('/dashboard')}>Back</button>
	<h1>Orders</h1>
	<p>Welcome, {user.email} to the orders page</p>
		<button on:click={async () => await downloadOrders(user.apikey, data.fetch)}
			>Download Order Logs</button
		>
	<button on:click={async () => await goto('orders/barcode')}>Barcode Edit Mode</button>
	<select on:change={async () => {
		const status = (document.querySelector('select')).value
		if(status === 'all') {
			await loadOrders()
		} else {
			await loadOrders()
			orders = orders.filter((order) => order.status.includes(status))
		}
	}}>
		<option value="all" selected >All</option>
		<option value="1">1 - Subculture</option>
		<option value="2">2 - Multiplication</option>
		<option value="3">3 - Transplant</option>
		<option value="4">4 - Rooting</option>
		<option value="5">5 - Finished Product</option>
		<option value="6">6 - Fulfilled</option>
			<option value="Voided">Voided</option>
	</select>
	{#if !orders}
		<p>No orders found</p>
		{:else}
		<table>
		<thead>
			<tr>
				<th>Use</th>
				<th>Customer Name</th>
				<th>Placed By</th>
				<th>Date Placed</th>
				<th>Date Due</th>
				<th>PO No.</th>
				<th>Strain - Quantity</th>
				<th>Strain Release Date</th>
				<th>Status</th>
				<th>Change Status</th>
				<th>Order Notes</th>
			</tr>
		</thead>
		<tbody>
			{#each orders as order}
				<tr>
					<td>{order.use}</td>
					<td>{order.customerName}</td>
					<td>{order.placedBy}</td>
					<td>{order.datePlaced}</td>
					<td>{order.dateRequired}</td>
					<td>{order.orderNumber}</td>
					<td>{order.clones.map((clone) => clone.name + ' - ' + clone.quantity)}</td>
					<td>{getMonthName(Number(order.clones[0].date.split('/')[0])) + ' - ' + order.clones[0].date.split('/')[1]}</td>
					<td>{order.status}</td>
					<td>
						<select
							id="status"
							on:change={async (e) => {
								let reason;
								if(e.target.value === order.status) return;
								if(e.target.value === 'Voided') {
									reason = prompt('Please enter a reason for voiding this order');
									 if(reason === null) {
										 e.target.value = order.status;
										 return
									 }
									 if(order.notes === '') {
									 	await updateOrderNotes(order.orderNumber, order.notes + 'Voided by: ' + user.email + ' for: \n\n' + reason, user.apikey, data.fetch)
									 } else {
										await updateOrderNotes(order.orderNumber, order.notes + '\n\nVoided by: ' + user.email + ' for: \n\n' + reason, user.apikey, data.fetch)
									 }
									await setOrderStatus(order.orderNumber, e.target.value, user.apikey, data.fetch, reason)
									await loadOrders();
								} else {
									await setOrderStatus(order.orderNumber, e.target.value, user.apikey, data.fetch)
									await loadOrders();
							}}}
						>

							{#each ['1 - Subculture', '2 - Multiplication', '3 - Transplant', '4 - Rooting', '5 - Finished Product', '6 - FulFilled'] as status}
								<option value={status} selected={order.status === status}>{status}</option>
							{/each}
							{#if user.role.includes('admin')}
								<option value="Voided" selected={order.status === 'Voided'}>Voided</option>
							{/if}
						</select>
					</td>
					<td
						><textarea
							id="notes"
							bind:value={order.notes}
							on:change={async () => {
								if(order.notes === '') {
								await updateOrderNotes(order.orderNumber, '', user.apikey, data.fetch)
								} else {
								await updateOrderNotes(order.orderNumber, order.notes, user.apikey, data.fetch)
							}}}
						rows="5"/>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
	{/if}
</main>

<style>
	main {
		padding: 1rem;
	}

	table {
		width: 100%;
		border-collapse: collapse;
	}

	th,
	td {
		padding: 0.25rem;
		border: 1px solid #ccc;
	}

	th {
		background: #eee;
	}

	tr:nth-child(even) {
		background: #fafafa;
	}

	tr:hover {
		background: #f4f4f4;
	}

	button {
		background-color: #4caf50;
		border: none;
		color: white;
		padding: 15px 32px;
		margin: 4px 2px;
		cursor: pointer;
	}

	td {
		text-align: center;
	}
</style>
