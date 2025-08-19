<script>
	import { onMount } from "svelte";
	import { createItem, updateItemContent, getItemContent } from "$lib/api.js";

	import DocumentSearch from "$lib/DocumentSearch.svelte";

	import { basicSetup } from "codemirror";
	import { EditorView, keymap } from "@codemirror/view";
	import { EditorState } from "@codemirror/state";
	import { vim } from "@replit/codemirror-vim";
	import { markdown } from "@codemirror/lang-markdown";
	import { languages } from "@codemirror/language-data";
	import { indentWithTab } from "@codemirror/commands"; 


	let viewTarget;
	let view = undefined;
	let itemInfo = null;

	let showSearch = false;
	let selectedItemInfo = null;

	$: {
		if (view) {
			selectedItemInfo // For reactivity
			if (view.state.doc.toString().length > 0) {
				saveDocument()
			}

			if (selectedItemInfo) {
				getItemContent(selectedItemInfo.id).
					then(res => res.text()).
					then((text) => {
						console.log(text);
						view.dispatch({
							changes: {
								from: 0,
								to: view.state.doc.toString().length,
								insert: ""
							}
						});
						view.dispatch({
							changes: {
								to: 0,
								from: 0,
								insert: text
							}
						});
						itemInfo = selectedItemInfo;
					}).
					catch(err => console.log(error));
			}
		}
	}

	function newDocument() {
		itemInfo = null;
		selectedItemInfo = null;

		// There may be some unsaved content
		saveDocument()


		view.dispatch({
			changes: {
				from: 0,
				to: view.state.doc.toString().length,
				insert: ""
			}
		});

		itemInfo = null;
	}

	async function saveDocument() {
		// Check if we need to make a new entry
		if (itemInfo === null) {
			const nameResult = prompt("Enter a name for the document:", "New Note");
			if (nameResult === null) return;

			const createReq = await createItem(nameResult, ".md");
			itemInfo = createReq;

		}

		updateItemContent(itemInfo.id, view.state.doc.toString())
	}

	async function saladPasteHandler(event, view) {
		console.log(event, view, event.clipboardData.dataTransfer);

		const cursor = view.state.selection.main.head;

		// If event is just a normal text paste, we just insert the text
		// and move on
		if (event.clipboardData.files.length == 0) {
			const newText = event.clipboardData.getData("text/plain");
			view.dispatch({
				changes: {
					from: cursor,
					to: cursor, 
					insert: newText
				},
				selection: {
					anchor: cursor + newText.length
				}
			})

			return;
		}


		// If file event, we want to make a link to the document
		for (const f of event.clipboardData.files) {
			console.log("FILE", f);

			if (!(f.type in mimeTypeMap)) {
				console.error("MIME Type paste not supported")
				return;
			}

			const ext = mimeTypeMap[f.type];

			const pastedName = prompt("Name pasted image:", "Pasted Image");
			if (pastedName === null) return;

			const createReq = await createItem(pastedName, ext)
			updateItemContent(createReq.id, f)

			const imageText = `![${pastedName}](saladbowl/${createReq.id})`
			
			view.dispatch({
				changes: {
					from: cursor,
					to: cursor,
					insert: imageText
				},
				selection: {
					anchor: cursor + imageText.length
				}
			})
		}
	}

	const mimeTypeMap = {
		"image/png": ".png"
	};

	onMount(() => {
		const eventHandlers = EditorView.domEventHandlers({
			async paste(event, view) {
				saladPasteHandler(event, view);
			}
		});

		let state = EditorState.create({
			extensions: [
				vim(),
				basicSetup, 
				markdown({codeLanguages: languages}),
				keymap.of([indentWithTab, {
					key: "Ctrl-s",
					run() { console.log("Saving"); saveDocument(); return true; }
				}]),
				eventHandlers
				//////////////////				EditorView.clipboardInputFilter.of((text, state) => { console.log(text, state); return "ASDASDASD" })
			]
		})


		view = new EditorView({
			state,
			parent: viewTarget,
			viewportMargin: Infinity,
		});

	})
</script>

<div>
	<button on:click={() => showSearch = true}>Open Document</button>
</div>

<DocumentSearch bind:displaySearch={showSearch} bind:selectedItem={selectedItemInfo}/>


<div>
	<p>{itemInfo ? itemInfo.name : "Untitled document"}</p>
	<button on:click={newDocument}>New Document</button>
</div>
<div bind:this={viewTarget} style="width: 100%; height: 100vh;"></div>
