# Exercise1: What credentials your PC has
## Goal
- Understand what credentials your PC has

## Exercise
- Investigate Chrome's profile
- Check GitHub's credentials
- Check Google Cloud's credentials
- Check SSH Key

## Additional Exercise
- Investigate other tokens
  - Password Manager(like 1Password, Lastpass)
  - Local Application(Slack, etc)
  - Other public cloud services

## Exercise's procedure
### 1. Investigate Chrome's profile
#### Profile Path
- Windows (Please use Windows side, not WSL)
  - `%HOMEPATH%\AppData\Local\Google\Chrome\User Data\<Your Profile Name>`

- Mac
  - `~/Library/Application\ Support/Google/Chrome/<Your Profile Name>`

#### What you can find
- Search `Cookies`(might be under /Network dir), `Login Data`, `History`, `Local Storage`
- Use `file` command to check the file format
- Use `strings` command to `Local Storage`. You can confirm that it's not encrypted

#### See your Cookie and Password via this tool

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183426308-58e1915d-c6bf-4057-abe6-5027fa2f31d0.png" height="300"> </kbd>

- https://github.com/rung/HackChromeData
  - docker is required for `make build`
  - Usage is [here](https://github.com/rung/HackChromeData#usage).
```
# build
gh repo clone rung/HackChromeData
cd HackChromeData
make build
```

- You can get Cookie and Password
  - without any additional authentication (Windows)
  - using user password (Mac)

### 2. Check GitHub's credentials
- You can get GitHub Cookie from Browser
- You can access ` ~/.config/gh/hosts.yml`, then you can get oauth_token, which is created via `gh` command during [preparation](https://github.com/rung/seccamp2022-device-cicd/tree/main/0-preparation)

### 3. Check Google Cloud's credentials
- You can get Google Cloud Cookie from Browser
- You can access the json file of `Service Account Key` you downloaded during [preparation](https://github.com/rung/training-devenv-security/blob/main/0-preparation/README.md)
- You can access to `~/.config/gcloud/legacy_credentials/<your emailaddress>/adc.json`
  - It's legacy token, but still available
- You can access `credential.db` too
```
$ file ~/.config/gcloud/credentials.db
**/.config/gcloud/credentials.db: SQLite 3.x database, last written using SQLite version 3036000

$ sqlite3 ~/.config/gcloud/credentials.db "select value from credentials;"
```
- (Reference) https://book.hacktricks.xyz/cloud-security/gcp-security/gcp-persistance
- (Reference) https://about.gitlab.com/blog/2020/02/12/plundering-gcp-escalating-privileges-in-google-cloud-platform/#steal-gcloud-authorizations

### 4. Check SSH Key
```
ls -l ~/.ssh/
```
