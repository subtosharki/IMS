<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { getClones } from '$lib/functions/clones/getClones';
	import { getMonthName } from '$lib/functions/utils/getMonthName.js';
	import type { Clone, User } from '$lib/types';
	import {getUserByAPIKey} from "$lib/functions/users/getUserByAPIKey";
	import {decrypt} from "$lib/functions/utils/crpyt";

	export let data;
	let clones: Clone[] = [],
		user = {
			email: '',
			apikey: '',
			role: ''
		} as User;
	onMount(async () => {
		user = await getUserByAPIKey(decrypt(data.apikey), data.fetch!)
		clones = await getClones(user.apikey, data.fetch!);
		if(clones.length === 0) return;
		clones.sort((a, b) => {
			if (a.date.split('/')[1] > b.date.split('/')[1]) return 1;
			if (a.date.split('/')[1] < b.date.split('/')[1]) return -1;
			if (a.date.split('/')[0] > b.date.split('/')[0]) return 1;
			if (a.date.split('/')[0] < b.date.split('/')[0]) return -1;
			return 0;
		})
	});
</script>

<svelte:head>
	<title>Dashboard</title>
	<meta name="description" content="Dashboard" />
</svelte:head>

<main>
	<h1>Dashboard</h1>
	<p>Welcome, {user.email} to the dashboard</p>
	<button on:click={async () => await goto('/order')}>Place Order</button>
	<button on:click={async () => await goto('/orders')}>View Orders</button>
	{#if user.role.includes('admin')}
		<button on:click={async () => await goto('/admin')}>Admin</button>
	{/if}
	<button style="background-color: #ff3e00" on:click={async () => await goto('/logout')}
		>Logout</button
	>
	<h2>Clones</h2>
	{#if clones.length === 0}
		<p>No clones found</p>
	{:else}
	<table>
		<thead>
			<tr>
				<th>Cultivar</th>
				<th>Month</th>
				<th>Year</th>
				<th>Quantity</th>
			</tr>
		</thead>
		<tbody>
			{#each clones as clone}
				<tr>
					<td>{clone.name}</td>
					<td>{getMonthName(Number(clone.date.split('/')[0]))}</td>
					<td>{clone.date.split('/')[1]}</td>
					<td>{clone.quantity}</td>
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
