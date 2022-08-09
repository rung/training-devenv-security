# Preparation: Setup Google Cloud and GitHub
- You need to register Google Cloud and GitHub to do this training
- For next section, you login to each service in this preparation

## Procedure
### GitHub
#### Register SSH Key
- Register your SSH Key
  - if you don't have it, you can make it via `ssh-keygen -t ed25519`

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183255148-879c34fd-dd9c-4e75-9ef5-e9d706670b51.png" height="300"> </kbd>

#### gh command
- Install (CloudShell has `gh` by default)
  - https://github.com/cli/cli#installation
- Execute
```
$ gh auth login
? What account do you want to log into? GitHub.com
? What is your preferred protocol for Git operations? HTTPS
? How would you like to authenticate GitHub CLI? Login with a web browser

! First copy your one-time code: ****-****
Press Enter to open github.com in your browser...
```

### Google Cloud
#### Login from Web Browser
- https://console.cloud.google.com/
  - Go to your project, then please access to IAM page under IAM&Admin
  - You can change access control on the IAM page manually
<kbd> <img src="https://user-images.githubusercontent.com/1150301/183255698-635ee497-d216-4e36-a677-ead861a12db7.png" height="200"> </kbd>

#### Create a Service Account
- Go to "Service Accounts" page under IAM&Admin
https://console.cloud.google.com/iam-admin/serviceaccounts
  - Click "CREATE SERVICE ACCOUNT"

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183255757-93765552-bb01-40ca-bd41-06a546e61b41.png" height="50"> </kbd>

- Input the service account name
  - Name: training-sa

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183255874-a1ca05c9-b832-4bb0-9d62-aa648b14720c.png" height="400"> </kbd>


- Grant [Owner](https://cloud.google.com/iam/docs/understanding-roles#basic-definitions) to the service account

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183537330-0ebbcc99-e97b-4eaf-8dc4-bb196ba2bddf.png" height="300"> </kbd>

- Click the service account you created

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183256014-cd797ee5-bd37-486c-be42-6f2c2553e989.png" height="40"> </kbd>

- Create New Key

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183256041-1bd3adc2-0679-466b-a694-df90f22b966d.png" height="300"> </kbd>

- Choose JSON type, then download the json

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183256053-79741e15-73b1-4ac9-b6e1-d2df09af480a.png" height="300"> </kbd>


- Confirm the json file you downloaded

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183256170-73c32efe-290b-4abd-9165-d7077a8a0d35.png" height="300"> </kbd>


#### Login via gcloud command
- Install gcloud command (CloudShell has `gcloud` by default)
  - https://cloud.google.com/sdk/docs/install
- login
```
$ gcloud auth login
...
You are now logged in as [******@*******].
```
