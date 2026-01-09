# Auth0 Infrastructure

Manages Auth0 authentication with multiple applications in a single tenant (`eztrip.us.auth0.com`).

## Environment Variables Required

```bash
# Terraform Management API credentials (for infrastructure management)
export TF_VAR_auth0_client_id="your-terraform-management-client-id"
export TF_VAR_auth0_client_secret="your-terraform-management-client-secret"
```

## Resources Created

- **Web App (Dev)** - SPA with `localhost:4200` callbacks
- **Web App (Prod)** - SPA with `ez-trip.ai` callbacks
- **API (Dev)** - `https://api-dev.ez-trip.ai`
- **API (Prod)** - `https://api.ez-trip.ai`
- **API Management M2M** - For Go API user management
- **Database Connection** - Username/Password auth

## Usage

```bash
cp terraform.tfvars.example terraform.tfvars  # Configure non-sensitive values
tofu init
tofu plan
tofu apply
tofu output  # Get client IDs and audiences for app config
```

## Outputs

- `dev_web_app_client_id` - For Angular `environment.development.ts`
- `prod_web_app_client_id` - For Angular `environment.ts`
- `dev_api_identifier` - Dev API audience: `https://api-dev.ez-trip.ai`
- `prod_api_identifier` - Prod API audience: `https://api.ez-trip.ai`
- `api_management_client_id` - For Go API `.env`
