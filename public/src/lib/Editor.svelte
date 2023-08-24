<script>
	import  Display from "$lib/Display.svelte";
	import { onMount } from "svelte";
	import { updateMeta, deleteItem, getFileByUID, getRenderedMarkdown } from "$lib/api.ts";

	let aceEditor;
	export let curDoc = undefined;
	let curUID = "";
	let curName = "";
	let curTags = "";

	let displayElem;

	let displayContent;

	$: {
		if (curDoc) {
			getFileByUID(curDoc.uid)
				.then(data => {
					aceEditor.resize();
					aceEditor.setValue(data);
					aceEditor.clearSelection();
				});

			//curUID = curDoc.uid;
			//curName = curDoc.name;
			//curTags = curDoc.tags.join(" "); 
		} 

		setMeta();
	}

	function setMeta() {
		if (curDoc) {
			curUID = curDoc.uid;
			curName = curDoc.name;
			curTags = curDoc.tags.join(" ");
		} else {
			curName = "";
			curTags = "";
			if (aceEditor) {
				aceEditor.setValue("");
			}
		}
	}


	onMount(async () => {
    	aceEditor = ace.edit("editor");
    	aceEditor.setTheme("ace/theme/clouds");
    	aceEditor.session.setMode("ace/mode/markdown");
		aceEditor.session.setUseWrapMode(true);
		aceEditor.setKeyboardHandler("ace/keyboard/vim");
		aceEditor.setFontSize(14);

		document.getElementById('editor').style.fontSize='18px';
		aceEditor.commands.addCommand({
  		  	name: 'saveFile',
  		  	bindKey: {
    			win: 'Ctrl-S', mac: 'Command-S',
    			sender: 'editor|cli'
  		  	},
  		  	exec: function (env, args, request) {
  		  		const content = aceEditor.getValue();
  		  		const headers = {
  		  			"Content-Type" :"application/json",
  		  			"saladnotes-uid": curUID
  		  		};

				fetch("/api/writeItem", {
					method: "POST",
					headers: headers,
					body: content
				})
					.then(res => {

					});

    			console.log('saving...', env, args, request);
				console.log("Content:", aceEditor.getValue());

				getRenderedMarkdown(curUID)
					.then(html => {
						displayContent = html;
					})
  		  	}
		});

	})


	function updateMetaHandler(event) {
		event.preventDefault();
	
		updateMeta({
			uid: curUID,
			name: curName,
			tags: curTags.split(/[ ,\n]+/)
		})
	}

	function deleteDocHandler(event) {
		console.log("Delete Handler")
		if (confirm("Are you sure you want to delete this document?")) {
			deleteItem(curUID);
			curDoc = undefined;
		}
	}


</script>

<div id="editor-content">
	<div style="border: solid black 1px;">
		<div style="display: flex;">
			<form on:submit={updateMetaHandler} style="width: auto; flex-grow: 1;">
				<div>
					<label>Name</label>
					<input bind:value={curName} on:focusout={updateMetaHandler}/>
				</div>

				<div style="flex-grow: 1;">
					<label>Tags</label>
					<input style="width: 100%;" bind:value={curTags} on:focusout={updateMetaHandler}/>
				</div>




				<button style="display: none;"></button>
			</form>

			<div style="display: flex; align-items: center; margin: 0em 0.25rem;">
				<button on:click={deleteDocHandler} class="button has-background-danger has-text-light"
					style="font-weight: bold;">X</button>
			</div>
		</div>

		<div id="editor"></div>
	</div>

	<Display bind:content={displayContent}/>
</div>

<style>

	#editor {
		width: 100%;
		height: calc(80vh - 3.5rem);
	}

	#editor-content {
		display: flex;
		width: 100%;
		height: 80vh;
		flex-wrap: wrap;
	}

	#editor-content > * {
		min-width: 30ch;
		width: 50%;
		height: 100%;
		flex-grow: 1;
	}

	#editor-output {
		padding: 1rem;
		overflow-y: scroll;
	}

	form {
		display: flex;
	}

	form > div {
		display: flex;
		flex-direction: column;
		padding: 0.25rem;
	}

</style>
