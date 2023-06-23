<script lang="ts">
    import type {OrderStatus} from "$lib/types";
    import {setOrderStatus} from "$lib/functions/orders/setOrderStatus";
    import {onMount} from "svelte";
    import {getUserByAPIKey} from "$lib/functions/users/getUserByAPIKey";
    import {decrypt} from "$lib/functions/utils/crpyt";
    import {goto} from "$app/navigation";
    let mode, barcode, user
    export let data;

    onMount(async () => {
        user = await getUserByAPIKey(decrypt(data.apikey), data.fetch)
    })
</script>

<svelte:head>
    <title>Barcode Order Edit</title>
    <meta name="description" content="Edit an order by barcode">
</svelte:head>

<main>
    <button on:click={async () => await goto('./')}>Back</button>
    {#if !mode}
        <select on:change={() => mode = (document.querySelector('select')).value}>
            <option value="0" disabled selected>Select a mode</option>
            <option value="1 - Subculture">1 - Subculture</option>
            <option value="2 - Multiplication">2 - Multiplication</option>
            <option value="3 - Transplant">3 - Transplant</option>
            <option value="4 - Rooting">4 - Rooting</option>
            <option value="5 - Finished Product">5 - Finished Product</option>
            <option value="6 - FulFilled">6 - FulFilled</option>
        </select>
        {:else}
    <h1>Barcode Order Edit</h1>
    <p>Scan a barcode to edit an order</p>
    <p>Mode: {mode}</p>
    <form on:submit={async () => {
        const statuses = ['1 - Subculture', '2 - Multiplication', '3 - Transplant', '4 - Rooting', '5 - Finished Product', '6 - FulFilled']
        if(barcode.length === 1) {
            if(barcode !== '1' && barcode !== '2' && barcode !== '3' && barcode !== '4' && barcode !== '5' && barcode !== '6') {
                alert("Invalid mode")
                return
            }
            statuses.forEach(status => {
                if(status.startsWith(barcode)) {
                    mode = status
                }
            })
            return
        }
        try {
        await setOrderStatus(barcode, mode, user.apikey, data.fetch)
        } catch (e) {
            alert("Unable to find order")
        }
        barcode = ''
    }}>
    <input type="text" id="barcode" name="barcode" autofocus bind:value={barcode}>
    </form>
    {/if}
</main>