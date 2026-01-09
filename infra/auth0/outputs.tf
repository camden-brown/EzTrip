output "dev_web_app_client_id" {
  description = "Auth0 Development Web App Client ID (use in environment.development.ts)"
  value       = auth0_client.web_app_dev.client_id
}

output "prod_web_app_client_id" {
  description = "Auth0 Production Web App Client ID (use in environment.ts)"
  value       = auth0_client.web_app_prod.client_id
}

output "auth0_domain" {
  description = "Auth0 domain"
  value       = var.auth0_domain
}

output "dev_api_identifier" {
  description = "Auth0 Development API identifier (audience)"
  value       = auth0_resource_server.api_dev.identifier
}

output "prod_api_identifier" {
  description = "Auth0 Production API identifier (audience)"
  value       = auth0_resource_server.api_prod.identifier
}

output "database_connection_name" {
  description = "Auth0 database connection name"
  value       = data.auth0_connection.database.name
}

output "api_management_client_id" {
  description = "Auth0 Management API Client ID (for Go API to seed users)"
  value       = auth0_client.api_management.client_id
}
