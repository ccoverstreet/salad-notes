# Salad Notes

The goal for this project is to create a more robust version of the note-taking setup I've been using. Since I like doing as much work as possible in the terminal, I take markdown (Pandoc flavor) notes using vim and use Pandoc to generate a PDF file. I simultaneously have  PDF viewer open (zathura) which refreshes whenever the opened file is written to. This way, whenever I write any changes in vim, the rendered markdown file updates in the PDF viewer.

One of the shortcomings of this approach is that linking between files in different subdirectories (such as school work and personal research) requires special attention to the path in which the PDF viewer is opened. This problem could be mitigated by setting environment variables or using some form of start script when editing or creating a file.

~~This project is a web server that is started in the root of a notebook (directory with or without nested subdirectories). The client would connect to the web server and be able to view markdown files within the observed directory. Any writes to markdown files would send the updated file to any connected clients and updated the displayed markdown.~~

This project now is a web interface that uses the Ace browser editor to support VIM keys. Clients can view the webpage to create, search/filter, edit, and preview markdown files

## Usage

1. Start `salad-notes` in the directory that you want to use as the root of your notebook
	- This will create a directory that will contain all created files
2. Open a browser tab and go to `localhost:21345`
3. Click the add button in the top left
	- This will create a blank new document
	- You can change the name and add some tags
3. Start writing markdown. This project now uses gomarkdown to generate the HTML. I have plans to add support for multiple backends in the future
	- The editor uses VIM keys, but I plan on adding an option for this later
	- Ctrl-s saves the content of the file and refreshes the preview

## Implemented Features

- Client view refreshes on file write
- Can search or filter through markdown

## Planned Features

- Builtin in canvas drawing
	- Useful for quickly sketching diagrams in notes
- Support for image pasting
