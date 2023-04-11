<script>
	export let content = "";

	let contentElem;

	$: {
		if (content) {
			updateView()
		}
	}

	async function updateView() {
		console.log(content);
		contentElem.innerHTML = content;	
		
		await MathJax.typesetPromise();
		contentElem.scrollTop = contentElem.scrollHeight;
	}

</script>

<div id="display-content" class="content" bind:this={contentElem}>
</div>

<style>

	#display-content {
		height: 100%;
		width: 50%;
		flex-grow: 1;
		padding: 0.5rem;
		border: solid black 1px;
		overflow-y: scroll
	}

</style>
