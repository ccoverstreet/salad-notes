<script>
	export let openFileHandle;

	let items = [];
	let currentDir = "/";


	function listDirectory() {
		console.log(currentDir);
		fetch("/api/listDirectory", {
			method: "POST",
			body: JSON.stringify({directory: currentDir})
		})
			.then(res => res.json())
			.then(data => {
				console.log(data)
				items = [{name: "..", "dir": undefined, isDir: true}].concat(data);
			})
			.catch(err => {
				console.error(err)
			})
	}

	function simplifyPath(raw) {
		const tokens = raw.split("/");
		console.log("INPUT", raw, tokens);
		const out = [];
		for (var i = 0; i < tokens.length; i++) {
			console.log(out, i, tokens[i]);
			if (tokens[i] === "") continue;
			if (i === 0 && tokens[i] === "..") continue;
			if (tokens[i] === "..") {
				out.pop();
				continue;
			}

			out.push(tokens[i])
		}

		if (out.length === 0) {
			return "/"
		}

		return "/" + out.join("/");
	}

	function openItem(item) {
		console.log(item)
		if (item.isDir) {
			if (item.name == "..") {
				currentDir = simplifyPath(currentDir + "/" + item.name);
			} else {
				currentDir = simplifyPath(currentDir + "/" + item.name);
			}

			listDirectory()
		} else {
			getFile(item.dir + "/" + item.name)
				.then(data => {
					openFileHandle(item, data)
				})
		}
	}

	function getFile(filename) {
		return fetch("/api/getFile", {
			method: "POST",
			body: JSON.stringify({filename: filename})
		})
			.then(res => res.text())
			.catch(err => {
				console.error(err)
			})
	}

</script>

<div>
	<h1>File Selector</h1>
	<button on:click={listDirectory}>Refresh</button>
	<div>
		{#each items as item}
			{#if item.isDir}
				<button on:click={openItem(item)}>
					<p>{item.name}</p>
				</button>
			{:else}
				<button on:click={openItem(item)}>
					<p>{item.name}</p>
				</button>
			{/if}
		{/each}
	</div>
</div>
