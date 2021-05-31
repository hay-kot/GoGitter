# GoGitter

A tiny utility to quickly and easily grab all your repos from github and store them on your machine. 

## Features
- Configuration with Yaml
- Sort by Language
- Pull Repo if the already exists on disk
- Alias/Sort repos based of search criteria


## Configuration

```yaml
# GoGitter.yaml
forks: false # Clone/Pull Forked Repos
pull: true # Pull Existing Repos
zipAll: false
sortByLanguage: true # Sort Repos into Subfolder
source:
  github:
    token: your-secret-api-key
    user: your-github-username
destination:
  local: ./Path-To-Someplace
  zip: /some/abs/path/for/zip
keywords:
  cookiecutter: CookieCutter # If cookiecutter is found in the repo name stick it in a CookieCutter folder
```

## Usage

```
$ gogitter GoGitter.yaml
```