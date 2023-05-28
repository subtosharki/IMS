<script lang="ts">
	import { onMount } from 'svelte';
	import { getUsers } from '$lib/functions/users/getUsers';
	import { goto } from '$app/navigation';
	import { toggleAdmin } from '$lib/functions/users/toggleAdmin';
	import { deleteUser } from '$lib/functions/users/deleteUser';
	import type { User } from '$lib/types';

	export let data: App.PageData;
	let users: User[] = [],
		user =
			data.user ||
			({
				apikey: ''
			} as User);

	onMount(async () => {
		users = await getUsers(user.apikey, data.fetch!);
	});
</script>

<svelte:head>
	<title>Manage Users</title>
	<meta name="description" content="Manage Users Page" />
</svelte:head>

<main>
	<h1>Manage Users</h1>
	<button on:click={async () => await goto('./')}>Back</button>
	<button on:click={async () => await goto('/users/add')}>Add User</button>
	<p>Manage Users Page</p>

	<table>
		<thead>
			<tr>
				<th>Email</th>
				<th>Password</th>
				<th>API Key</th>
				<th>Admin</th>
				<th>Actions</th>
			</tr>
		</thead>
		<tbody>
			{#each users as loopedUser}
				<tr>
					<td>{loopedUser.email}</td>
					<td>{loopedUser.password}</td>
					<td>{loopedUser.apikey}</td>
					<td>{loopedUser.admin}</td>
					<td>
						<button
							on:click={async () => {
								await toggleAdmin(loopedUser.userId, user.apikey, data.fetch);
								users.forEach((user) => {
									if (user.userId === loopedUser.userId) {
										user.admin = !user.admin;
									}
								});
								users = users;
							}}>Toggle Admin</button
						>
						<button
							on:click={async () => {
								await deleteUser(loopedUser.userId, user.apikey, data.fetch);
								users = users.filter((user) => user.userId !== loopedUser.userId);
							}}>Delete</button
						>
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
</style>
