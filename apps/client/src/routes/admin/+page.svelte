<script lang="ts">
	import { goto } from '$app/navigation';
	import type { User } from '$lib/types';

	export let data: App.PageData;
	let user =
		data.user ||
		({
			email: ''
		} as User);
</script>

<svelte:head>
	<title>Admin Panel</title>
	<meta name="description" content="Admin Panel" />
</svelte:head>

<main>
	<h1>Admin Panel</h1>
	<p>Welcome, {user.email}</p>
	<button on:click={async () => await goto('/dashboard')}>Back</button>

	<button on:click={async () => await goto('/admin/users')}>Manage Users</button>
	<button on:click={async () => await goto('/admin/clones')}>Manage Clones</button>
	<button on:click={async () => await goto('/admin/customers')}>Manage Customers</button>
</main>

<style>
	main {
		text-align: center;
	}

	button {
		margin: 10px;
	}

	h1 {
		margin: 0;
	}

	p {
		margin: 0;
	}
</style>
