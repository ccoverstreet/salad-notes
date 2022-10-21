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

	retrieveFile: (file) => {
		console.log(file);

		fetch(file)
			.then(async data => {
				const rawHTML = await data.text();

				const splitFilePath = file.split("/");
				console.log(splitFilePath);
				const relRoot = splitFilePath.slice(0, splitFilePath.length - 1).join("/");
				
				const temp = salad.resolveRelativePaths(salad.createHTMLTemplate(rawHTML),
					relRoot);

				const pane = document.querySelector("#md-view-pane");
				pane.innerHTML = "";
				pane.appendChild(temp.content);
				MathJax.typeset();
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
