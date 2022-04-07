# ADO


## Usage

Call link.go with the two commit urls as command line arguments, the older commit url goes first

example:
```
go run link.go https://dev.azure.com/yourOrg/yourProject/_git/repoName/commit/25dc190b6fbe683976578abdfe7936f22607e/refName=refs/heads/main https://dev.azure.com/yourOrg/yourProject/_git/repoName/commit/1932847a190b6fbe683976578abdfe7936f22607e/refName=refs/heads/main

go run link.go OLD_COMMIT_URL NEW_COMMIT_URL

```