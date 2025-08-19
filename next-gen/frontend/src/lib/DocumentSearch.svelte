<script>
	import { getSearchItemsByName } from "$lib/api.js";

	export let selectedItem = undefined;
	export let displaySearch = false

	let searchItems = [];
	let searchContent = "";
	let warningMessage = ""

	function updateSearchItems() {
		getSearchItemsByName(searchContent).
			then(res => res.json()).
			then(json => {
				searchItems = json
			}).
			catch(err => {
				console.log(err);
				warningMessage = err;
			})

	}

	function selectCurrentItemGenerator(itemInfo) {
		return () => {
			selectedItem = itemInfo;
			console.log(selectedItem);
			displaySearch = false;
		}
	}
</script>

{#if displaySearch === true}
	<div style="position: fixed; top: 0; left: 0; background-color: #00000044;
	width: 100vw; height: 100vh; z-index: 100;">
		<div style="display: flex; align-items: center; width: 100%; 
		height: 100%; justify-content: center; flex-direction: column">
			<input on:change={updateSearchItems} bind:value={searchContent}
		  style="width: 100%; max-width: 40rem; margin: 2rem"/>		
			<div style="width: 100%; max-width: 40rem;">
				{#each searchItems as item}
					<div on:click={selectCurrentItemGenerator(item)} class="search-item">
						<div style="width: 100%">{item.name}</div>
						<div>{item.extension}</div>
					</div>
				{/each}
			</div>
		</div>
	</div>
{/if}

<style>
	.search-item {
		display: flex;
		margin: 0.5rem;
	}
</style>
