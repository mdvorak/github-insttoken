# github-insttoken

Go binary to get a GitHub App ephemeral Installation token for a repository to allow
git `clone` / `pull` etc operations over https.

This is to simplify the steps detailed in [authenticating-with-github-apps](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/)

## Usage:
```bash
./github-insttoken \
  --private-key-file YOURGITHUBAPP.2018-10-26.private-key.pem \
  --app-id APP_ID \
  --repo Organisation/project \
  --base-url https://github.example.com/api/v3
```
Returns:
```
v1.2a04[...snip...]5172
```

This can now be used in a git clone:
```bash
git clone https://x-access-token:v1.2a04[...snip...]5172@github.com/Organisation/project.git
```
