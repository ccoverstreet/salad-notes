<script>
	import MDViewer from "$lib/MDViewer.svelte"

	let inputDiv;
	let textInput;
	let renderedInput;
	let normalMode = true;

	function getCursorColumn(elem) {
		let index = elem.selectionStart;

		let colNumber = 0;
		while (index > 0 && elem.value[index] != "\n") {
			colNumber++;
			index--;
		}
		return colNumber;
	}

	function renderHTML() {
		fetch("/util/markdownToHTML", {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({content: textInput})
		})
			.then(async resp => {
				const data = await resp.json();
				console.log(data);
				renderedInput = data.markdown;
				setTimeout(() => { MathJax.typesetPromise()}, 0);
			})
	}
</script>


<div id="wrapper">
	<textarea bind:this={inputDiv} on:keydown={keydownHandler} bind:value={textInput}></textarea>
	<MDViewer id="viewer" content={renderedInput}
		   style=""></MDViewer>
</div>

<style>
	#wrapper {
		height: 60vh;
		width: 100%;
		display: flex;
		grid-template-columns: 50% 50%;
		flex-wrap: wrap;
		justify-content: center;
	}

	#wrapper > * {
		flex-grow: 1;
		min-width: 30rem;
	}

	#wrapper > textarea {
		tab-size: 4;
	}

</style>


