# Exercise3: Try continuous deployment and Infrastructure as code
## Goal
- Understand concept of Continuous Deployment and Infrastructure as code(Terraform)

## Preparation
- You need to setup lab environment (Google Cloud and GitHub) to proceed this exercise
  - [Lab Setup Procedure](lab-setup.md)

## Current State

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183403529-505390f7-ecfe-4cd4-9356-c7fcedfa5801.png" height="300"> </kbd>

## Exercises
- Modify Go code and see automatic deployment
- Add configuration via Terraform
- (optional) read code

## Exercises Procedure
### 1. Modify Go code and see automatic deployment
- Access Cloud Run
  - https://console.cloud.google.com/run
  - You can get URL from the console (Click the name of Cloud Run)
    - `https://training-helloworld-*****.a.run.app`
- Modify message
  - https://github.com/rung/training-devenv-security/blob/467bcce6babc26afec22e8549355aab88ea24423/template/devenv-security-app/main.go#L27

- Re-Access Cloud Run
  - Confirm an automated deployment

### 2. Add configuration via Terraform
- Document: https://registry.terraform.io/providers/hashicorp/google/latest/docs
- For instance, you can create a new service account like this configuration
```terraform
resource "google_service_account" "test" {
  account_id   = "service-account-id"
}
```

### 3. (optional) read code
- Understand what happens from the source codes
