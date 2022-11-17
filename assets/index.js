
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
	},

	listDir: (dirName) => {
		return fetch("/saladnotes/listDir", {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({dirName: dirName})
		})
			.then(async data => {
				return await data.json();
			})
	},

	showDir: async (dirName) => {
		files = await salad.listDir(dirName);

		console.log(files);
	}, 

	cleanFilePath: (raw) => {
		const split = raw.split("/")
		const cleaned = [];

		let ind = split.length;
		let skipCounter = 0;
		while(ind > 0) {
			ind--;
			if (split[ind] === "..") {
				skipCounter += 1;
				continue;
			}

			if (skipCounter > 0) {
				skipCounter -= 1;
				continue;
			}

			cleaned.push(split[ind]);
		}

		console.log(cleaned);

		const path = cleaned.reverse().join("/");
		if (path === "") {
			return "."
		}

		return path;
	}
}

class SaladFileExplorer extends HTMLElement {
	constructor() {
		super();
		
		this.currentDir = ".";
	}

	connectedCallback() {
		this.innerHTML = `
		<div class="salad-file-explorer-content">
		</div>
		`;

		this.showDirectory(this.currentDir);
	}

	async showDirectory(dirName) {
		dirName = salad.cleanFilePath(dirName);
		const files = await salad.listDir(dirName);
		if (files.err) {
			console.error(`Unable to list directory ${dirName}`);
			return
		}
		this.currentDir = dirName;
		const content = this.querySelector(".salad-file-explorer-content");


		console.log(files, this);

		const fileHolder = document.createElement("div");

		const aboveDir = document.createElement("div");
		aboveDir.textContent = "..";
		aboveDir.dataset.isDir = true;
		aboveDir.dataset.name = "..";
		aboveDir.style.padding = "0em 1em"
		aboveDir.onclick = (event) => {
			const elem = event.srcElement;

			console.log(elem.dataset);

			// If element is not a directory
			// Just retrieve the file
			if (elem.dataset.isDir !== 'true') {
				console.log("TODO: Retrieve file")
				return
			}

			this.showDirectory(this.currentDir+"/"+elem.dataset.name);
		}

		fileHolder.appendChild(aboveDir);

		for (const f of files) {
			const fileRow = document.createElement("div");
			
			const attributes = f["isDir"] === true ? "/" : "";

			fileRow.textContent = attributes + f.name;
			fileRow.dataset.isDir = f["isDir"];
			fileRow.dataset.name = f["name"];
			fileRow.style.padding = "0em 1em"

			fileRow.onclick = (event) => {
				const elem = event.srcElement;

				console.log(elem.dataset);

				// If element is not a directory
				// Just retrieve the file
				if (elem.dataset.isDir !== 'true') {
					salad.retrieveFile(this.currentDir+"/"+elem.dataset.name);
					console.log("TODO: Retrieve file");
					return
				}

				this.showDirectory(this.currentDir+"/"+elem.dataset.name);
			}

			fileHolder.appendChild(fileRow);
		}

		content.innerHTML = "";
		content.appendChild(fileHolder);
	}
}

customElements.define("salad-file-explorer", SaladFileExplorer);


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

		<div class="salad-md-viewer-content" tabindex="0">
		</div>
		`

		const filenameElem = this.querySelector(".salad-md-viewer-filename");
		filenameElem.onclick = () => {
			navigator.clipboard.writeText(filenameElem.textContent);
		}

		this.onkeydown = event => {
			event.preventDefault();
			console.log(event);
			switch (event.key) {
				case "r":
					location.reload();	
				case "e":
					if (event.ctrlKey) {
						this.scroll(20);
					}
					break;
				case "y":
					if (event.ctrlKey) {
						this.scroll(-20);
					}
					break;
			}
		}
	}

	async displayFile(filename, htmlText) {
		const isSameFile = this.currentFilename === filename;
		this.currentFilename = filename;

		const splitFilename = filename.split("/");

		const ext = splitFilename[splitFilename.length - 1].split(".").slice(-1)[0];
		console.log(ext);

		const contentView = this.querySelector(".salad-md-viewer-content");

		// Code file extentions
		const codeExtensions = ["py", "c", "cpp", "go", "rs"];
		if (codeExtensions.includes(ext)) {
			const pre = document.createElement("pre");
			const code = document.createElement("code");
			code.textContent = htmlText;

			pre.appendChild(code);
			contentView.innerHTML = "";
			contentView.appendChild(pre);

			this.querySelector(".salad-md-viewer-filename").textContent = filename;
		
			hljs.highlightAll();
			return
		}
		
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

		const firstDiffElem = isSameFile ? salad.getFirstDifferentElem(
			this.currentFileNodeRaw, newTemplate.content
		) : null;

		this.currentFileNodeRaw = newTemplate.content.cloneNode(true);

		// Add to view as child and render MathJax
		contentView.innerHTML = "";
		contentView.appendChild(newTemplate.content);

		this.querySelector(".salad-md-viewer-filename").textContent = filename;

		await MathJax.typesetPromise();
		hljs.highlightAll();

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

		contentView.focus();
	}

	scroll(pixels) {
		const content = this.querySelector(".salad-md-viewer-content")
		content.scrollBy({top: pixels, behavior: "smooth"});
	}
}

customElements.define("salad-md-viewer", SaladMDViewer);
