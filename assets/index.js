
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

				const viewer = document.querySelector("#salad-md-viewer-1");

				viewer.displayFile(file, rawHTML);
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
	},

	replaceSpecialChars: (text) => {
		return text.replaceAll("\\AA", "\\unicode{x212B}")
	},

	// Find the first element in a tree that is new compared 
	// to an existing tree.
	// Returns the element in the new tree
	getFirstDifferentElem: (old, updated) => {
		console.log(old, updated);
		const origChildren = old.children;
		const newChildren = updated.children;
		const traverseN = Math.min(origChildren.length, newChildren.length);

		for (var i = 0; i < traverseN; i++) {
			if (!origChildren[i].isEqualNode(newChildren[i])) {
				if (newChildren[i].children.length === 0) {
					return newChildren[i]
				}

				return salad.getFirstDifferentElem(origChildren[i], newChildren[i]);
			}
		}

		if (newChildren.length > origChildren.length) {
			return newChildren[origChildren.length];
		}

		return null;
	}

}



class SaladMDViewer extends HTMLElement {
	constructor() {
		super();

		this.currentFilename = undefined;
		this.currentFileNodeRaw = undefined;
	}

	connectedCallback() {
		console.log("ASDALSJDHALKSJDHLAKSJh");
		this.innerHTML = `
		<div class="salad-md-viewer-filename"></div>

		<div class="salad-md-viewer-content" style="height: 500px; overflow: scroll;">
		</div>
		`

		const filenameElem = this.querySelector(".salad-md-viewer-filename");
		filenameElem.onclick = () => {
			navigator.clipboard.writeText(filenameElem.textContent);
		}
	}

	async displayFile(filename, htmlText) {
		const isSameFile = this.currentFilename = filename;
		this.currentFilename = filename;

		const contentView = this.querySelector(".salad-md-viewer-content");
		
		const splitFilePath = filename.split("/");
		const relRoot = splitFilePath
			.slice(0, splitFilePath.length - 1)
			.join("/");

		console.log(`Relative root for ${filename}: ${relRoot}`);

		const newTemplate = salad.resolveRelativePaths(
			salad.createHTMLTemplate(
				salad.replaceSpecialChars(htmlText)
			),
			relRoot
		);

		const firstDiffElem = this.currentFileNodeRaw ? salad.getFirstDifferentElem(
			this.currentFileNodeRaw, newTemplate.content
		) : null;

		this.currentFileNodeRaw = newTemplate.content.cloneNode(true);

		// Add to view as child and render MathJax
		contentView.innerHTML = "";
		contentView.appendChild(newTemplate.content);

		this.querySelector(".salad-md-viewer-filename").textContent = filename;

		await MathJax.typesetPromise();

	 	if (isSameFile && firstDiffElem) {
			console.log("Auto-scrolling to change");
			console.log(`First different element`, firstDiffElem);

			// Not sure why this is required for correct behavior
			setTimeout(() => {
				firstDiffElem.scrollIntoView();
			}, 0)

			firstDiffElem.style.backgroundColor = "var(--clr-secondary)";
			setTimeout(() => {
				firstDiffElem.style.backgroundColor = "#ffffff";
			}, 3000)
	 	}

	}
}

customElements.define("salad-md-viewer", SaladMDViewer);
