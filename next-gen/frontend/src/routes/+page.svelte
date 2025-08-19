<script>
	import Editor from "$lib/Editor.svelte";

	import { getMarkdownAsHTML } from "$lib/api.js";

	const mimeTypeMap = {
		"image/png": ".png"
	};

	async function pasteHandler(event) {
		console.log("WORKING HANDLER", event.clipboardData.files);
		for (const f of event.clipboardData.files) {
			console.log(f);

			if (!(f.type in mimeTypeMap)) {
				console.error("MIME Type not supported")
				return;
			}

			const ext = mimeTypeMap[f.type];

			const createReq = await fetch("/api/createItem", {
				method: "POST",
				body: JSON.stringify({
					name: "Pasted image",
					extension: ext
				})
			}).then(res => res.json())

			console.log(createReq)

			let formData = new FormData();
			formData.append("file", f);
			formData.append("itemId",  createReq.id);
			console.log(formData)

			fetch("api/updateItemContent", {
				method: "POST",
				body: formData
			})
		}

	}

	async function mdSaveHandler(event) {
		console.log(event);

		const createReq = await fetch("/api/createItem", {
			method: "POST",
			body: JSON.stringify({
				name: "Pasted image",
				extension: ".md"
			})
		}).then(res => res.json())

		console.log(event.target.value)

		let formData = new FormData();
		formData.append("file", new Blob([event.target.value]));
		formData.append("itemId",  createReq.id);
		console.log(formData)

		fetch("api/updateItemContent", {
			method: "POST",
			body: formData
		})
	}


	function tester() {
		getMarkdownAsHTML("67e9fded-c336-4a5a-b9e0-47a2b3c94025").
			then(res => res.text()).
			then(text => console.log(text));
	}

</script>


<Editor/>


<button on:click={tester}>TESTER</button>
