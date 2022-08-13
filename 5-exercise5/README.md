# Exercise5: Secure your CI/CD pipeline
## Goal
- Try to secure CI/CD pipeline from attacks and understand the limitation

## Exercises
- Configure Branch Protection
  - Caution: GitHub Free version doesn't support it. You need to use GitHub Pro or Team
- Configure OIDC, then try keyless between GitHub actions and Google Cloud

## Additional Exercises
- Enter GitHub Actions using Tailscale and reverseshell (The procedure is in the slide)
- Steal Google Service Account Token in the keyless environment
- Steal IaC's CI(read) token, then see tfstate file on Google Storage
  - Then Consider what role CI should have to do `Least Privilege` policy

## Exercises procedure
### 1. Configure Branch Protection
- Target Repository: `devenv-security-iac`

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183426687-624dad37-10ba-44ea-99f3-b82eed42ad2c.png" height="400"> </kbd>


### 2. Configure OIDC, then try keyless between GitHub actions and Google Cloud
- Target Repository: `devenv-security-iac`
- Rename `google_actions_oidc.tf_` to `google_actions_oidc.tf` (`devenv-security-iac/terraform/training-project/`)
  - It enables Workload Federation

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183426987-2ba5d9ce-2d9d-4e33-882e-e0e732f3568c.png" height="200"> </kbd>


- Modify Iac's Actions Workflows
  - `devenv-security-iac/.github/workflows/apply.yaml`, `devenv-security-iac/.github/workflows/plan.yaml`
    - Uncomment `id-token: 'write'`
    - Comment out `credentials_json`
    - Uncomment `workload_identity_provider: 'projects/<Project Number>/locations/global/workloadIdentityPools/training-pool/providers/training-provider'`
    - Uncomment `service_account: 'iac-actions-cd@<Project ID>.iam.gserviceaccount.com'`
- You can do the same thing to App too.
