# Exercises for "Dangerous attack paths: Development Environment Security - Devices and CI/CD pipelines"

This repository contains the training materials of "Dangerous Attack Path: Development Environment Security - Devices and CI/CD pipelines" at Security Camp 2022.

## Course Abstract
TBD

## Disclaimer
- This repository's purpose is education and security research to cybersecurity and software engineering community.

## Presentation Slide
- I plan to publish slides (English Version and Japanese Version) within Aug

## Lab Environment
- Google Cloud
- GitHub

## Terminal
- I recommend Cloud Shell for this exercise

[![Open in Cloud Shell](http://gstatic.com/cloudssh/images/open-btn.png)](https://console.cloud.google.com/cloudshell/open?git_repo=https://github.com/rung/seccamp2022-devenv-security-training)

- You can also use your own shell envionment on your device
  - Mac: Terminal
  - Windows: WSL

## Exercises
#### Each section has Additional Exercises. You can try them too.

### [Preparation: Setup Google Cloud and GitHub](0-preparation/README.md)
- Goal: Log in to each service

### [Exercise1: What credentials your PC has](./1-exercise1/README.md)
- Goal: Understand what credentials your PC has
- Exercises: Investigate Chrome's profile(e.g. decrypt Cookie), Check GitHub's credentials, Check Google Cloud's credentials, Check SSH Key

### [Exercise 2: Try to secure your token](./2-exercise2/README.md)
- Goal: Try some mitigation method from the slide
- Exercise: Try Webauthn, Try least privilege on Google Cloud, Assign temporary role via IAM Condition on Google Cloud, Try Keyless (within Cloud)

### [Exercise3: Try continuous deployment and Infrastructure as code](./3-exercise3/README.md)
- You need to do [Lab Setup](./3-exercise3/lab-setup.md) before this Exercise
- Goal: Understand concept of Continuous Deployment and Infrastructure as code(Terraform)
- Exercise: Modify Go code and see automatic deployment, Add configuration via Terraform

### [Exercise4: CI/CD Attacks](./4-exercise4/README.md)
- Goal: Attacks CI/CD pipeline and understand the attack surface
- Exercise: Overwrite source code without any review, Steal secret from non-protected branch, Try Supply-Chain attacks via Actions the repository uses

### [Exercise5: Secure your CI/CD pipeline](./5-exercise5/README.md)
- Goal: Try to secure CI/CD pipeline from attacks
- Exercise: Configure Branch Protection, Configure OIDC, then try keyless between GitHub actions and Google Cloud
