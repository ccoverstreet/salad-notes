<script>
	import { searchByName, searchByTags } from "$lib/api.ts";

	let nameFrag;
	let curTags;
	let centererElem;
	let initialInput

	let results = [];
	export let selectedUID = "";
	export let selectedDoc = undefined;

	async function searchByNameHandler(event) {
		console.log(event);

		const res = await searchByName(nameFrag);

		results = [];
		results = res;
	}

	async function searchByTagsHandler(event) {
		event.preventDefault();
		const tags = curTags.split(/[ ,\n]+/);

		const res = await searchByTags(tags);

		results = [];
		results = res;
	}

	function setCurrentDocument(doc) {
		console.log(doc);

		selectedUID = doc.uid;
		selectedDoc = doc;
	}

	function keydownHandler(event) {
		console.log(event);
		if (event.code == "Escape") {
			centererElem.style.display = "none";
			results = [];
		}

	}

	export function show() {
		centererElem.style.display = "flex";
		initialInput.focus();
	}

</script>

<div id="centerer" bind:this={centererElem} on:keydown={keydownHandler} tabindex="0">
	<div>
		<div id="file-search" class="has-background-white">
			<form on:submit={searchByNameHandler}>
				<label>Search by Name</label>
				<div>
					<input class="input" bind:this={initialInput} bind:value={nameFrag}/>
					<button class="button has-background-primary has-text-light">Search</button>
				</div>
			</form>
			<form on:submit={searchByTagsHandler}>
				<label>Search by Tags</label>
				<div>
					<input class="input" bind:value={curTags}/>
					<button class="button has-background-primary has-text-light">Submit Tags</button>
				</div>
			</form>

			<ul>
			{#each results as d}
				<li class="search-entry">
					<button class="file-link-button has-background-white" on:click={setCurrentDocument(d)}>{d.name}</button>
					<div>
						{#each d.tags as t}
							<p>{t}</p>
						{/each}
					</div>
				</li>
			{/each}
			</ul>
		</div>
	</div>
</div>

<style>
	#centerer {
		position: fixed;
		top: 0;
		left: 0;
		/*display: flex;*/
		display: none;
		align-items: center;
		justify-content: center;
		width: 100vw;
		height: 100vh;
		z-index: 1000;
		margin: auto;
		background-color: rgba(0, 0, 0, 0.3)	
	}

	#file-search {
		z-index: 1000;
		padding: 1rem;
		min-width: 30ch;
		max-width: 60ch;
		width: 100%;
	}

	form > div {
		display: flex;
	}

	.file-link-button {
		border: 0;
	}

	.search-entry {
		display: flex;
		text-align: center;
		min-height: 2rem;
	}

	.search-entry > button {
		display: flex;
		text-align: center;
		align-items: center;
		width: 20ch;
		min-width: 20ch;
	}

	.search-entry > div {
		display: flex;
		font-weight: bold;
		flex-wrap: wrap;
		gap: 0.25rem;
		margin: 0.25rem;
	}

	.search-entry > div > p{ 
		display: flex;
		font-weight: bold;
	}
</style>
