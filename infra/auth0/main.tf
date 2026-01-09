# =============================================================================
# Auth0 Application - Web App Development (SPA)
# =============================================================================

resource "auth0_client" "web_app_dev" {
  name        = "${var.app_name} Web App - Development"
  description = "Angular SPA for ${var.app_name} development (allows localhost)"
  app_type    = "spa"

  # Development URLs - localhost only
  callbacks = [
    "http://localhost:4200"
  ]
  allowed_logout_urls = [
    "http://localhost:4200"
  ]
  web_origins = [
    "http://localhost:4200"
  ]

  # OAuth settings
  oidc_conformant = true

  jwt_configuration {
    alg = "RS256"
  }

  # Grant types for SPA (uses PKCE by default)
  grant_types = [
    "authorization_code",
    "refresh_token"
  ]
}

# =============================================================================
# Auth0 Application - Web App Production (SPA)
# =============================================================================

resource "auth0_client" "web_app_prod" {
  name        = "${var.app_name} Web App - Production"
  description = "Angular SPA for ${var.app_name} production (production URLs only)"
  app_type    = "spa"

  # Production URLs - no localhost
  callbacks = [
    "https://${var.production_domain}",
    "https://www.${var.production_domain}"
  ]
  allowed_logout_urls = [
    "https://${var.production_domain}",
    "https://www.${var.production_domain}"
  ]
  web_origins = [
    "https://${var.production_domain}",
    "https://www.${var.production_domain}"
  ]

  # OAuth settings
  oidc_conformant = true

  jwt_configuration {
    alg = "RS256"
  }

  # Grant types for SPA (uses PKCE by default)
  grant_types = [
    "authorization_code",
    "refresh_token"
  ]
}

# =============================================================================
# Auth0 API - Backend GraphQL API (Development)
# =============================================================================

resource "auth0_resource_server" "api_dev" {
  name       = "${var.app_name} API - Development"
  identifier = var.dev_api_identifier

  # Token settings
  token_lifetime                                  = 86400 # 24 hours
  signing_alg                                     = "RS256"
  allow_offline_access                            = true # Enable refresh tokens
  skip_consent_for_verifiable_first_party_clients = true
}

# =============================================================================
# Auth0 API - Backend GraphQL API (Production)
# =============================================================================

resource "auth0_resource_server" "api_prod" {
  name       = "${var.app_name} API - Production"
  identifier = var.prod_api_identifier

  # Token settings
  token_lifetime                                  = 86400 # 24 hours
  signing_alg                                     = "RS256"
  allow_offline_access                            = true # Enable refresh tokens
  skip_consent_for_verifiable_first_party_clients = true
}

# =============================================================================
# Machine-to-Machine Application - API Management Client
# =============================================================================

resource "auth0_client" "api_management" {
  name        = "${var.app_name} API - Management Client"
  description = "M2M client for Go API to manage Auth0 users (seeding, user operations)"
  app_type    = "non_interactive"

  # M2M apps use client credentials grant
  grant_types = [
    "client_credentials"
  ]

  jwt_configuration {
    alg = "RS256"
  }
}

# Grant permissions to Management API
resource "auth0_client_grant" "api_management_grant" {
  client_id = auth0_client.api_management.id
  audience  = "https://${var.auth0_domain}/api/v2/"

  scopes = [
    "read:users",
    "update:users",
    "delete:users",
    "create:users",
    "read:users_app_metadata",
    "update:users_app_metadata",
    "delete:users_app_metadata",
    "create:users_app_metadata",
    "read:user_idp_tokens"
  ]
}

# =============================================================================
# Database Connection - Username/Password
# =============================================================================
# Note: Auth0 creates "Username-Password-Authentication" by default
# We reference the existing connection and just enable it for our apps

data "auth0_connection" "database" {
  name = "Username-Password-Authentication"
}

resource "auth0_connection_clients" "database_clients" {
  connection_id = data.auth0_connection.database.id

  enabled_clients = [
    auth0_client.web_app_dev.id,
    auth0_client.web_app_prod.id
  ]
}
