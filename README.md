# Salad Notes

The goal for this project is to create a more robust version of the note-taking setup I've been using. Since I like doing as much work as possible in the terminal, I take markdown (Pandoc flavor) notes using vim and use Pandoc to generate a PDF file. I simultaneously have  PDF viewer open (zathura) which refreshes whenever the opened file is written to. This way, whenever I write any changes in vim, the rendered markdown file updates in the PDF viewer.

One of the shortcomings of this approach is that linking between files in different subdirectories (such as school work and personal research) requires special attention to the path in which the PDF viewer is opened. This problem could be mitigated by setting environment variables or using some form of start script when editing or creating a file.

This project is a web server that is started in the root of a notebook (directory with or without nested subdirectories). The client would connect to the web server and be able to view markdown files within the observed directory. Any writes to markdown filess would send the updated file to any connected clients and updated the displayed markdown.


