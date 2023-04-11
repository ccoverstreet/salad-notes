<script>
	import { onMount } from "svelte";
	import FileBrowser from "$lib/FileBrowser.svelte";
	import Editor from "$lib/Editor.svelte";
	import { newDoc } from "$lib/api.ts";

	let editor = undefined;
	let selectedDoc = undefined;
	let browserElemShow;

	$: {
		if (selectedDoc) {
			//editor.open(selectedDoc);
		} 
	}


	function showBrowser() {
		browserElemShow();
	}

	function createNote() {
		newDoc().then(newDoc => {
			selectedDoc = newDoc;
		})
	}
</script>

<div class="has-background-dark" style="display:flex; justify-content: center; align-items: center;">
	<button class="button has-background-primary has-text-light"
		 style="margin-left: 0.25rem; margin-right: 0.25rem; font-weight: bold; "
		on:click={createNote}>+</button>
	<button class="button has-background-primary has-text-light"
		 on:click={showBrowser}
		 style="">Search</button> 
	<div style="flex-grow: 1;"></div>
	<div class="has-text-primary-light" style="display: flex; align-items: center; font-size: 1.5rem; 
		font-weight: bold; text-align: center">SaladNotes</div>
	<div style="flex-grow: 1;"></div>
</div>
<FileBrowser bind:show={browserElemShow} bind:selectedDoc={selectedDoc}/>
<div>
	<Editor bind:curDoc={selectedDoc}/>
</div>

<style>
	#wrapper {
		display: wrap;
		flex-wrap: wrap;
		height: 80vh;
	}
</style>


