# What is this

This repository is Hands-on Exercises for "[Dangerous Attack Path: Modern Development Environment Security - Devices and CI/CD pipelines](https://speakerdeck.com/rung/training-devenv-security-en)" at Security Camp (Japan) 2022.  
This training is created by Hiroki SUEZAWA([@rung](https://twitter.com/rung/)), 
Author of [Common Threat Matrix for CI/CD Pipeline
](https://github.com/rung/threat-matrix-cicd) and reviewer of [Top10 CI/CD Security Risks](https://www.cidersecurity.io/top-10-cicd-security-risks/).

(Japanese: このレポジトリはセキュリティ・キャンプ全国大会2022のトレーニング「[開発環境のセキュリティおよびCI/CDパイプラインのセキュア化](https://speakerdeck.com/rung/training-devenv-security-ja)」の演習です)

## ℹ️ Course Abstract

Over the past ten years, the development environment in which software is being developed has changed dramatically: with the spread of DevOps culture and the increased use of Cloud infrastructures, and applications are now deployed through CI/CD pipelines. In addition, development is now conducted not only in the office, but also outside the company.

In this training, we will discuss how to attack and secure modern production environments, mainly from the perspective of client-side attacks using malware and supply-chain attacks, and explain comprehensive attack methods and measures, followed by hands-on exercises.

In hands-on exercises, You can decrypt your browser's cookie and password, and other credentials. Then you create a new CI/CD pipeline for automated deployment and Infrastructure as Code, attacking and securing them on your hand!

## 🛑 Disclaimer
- This repository's purpose is education and security research to cybersecurity and software engineering community.
- We use GitHub as Source Code Management and Google Cloud as a public cloud in this exercise, but the contents of the slide can apply to others.

## 📖 Presentation Slide
|English|Japanese (日本語)|
|---|---|
|[![English](https://user-images.githubusercontent.com/1150301/184302502-b0c6138e-6498-4fc8-ab2d-29ea69d5e05f.png)](https://speakerdeck.com/rung/training-devenv-security-en)|[![Japanese](https://user-images.githubusercontent.com/1150301/184300825-f351f68d-7704-43de-a193-9b7a98b297cb.png)](https://speakerdeck.com/rung/training-devenv-security-ja)|


## 🎫 Requirements for this exercise
- Google Cloud Account
  - (We use IAM, small Storage, small GCE instance and Cloud Run only, so free credits are enough)
- GitHub Account
  - Free plan users can do most of the exercises
  - But we try `branch protection,` so Pro plan is recommended

## 💻 Terminal
- You can use your Unix-like terminal on your device
  - Mac: Terminal
  - Windows: WSL

- You can also use Cloud Shell for this exercise

[![Open in Cloud Shell](http://gstatic.com/cloudssh/images/open-btn.png)](https://console.cloud.google.com/cloudshell/open?git_repo=https://github.com/rung/seccamp2022-devenv-security-training)

## ✍️ Exercises
#### Each section has Additional Exercises. You can try them too.

### [Preparation: Setup Google Cloud and GitHub](0-preparation/README.md)
- Goal: Log in to each service

### [Exercise1: What credentials your PC has](./1-exercise1/README.md)
- Goal: Understand what credentials your PC has
- Exercises: Investigate Chrome's profile(e.g. decrypt Cookie), Check GitHub's credentials, Check Google Cloud's credentials, Check SSH Key

### [Exercise 2: Try to secure your token](./2-exercise2/README.md)
- Goal: Try some mitigation methods from the slide
- Exercise: Try Webauthn, Try Keyless (within Cloud), Assign temporary role via IAM Condition on Google Cloud, Try Least Privilege on Google Cloud, 

### [Exercise3: Make and Try continuous deployment and Infrastructure as code](./3-exercise3/README.md)
- You need to do [Lab Setup](./3-exercise3/lab-setup.md) before this Exercise
- Goal: Understand the concept of Continuous Deployment and Infrastructure as code(Terraform)
- Exercise: Modify Go code and see automatic deployment, Add configuration via Terraform

### [Exercise4: Attack against CI/CD](./4-exercise4/README.md)
- Goal: Attack on CI/CD pipelines and understanding the attack surface
- Exercise: Overwrite source code without any review, Steal secrets from a non-protected branch, Try Supply-Chain attacks via Actions the repository uses

### [Exercise5: Secure your CI/CD pipeline](./5-exercise5/README.md)
- Goal: Try to secure CI/CD pipeline from attacks
- Exercise: Configure Branch Protection, Configure OIDC, then try keyless between GitHub actions and Google Cloud
