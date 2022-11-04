# Salad Notes

The goal for this project is to create a more robust version of the note-taking setup I've been using. Since I like doing as much work as possible in the terminal, I take markdown (Pandoc flavor) notes using vim and use Pandoc to generate a PDF file. I simultaneously have  PDF viewer open (zathura) which refreshes whenever the opened file is written to. This way, whenever I write any changes in vim, the rendered markdown file updates in the PDF viewer.

One of the shortcomings of this approach is that linking between files in different subdirectories (such as school work and personal research) requires special attention to the path in which the PDF viewer is opened. This problem could be mitigated by setting environment variables or using some form of start script when editing or creating a file.

This project is a web server that is started in the root of a notebook (directory with or without nested subdirectories). The client would connect to the web server and be able to view markdown files within the observed directory. Any writes to markdown files would send the updated file to any connected clients and updated the displayed markdown.

## Usage

1. Start `salad-notes` in the directory that you want to use as the root of your notebook
2. Open a browser tab and go to `localhost:33322`
3. Start writing to a `.md` file. The client will automatically view the file once you write to it

### Usage with Vim/Neovim

- Open the root directory of the notebook in vim and navigate using netrw
- Linking to other files can be done by filename completion

## Implemented Features

- Client view refreshes on file write
- Simple file explorer for viewing files within the notebook
- Clicking on the current filename copies the filename text to clipboard
	- Useful for quickly jumping to files in Vim

## Planned Features

- Multiple viewing panes
	- Once file linking is sorted, a primary and secondary view pane would allow for users to have multiple notes open. 
