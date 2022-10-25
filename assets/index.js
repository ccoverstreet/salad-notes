salad = {
	connectClient: () => {
		const conn = new WebSocket(`ws://${document.location.host}/saladnotes/connectClient`);
		salad.WSCONNECTION = conn;
		conn.addEventListener("message", msg => {
			rawData = msg.data;
			console.log("ASDKASDALDJA");
			data = JSON.parse(rawData);
			salad.retrieveFile(data["modifiedFile"]);
			console.log(data);
		})
	},

	WSCONNECTION: undefined,
	CURRENTFILE: undefined,
	CURRENTFILERAWNODE: undefined,

	retrieveFile: (file) => {
		PRIORFILE = salad.CURRENTFILE
		salad.CURRENTFILE = file;

		const isSameFile = PRIORFILE === salad.CURRENTFILE;

		fetch(file)
			.then(async data => {
				const rawHTML = await data.text();

				const splitFilePath = file.split("/");
				console.log(splitFilePath);
				const relRoot = splitFilePath
					.slice(0, splitFilePath.length - 1)
					.join("/");

				const temp = salad.resolveRelativePaths(
					salad.createHTMLTemplate(
						rawHTML.replaceAll("\\AA", "\\unicode{x212B}")),
					relRoot);

				const pane = document.querySelector("#md-view-pane");

				// Find location of first change
				// Find the deepest first changed element
				var firstMismatchElem = undefined;
				if (isSameFile) {
					const origChildren = salad.CURRENTFILERAWNODE.children;
					const newChildren = temp.content.children;

					console.log(origChildren);
					console.log(newChildren);

					traverseN = Math.min(origChildren.length, newChildren.length);

					for (var i = 0; i < traverseN; i++) {
						// Search for mismatch
						// return a reference
						if (!origChildren[i].isEqualNode(newChildren[i])) {
							console.log("Mismatch found", i, origChildren[i].innerHTML, newChildren[i].innerHTML);
							firstMismatchElem = newChildren[i];
							break;
						}
					}
				}

				const priorRawNode = salad.CURRENTFILERAWNODE
				salad.CURRENTFILERAWNODE = temp.content.cloneNode(true);

				pane.innerHTML = "";
				pane.appendChild(temp.content);

				await MathJax.typesetPromise();

				if (isSameFile && firstMismatchElem) {
					console.log("Auto-scrolling to change")
					/*
					const holder = document.querySelector("#md-view-holder");
					holder.scrollTop = holder.scrollHeight;
					*/
					firstMismatchElem.scrollIntoView(true);
				}
			});
	},

	createHTMLTemplate: (htmlText) => {
		const frag = document.createElement("template");
		frag.innerHTML = htmlText;
		return frag;
	},

	resolveRelativePaths: (template, relpath) => {
		const imgs = template.content.querySelectorAll("img")

		for (const img of imgs) {
			const rawSrc = img.getAttribute("src")

			if (rawSrc.startsWith("./")) {
				img.src = relpath + "/" + rawSrc;
			}

			img.style.width = "100%";
		}
		
		return template;
	}

}
