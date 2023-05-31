<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { getClones } from '$lib/functions/clones/getClones';
	import { deleteClone } from '$lib/functions/clones/deleteClone';
	import { editCloneName } from '$lib/functions/clones/editCloneName';
	import { setCloneQuantity } from '$lib/functions/clones/setCloneQuantity';
	import { getMonthName } from '$lib/functions/getMonthName';
	import type { Clone, User } from '$lib/types';

	export let data;
	let clones: Clone[] = [],
		user =
			data.user ||
			({
				apikey: ''
			} as User);

	onMount(async () => {
		clones = await getClones(user.apikey, data.fetch);
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
	<title>Manage Clones</title>
	<meta name="description" content="Manage Clones Page" />
</svelte:head>

<main>
	<h1>Manage Clones</h1>
	<p>Manage Clones Page</p>
	<button on:click={async () => await goto('./')}>Back</button>
	<button on:click={async () => await goto('clones/new')}>Create New Clone</button>

	<table>
		<tr>
			<th>Clone Name</th>
			<th>Month</th>
			<th>Year</th>
			<th>Clone Quantity</th>
			<th>Actions</th>
		</tr>

		{#each clones as clone}
			<tr>
				<td>
					<textarea placeholder="Clone Name" bind:value={clone.name} />
					<button
						on:click={async () =>
							await editCloneName(clone.cloneId, clone.name, user.apikey, data.fetch)}>Save</button
					>
				</td>
				<td>{getMonthName(Number(clone.date.split('/')[0]))}</td>
				<td>{clone.date.split('/')[1]}</td>
				<td>
					<input type="number" min="0" placeholder="0" bind:value={clone.quantity} />
					<button
						on:click={async () =>
							await setCloneQuantity(clone.cloneId, clone.quantity, user.apikey, data.fetch)}
						>Save</button
					>
				</td>
				<td>
					<button
						on:click={async () => {
							await deleteClone(clone.cloneId, user.apikey, data.fetch);
							clones = clones.filter((c) => c.cloneId !== clone.cloneId);
						}}>Delete</button
					>
				</td>
			</tr>
		{/each}
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
</style>
