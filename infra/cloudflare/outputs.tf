output "zone_id" {
  description = "Cloudflare zone ID"
  value       = data.cloudflare_zone.main.zone_id
}

output "name_servers" {
  description = "Cloudflare nameservers for your domain"
  value       = data.cloudflare_zone.main.name_servers
}

output "domain" {
  description = "The managed domain"
  value       = var.domain
}

output "frontend_url" {
  description = "Frontend URL"
  value       = "https://${var.domain}"
}

output "api_url" {
  description = "API URL"
  value       = var.api_server_ip != "" ? "https://api.${var.domain}" : "Not configured"
}
