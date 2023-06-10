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
		voidedOrders: Order[] = [],
		seeVoided = false,
			user = {
				email: '',
				role: ''
			} as User

	async function loadOrders() {
		orders = await getOrders(user.apikey, data.fetch!);
		voidedOrders = orders.filter((order) => order.status !== 'In Progress');
		orders = orders.filter((order) => order.status === 'In Progress');
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
	<h1>{!seeVoided ? 'Pending Orders' : 'Fulfilled/Voided Orders'}</h1>
	<p>Welcome, {user.email} to the orders page</p>
	<button on:click={() => (seeVoided = !seeVoided)}
		>{!seeVoided ? 'View Fulfilled/Voided Orders' : 'View Pending Orders'}</button
	>
	{#if user.role.includes('admin')}
		<button on:click={async () => await downloadOrders(user.apikey, data.fetch)}
			>Download Order Logs</button
		>
	{/if}
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
			{#each !seeVoided ? orders : voidedOrders as order}
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

							{#each ['In Progress', 'Fulfilled'] as status}
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
