# Exercise 2: Try to secure your token
## Goal
- Try some mitigation method from the slide

## Exercise
- Try Webauthn
- Assign temporary role via IAM Condition on Google Cloud
- (Optional) Try Keyless (within Cloud)
- (Optional) Try least privilege on Google Cloud

## Additional Exercise
- Network Restriction to Google Cloud API via [VPC Service Controls](https://cloud.google.com/vpc-service-controls)
  - Caution: Org-level permission is required

## Procedure
### 1. Try WebAuthn on your device
- Caution: GitHub supports WebAuthn, so that you can try WebAuthn
- Register your key
  - https://github.com/settings/two_factor_authentication/configure

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183346202-dee50d9a-797c-4c28-a102-3bdfa65e848f.png" height="150"> </kbd>

- Log in via WebAuthn (Windows Hello, Touch/Face ID)
  

### 2. Assign temporary role via IAM Condition on Google Cloud
- Service Account: `training-sa` (You created it in [0.preparation](../0-preparation/README.md#create-a-service-account))
- IAM configuration: https://console.cloud.google.com/iam-admin/iam

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183404485-1e9c43ac-ff54-48be-b54e-abf9cba60694.png" height="150"> </kbd>

- You use the SA later, so please don't set a close expiration date

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183404663-28aec4e4-8191-4326-b305-3ce969facf31.png" height="150"> </kbd>

### 3. (Optional) Try Keyless (within Cloud)
- You can make new GCE instance from [here](https://console.cloud.google.com/compute/instances)
  - Machine Type: e2-micro
  - Service Account: any SA is fine
- SSH to the instance (`Open in browser window`)

<kbd> <img src="https://user-images.githubusercontent.com/1150301/183394377-3294ac7e-3032-4070-a99d-474c91ec23db.png" height="300"> </kbd>

- Confirm Metadata (Ref: https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances#applications)
  - You can get the SA's temporary token from metadata endpoint
  - (SSRF attack steals this token)
```bash
# Confirm Token (expire: 3600sec)
$ curl "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/token" -H "Metadata-Flavor: Google"

# Get the token info
$ curl "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=ya29.c.b0AXv0zTODIOa_oxONq..."
{
  "issued_to": "*",
  "audience": "*",
  "scope": "*",
  "expires_in": 3245,
  "access_type": "online"
}
```

- Google Compute Engine Instance has Identity on Google Cloud, so without static key, it can get temporary token through metadata service

### 4. (Optional) Try least privilege on Google Cloud
- The purpose is to understand [IAM](https://cloud.google.com/iam/docs/understanding-roles)
- Go to your project on Google Cloud
  - https://console.cloud.google.com/iam-admin/serviceaccounts
    - You can try to make `custom role` on https://console.cloud.google.com/iam-admin/roles
