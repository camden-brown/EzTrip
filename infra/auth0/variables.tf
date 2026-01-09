variable "auth0_domain" {
  description = "Auth0 domain (e.g., eztrip.us.auth0.com)"
  type        = string
}

variable "auth0_client_id" {
  description = "Auth0 Management API client ID"
  type        = string
  sensitive   = true
}

variable "auth0_client_secret" {
  description = "Auth0 Management API client secret"
  type        = string
  sensitive   = true
}

variable "app_name" {
  description = "Application name"
  type        = string
  default     = "EzTrip"
}

variable "production_domain" {
  description = "Production domain"
  type        = string
  default     = "ez-trip.ai"
}

variable "dev_api_identifier" {
  description = "Auth0 API identifier for development"
  type        = string
  default     = "https://api-dev.ez-trip.ai"
}

variable "prod_api_identifier" {
  description = "Auth0 API identifier for production"
  type        = string
  default     = "https://api.ez-trip.ai"
}
