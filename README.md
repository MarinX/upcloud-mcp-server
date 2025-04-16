# Upcloud MCP Server

## Work in progress!!!

This is a work in progress and not yet ready for any type of use. Please check back later for updates.

## Overview

The Upcloud MCP Server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction)
server that provides seamless integration with Upcloud APIs, enabling advanced
automation and interaction capabilities for developers and tools.
The idea comes from [Github MCP server](https://github.com/github/github-mcp-server/)

## Use Cases

- Extracting and analyzing data from Upcloud services.
- Building AI powered tools and applications that interact with Upcloud's ecosystem.

## Prerequisites

1. To run the server in a container, you will need to have [Docker](https://www.docker.com/) installed.
2. Once Docker is installed, you will also need to ensure Docker is running.
3. Lastly you will need to populate UPCLOUD_USERNAME and UPCLOUD_PASSWORD variables.

## Installation

### Local usage

```sh
go build cmd/upcloud-mcp-server
```

and move the binary to your `bin` location and move the `claude_desktop_config.json` to Claude configuration directory. Don't forget to populate ENV variables!

On Mac: ~/Library/Application\ Support/Claude/claude_desktop_config.json

On Windows: AppData\Claude\claude_desktop_config.json

### Usage with Claude Desktop

{
"mcpServers": {
"github": {
"command": "docker",
"args": [
"run",
"-i",
"--rm",
"-e",
"UPCLOUD_USERNAME",
"-e",
"UPCLOUD_PASSWORD",
"ghcr.io/marinx/upcloud-mcp-server"
],
"env": {
"UPCLOUD_USERNAME": "<YOUR_USERNAME>",
"UPCLOUD_PASSWORD": "<YOUR_PASSWORD>"
}
}
}
}
![tools](./static/tools.png)
![ask](./static/ask.png)

## Tools

### Account

- **get_account** - Get account of current user
- **get_account_details** - Get account details by username
  - `username`: Account username (string, required)
- **get_account_list** - Get account list

### Database

- **get_database** - Get managed database by uuid
  - `uuid`: Database UUID (string, required)

### Server

- **get_servers** - Get servers
- **get_server_details** - Get server details
  - `uuid`: Server UUID (string, required)

### Kubernetes

- **get_kubernetes_clusters** - Get kubernetes clusters
- **get_kubernetes_cluster** - Get kubernetes cluster
  - `uuid`: Cluster UUID (string, required)
- **get_kubernetes_plans** - Get kubernetes plans

## License

This project is licensed under the terms of the MIT open source license. Please refer to [MIT](./LICENSE) for the full terms.
