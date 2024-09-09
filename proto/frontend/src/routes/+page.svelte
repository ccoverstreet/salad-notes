<script>
	import { onMount } from "svelte";
	import { EditorView, basicSetup } from "codemirror";
	import {EditorState, Compartment} from "@codemirror/state";
	import {javascript} from "@codemirror/lang-javascript";
	import {markdown} from "@codemirror/lang-markdown";
	import {languages} from "@codemirror/language-data";
	import { indentWithTab } from "@codemirror/commands";
	import{ keymap } from "@codemirror/view";
	//import {python} from "@codemirror/lang-python"
	import { vim } from "@replit/codemirror-vim"

	import FileSelector from "$lib/FileSelector.svelte";

	let language = new Compartment;
	let tabsize = new Compartment;
	let view = null;
	let fileDetails = {};

	function openFileHandle(details, fileContent) {
		console.log(view);
		console.log(fileContent);
		fileDetails = details
		view.dispatch({ changes: {from: 0, to: view.state.doc.length, insert: fileContent}});
	}

	function saveFile() {
		console.log(view.state.doc.toString());
	}

	onMount(() => {
		console.log("Hello");

		//var editor = ace.edit("editor");

		let state = EditorState.create({
			extensions: [
				vim(),
				basicSetup, 
				markdown({codeLanguages: languages}),
				keymap.of([indentWithTab, {
					key: "Ctrl-s",
					run() { console.log("Saving"); saveFile(); return true; }
				}])
			]
		})

		view = new EditorView({
			state,
			parent: document.getElementById("editor-container"),
			viewportMargin: Infinity,
		});


		console.log("ASD")
	})
</script>
<h1>Salad Notes</h1>

<FileSelector openFileHandle={openFileHandle}/>

<div id="main-content">
	<div id="editor-container">
	</div>
	<div>PREVIEW</div>
</div>

<style global>
	#main-content {
		display: flex;
	}

	#editor-container {
		height: 50vh;
		width: 50%;
		box-sizing: border-box;
	}

	#editor-container > * {
		height: 50vh;
		width: 500px;
		flex-grow: 1;
	}

</style>
