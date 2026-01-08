variable "cloudflare_api_token" {
  description = "Cloudflare API token with DNS edit permissions"
  type        = string
  sensitive   = true
}

variable "cloudflare_account_id" {
  description = "Cloudflare account ID"
  type        = string
}

variable "domain" {
  description = "Domain name to manage (e.g., tripscout.app)"
  type        = string
}

variable "api_server_ip" {
  description = "IP address of the API server (Cloud Run or VM)"
  type        = string
  default     = ""
}

variable "enable_proxy" {
  description = "Enable Cloudflare proxy (orange cloud) for DDoS protection and caching"
  type        = bool
  default     = true
}

# =============================================================================
# Cloudflare Pages Configuration
# =============================================================================

variable "pages_project_name" {
  description = "Name of the Cloudflare Pages project"
  type        = string
  default     = "eztrip-web"
}

variable "pages_production_branch" {
  description = "Git branch to use for production deployments"
  type        = string
  default     = "main"
}

variable "github_repo_owner" {
  description = "GitHub repository owner (username or organization)"
  type        = string
}

variable "github_repo_name" {
  description = "GitHub repository name"
  type        = string
  default     = "eztrip"
}

variable "enable_web_analytics" {
  description = "Enable Cloudflare Web Analytics for the Pages site"
  type        = bool
  default     = true
}

variable "node_version" {
  description = "Node.js version to use for builds (should match .nvmrc)"
  type        = string
  default     = "24"
}
