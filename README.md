
<h1 align="center">
  <br>
  <p align="center"> <img src="https://upload.wikimedia.org/wikipedia/commons/2/2e/Pocket_App_Logo.png" alt="Pocket Logo" width="350"/> &nbsp;&nbsp;&nbsp; <img src="https://upload.wikimedia.org/wikipedia/commons/e/e9/Notion-logo.svg" alt="Notion Logo" width="100"/> </p>
  Pocket to Notion Sync Tool
  <br>
</h1>

<h4 align="center">A minimal CLI tool to synchronize all your saved articles in Pocket with your Notion database.</h4>

This tool allows you to easily import articles you've saved on Pocket into your Notion workspace. It helps organize and manage your saved content with full control over templates and rich metadata, such as article titles, URLs, and tags. No more manual copying and pasting—let the tool do the work for you!

For more detailed background and specific goals of the project, feel free to check out my [article](https://secured-polygon-902513.framer.app/projects/pocket-to-notion-sync)

![screenshot-sync](https://github.com/user-attachments/assets/07e028d9-d344-490f-bdfd-234e970e7153)

## Key Features

* Sync Pocket articles to a Notion database. 
* Saves article metadata (title, URL, tags). 

* Simple CLI tool.

## How To Use

Prerequisites:
- [Go](https://golang.org/) installed. 
- Pocket API and Notion API tokens.
```bash 
git clone https://github.com/BahaBoualii/pocket-to-notion-sync.git 
cd pocket-to-notion-sync 
go mod tidy 
go build -o pocket-to-notion-sync
```
- Create a new date property in your Notion database called `Added`.

> **Note**  
> Set environment variables:  
    
>  ```bash    
>   export POCKET_CONSUMER_KEY=your_pocket_api_token
>   export NOTION_KEY=your_notion_api_token
 
> Or create a `.env` file with the above variables.

Then just run this command:  
`./pocket-to-notion-sync`

Or run it with an optional `notion-db` flag:  
`./pocket-to-notion-sync --notion-db=your_notion_database_id`

## Roadmap

 - [ ] Add automatic sync scheduling.
 - [ ] Filter Pocket articles by tags/date.
 - [ ] Enhance logging and error handling.
 - [ ] Build more customization options for the Notion page properties.

## Contributions

Feel free to contribute by opening issues or submitting PRs.

