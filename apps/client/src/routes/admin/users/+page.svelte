<script lang="ts">
	import { onMount } from 'svelte';
	import { getUsers } from '$lib/functions/users/getUsers';
	import { goto } from '$app/navigation';
	import { toggleAdmin } from '$lib/functions/users/toggleAdmin';
	import { deleteUser } from '$lib/functions/users/deleteUser';
	import type { User } from '$lib/types';
	import {decrypt} from "$lib/functions/crpyt";
	import {getUserByAPIKey} from "$lib/functions/users/getUserByAPIKey";

	export let data;
	let users: User[] = [],
		user =
			{
				apikey: ''
			} as User;

	onMount(async () => {
		user = await getUserByAPIKey(decrypt(data.apikey), data.fetch!)
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
	<button on:click={async () => await goto('users/add')}>Add User</button>
	<p>Manage Users Page</p>

	<table>
		<thead>
			<tr>
				<th>Name</th>
				<th>Email</th>
				<th>Password</th>
				<th>Role</th>
				<th>Actions</th>
			</tr>
		</thead>
		<tbody>
			{#each users as loopedUser}
				<tr>
					<td>{loopedUser.firstName + ' ' + loopedUser.lastName}</td>
					<td>{loopedUser.email}</td>
					<td>{loopedUser.password}</td>
					<td>{loopedUser.role}</td>
					<td>
						<button
							on:click={async () => {
								await toggleAdmin(loopedUser.userId, user.apikey, data.fetch);
								users.forEach((user) => {
									if (user.userId === loopedUser.userId) {
										if(user.role === 'admin') {
											user.role = 'sales';
										} else {
											user.role = 'admin';
										}
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
