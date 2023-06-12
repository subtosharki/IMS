<script lang="ts">
	import { onMount } from 'svelte';
	import { getCustomers } from '$lib/functions/customers/getCustomers';
	import type { Customer, User } from '$lib/types';
	import { goto } from '$app/navigation';
	import { deleteCustomer } from '$lib/functions/customers/deleteCustomer';
	import { updateCustomerNotes } from '$lib/functions/customers/updateCustomerNotes';
	import {getUserByAPIKey} from "$lib/functions/users/getUserByAPIKey";
	import {decrypt} from "$lib/functions/utils/crpyt";

	export let data;
	let customers = [] as Customer[],
		user = {} as User;

	onMount(async () => {
		user = await getUserByAPIKey(decrypt(data.apikey!), data.fetch!)
		customers = await getCustomers(user.apikey, data.fetch!);
	});
</script>

<svelte:head>
	<title>Manage Customers</title>
	<meta name="description" content="Manage Customers Page" />
</svelte:head>

<main>
	<h1>Manage Customers</h1>
	<button on:click={async () => await goto('./')}>Back</button>
	<button on:click={async () => await goto('customers/add')}>Add Customer</button>
	<p>Manage Customers Page</p>

	{#if !customers}
		<p>No customers found.</p>
		{:else}
	<table>
		<thead>
			<tr>
				<th>Name</th>
				<th>Email</th>
				<th>Sub. Id</th>
				<th>Phone #</th>
				<th>Address</th>
				<th>Notes</th>
				<th>Actions</th>
			</tr>
		</thead>
		<tbody>
			{#each customers as customer}
				<tr>
					<td>{customer.name}</td>
					<td>{customer.email}</td>
					<td>{customer.subscriptionId}</td>
					<td>{customer.phone}</td>
					<td>{customer.address}</td>
					<td
						><textarea
							id="notes"
							bind:value={customer.notes}
							on:change={async () =>
								await updateCustomerNotes(
									customer.customerId,
									customer.notes,
									user.apikey,
									data.fetch
								)}
						/></td
					>
					<td>
						<button
							on:click={async () => {
								await deleteCustomer(customer.customerId, user.apikey, data.fetch);
								customers = customers.filter((c) => c.customerId !== customer.customerId);
							}}
						>
							Delete
						</button>
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
</style>
