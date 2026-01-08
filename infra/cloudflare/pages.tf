# =============================================================================
# Cloudflare Pages Project
# =============================================================================

resource "cloudflare_pages_project" "web_app" {
  account_id        = var.cloudflare_account_id
  name              = var.pages_project_name
  production_branch = var.pages_production_branch

  # Build configuration
  build_config = {
    build_command       = "npx nx run web-app:build:production"
    destination_dir     = "dist/apps/web-app/browser"
    root_dir            = ""
    web_analytics_tag   = var.enable_web_analytics ? cloudflare_web_analytics_site.web_app[0].site_tag : null
    web_analytics_token = var.enable_web_analytics ? cloudflare_web_analytics_site.web_app[0].site_token : null
  }

  # Source configuration (GitHub)
  source = {
    type = "github"
    config = {
      owner                         = var.github_repo_owner
      repo_name                     = var.github_repo_name
      production_branch             = var.pages_production_branch
      pr_comments_enabled           = true
      deployments_enabled           = true
      production_deployment_enabled = true
      preview_deployment_setting    = "all"
    }
  }

  # Deployment configurations
  deployment_configs = {
    production = {
      environment_variables = {
        NODE_VERSION = var.node_version
      }
    }

    preview = {
      environment_variables = {
        NODE_VERSION = var.node_version
      }
    }
  }
}

# =============================================================================
# Custom Domain Configuration
# =============================================================================

resource "cloudflare_pages_domain" "root" {
  account_id   = var.cloudflare_account_id
  project_name = cloudflare_pages_project.web_app.name
  name         = var.domain
}

resource "cloudflare_pages_domain" "www" {
  account_id   = var.cloudflare_account_id
  project_name = cloudflare_pages_project.web_app.name
  name         = "www.${var.domain}"
}

# =============================================================================
# Web Analytics (Optional)
# =============================================================================

resource "cloudflare_web_analytics_site" "web_app" {
  count        = var.enable_web_analytics ? 1 : 0
  account_id   = var.cloudflare_account_id
  zone_tag     = data.cloudflare_zone.main.zone_id
  auto_install = true
}
