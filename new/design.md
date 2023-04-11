# Design

## Information stored

- Notes
- Images (png, svg, eps, etc)
- HTML canvas states/editable images
	- This could just be saved as a png

## Abstraction

- The UID is constructed as follows
	- <file extension>-<32 Base64 characters>

### Inserting Items

- Required information is item content, tags, name (does not have to be unique), type (will be an enum)
	- Type enums
		- Markdown
		- Images
- For images, the content will be the file identifier for the corresponding image stored on the file system

### Updating Items

- Item specified by UID

### Removing Items

- Item specified by UID

### Query Items

- Items can be retrieved quickly using a global UID
- Items can be searched by tag
- Items can be searched by title
- Future optimizations
	- Caching titles/tags and the corresponding UIDs
		- Maybe making a secondary table



```json
{
	"UID": "SOMERANDOMUNIQUESTRING",
	"name": "maybe/uri/structure",
	"content": "Markdown, base64",
	"tags": [
		"array",
		"of",
		"tags"
	]
}
```
