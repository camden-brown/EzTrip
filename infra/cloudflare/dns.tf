# =============================================================================
# Cloudflare Zone (Domain)
# =============================================================================

# Reference existing zone (domain purchased on Cloudflare)
data "cloudflare_zone" "main" {
  filter = {
    name = var.domain
  }
}

# =============================================================================
# DNS Records
# =============================================================================

# Root domain - points to Cloudflare Pages
resource "cloudflare_dns_record" "root" {
  zone_id = data.cloudflare_zone.main.zone_id
  name    = "@"
  content = cloudflare_pages_project.web_app.subdomain
  type    = "CNAME"
  ttl     = 1  # Automatic
  proxied = true
  comment = "Root domain - Cloudflare Pages"
}

# WWW subdomain - points to Cloudflare Pages
resource "cloudflare_dns_record" "www" {
  zone_id = data.cloudflare_zone.main.zone_id
  name    = "www"
  content = cloudflare_pages_project.web_app.subdomain
  type    = "CNAME"
  ttl     = 1  # Automatic
  proxied = true
  comment = "WWW subdomain - Cloudflare Pages"
}

# API subdomain (for future backend deployment)
resource "cloudflare_dns_record" "api" {
  count   = var.api_server_ip != "" ? 1 : 0
  zone_id = data.cloudflare_zone.main.zone_id
  name    = "api"
  content = var.api_server_ip
  type    = "A"
  ttl     = 1 # 1 = automatic
  proxied = var.enable_proxy
  comment = "API server"
}
# =============================================================================
