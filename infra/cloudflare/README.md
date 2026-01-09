# Cloudflare Infrastructure

Manages DNS, Cloudflare Pages deployment, and domain configuration for `ez-trip.ai`.

## Environment Variables Required

```bash
# Cloudflare API Token (provider reads this directly)
export CLOUDFLARE_API_TOKEN="your-cloudflare-api-token"
```

Get token from: [Cloudflare Dashboard](https://dash.cloudflare.com/profile/api-tokens) → Create Token → "Edit zone DNS"

## Resources Created

- **DNS Zone** - Manages `ez-trip.ai` zone
- **DNS Records** - `ez-trip.ai`, `www.ez-trip.ai`, `api.ez-trip.ai`
- **Pages Project** - `eztrip-web` connected to GitHub `camden-brown/EzTrip`
- **Custom Domains** - Maps root and www to Pages deployment

## Usage

```bash
cp terraform.tfvars.example terraform.tfvars  # Configure account ID and domain
tofu init
tofu plan
tofu apply
tofu output  # Get deployment URLs
```

## Variables

**Required in `terraform.tfvars`:**
- `cloudflare_account_id` - From Cloudflare dashboard
- `domain` - `ez-trip.ai`
- `github_repo_owner` - `camden-brown`

**Optional:**
- `api_server_ip` - Leave empty until API deployed
- `node_version` - Default: `24`
