![GitHub License](https://img.shields.io/github/license/antfie/obsidian-tools)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/antfie/obsidian-tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/antfie/obsidian-tools)](https://goreportcard.com/report/github.com/antfie/obsidian-tools)
![GitHub Release](https://img.shields.io/github/v/release/antfie/obsidian-tools)
![GitHub Downloads (all assets, latest release)](https://img.shields.io/github/downloads/antfie/obsidian-tools/total)
![Docker Image Size](https://img.shields.io/docker/image-size/antfie/obsidian-tools/latest)
![Docker Pulls](https://img.shields.io/docker/pulls/antfie/obsidian-tools)

# obsidian-tools

This is an Obsidian toolkit. It has the following features:

| Tool                       | Description                                                                     |
|----------------------------|---------------------------------------------------------------------------------|
| `move`                     | Move a file from one Obsidian vault to another, removing any unreferenced files |
| `copy`                     | Copy a file from one Obsidian vault to another                                  |
| `delete`                   | Delete a note and any unreferenced attachments                                  |
| `find_missing_attachments` | Identify missing attachments                                                    |
| `find_duplicates`          | Identify duplicate files                                                        |
| `find_empty_files`         | Identify empty files                                                            |
| `find_sync_conflicts`      | Identify Syncthing conflicts                                                    |

## How Do I Run It?

You can run this wherever you like. Just download the appropriate binary from [here](https://github.com/antfie/obsidian-tools/releases/latest).

### Using Docker

```bash
docker pull antfie/obsidian-tools
docker run --rm -it -v "$(pwd):/vault" antfie/obsidian-tools find_empty_files /vault
```

## How Can I Support This?

W welcome fixes, features and donations.