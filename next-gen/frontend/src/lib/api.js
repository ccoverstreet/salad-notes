
export function createItem(name, extension) {
	return fetch("/api/createItem", {
				method: "POST",
				body: JSON.stringify({
					name: name,
					extension: extension
				})
			}).then(res => res.json())
}

export function updateItemContent(itemId, content) {
	let formData = new FormData();
	formData.append("file", content);
	formData.append("itemId", itemId);

	return fetch("/api/updateItemContent", {
		method: "POST",
		body: formData
	});
}

export function getSearchItemsByName(textFragment) {
	return fetch("/api/getItemsByName", {
		method: "POST",
		body: JSON.stringify({ fragment: textFragment })
	})	
}

export function getItemContent(itemId) {
	return fetch(`/api/getItemContent/${itemId}`)
}

export function getMarkdownAsHTML(itemId) {
	return fetch(`/api/markdownAsHTML/${itemId}`)	
}
