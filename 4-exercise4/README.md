# Exercise4: CI/CD Attacks
## Goal
- Attack CI/CD pipeline and understand the attack surface
  - Assume a developer's device is compromised
  - Assume tool/software dependencies are compromised

## Exercises
- Overwrite source code without any review
- Steal secrets from a non-protected branch
- Try Supply-Chain attacks via Actions the repository uses

## Additional Exercises
- Steal IaC's CI(read) token, then see tfstate file on Google Storage
- Steal App's deployment token. Even if it's editor, it can be escalated to Owner. How?
  - Reference: https://rhinosecuritylabs.com/gcp/privilege-escalation-google-cloud-platform-part-1/

## Exercises Procedure
### 1. Overwrite source code without any review
- Push source code to the main branch directly

### 2. Steal secret from non-protected branch
- `echo ${{ secrets.GCP_SA_KEY }} | base64` on CI  

### 3. Try Supply-Chain attacks via Actions the repository uses
- It's not real supply-chain attacks. but we assume our app's CD depends on "supplychain" actions.
  - https://github.com/rung/training-devenv-security/blob/main/template/devenv-security-app/supplychain/action.yml#L8
  - You can modify this yml, then it affects CD on App
