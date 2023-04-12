function JSONRequest(url, data) {
	console.log(url, data);
	console.log(JSON.stringify(data));
	return fetch(url, {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify(data)
	})
	.then(async res => {
		if (res.status < 200 || res.status >= 400) {
			throw new Error(await res.text());
		}

		return res.json();
	})
}

export function newDoc() {
	return JSONRequest("/api/addItem", {name: "New Document", tags: []})
}

export function deleteItem(uid) {
	return JSONRequest("/api/deleteItem", {uid: uid});
}

export function getFileByUID(uid) {
	return fetch(`/api/uid/${uid}`)
	.then(res => { return res.text(); });
}

export function getRenderedMarkdown(uid) {
	return fetch(`/api/render/${uid}`)
		.then(res => { return res.text(); });
}

export function searchByName(name) {
	return JSONRequest("/api/getByName", {name: name});
}

export function searchByTags(tags) {
	return JSONRequest("/api/getByTags", {tags: tags});
}

export function updateMeta(meta) {
	return JSONRequest("/api/updateMeta", meta);
}
