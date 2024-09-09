# New Design

- Markdown files 
    - Likely using GitHub markdown format
    - Execute code cells sort of like Jupyter
        - Code within backticks will be executed in the directory
        - Useful for generating plots or figures using Python
            - Plots should be saved as an image and then displayed in Markdown using the normal image syntax
        - For now, all cells will be isolated, no preserved state for markdown file
            - Likely to change in the future to improve efficiency
- Organized on the filesystem in a normal directory structure
- Can be rendered to static HTML to be served on a blog or webpage
- YAML header for adding tags and other metadata
