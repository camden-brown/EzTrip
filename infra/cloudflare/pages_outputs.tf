# =============================================================================
# Cloudflare Pages Outputs
# =============================================================================

output "pages_project_name" {
  description = "Name of the Cloudflare Pages project"
  value       = cloudflare_pages_project.web_app.name
}

output "pages_subdomain" {
  description = "Default Cloudflare Pages subdomain"
  value       = cloudflare_pages_project.web_app.subdomain
}

output "pages_domains" {
  description = "Custom domains configured for the Pages project"
  value = [
    cloudflare_pages_domain.root.name,
    cloudflare_pages_domain.www.name
  ]
}

output "pages_deployment_url" {
  description = "Production deployment URL"
  value       = "https://${cloudflare_pages_project.web_app.subdomain}.pages.dev"
}

output "web_analytics_tag" {
  description = "Cloudflare Web Analytics site tag"
  value       = var.enable_web_analytics ? cloudflare_web_analytics_site.web_app[0].site_tag : null
  sensitive   = true
}
