resource "random_string" "cf_tunnel_secret" {
  length  = 64
  special = false
}

resource "cloudflare_zero_trust_tunnel_cloudflared" "cf_tunnel" {
  account_id = var.cf_account_id
  name       = "crtunnel"
  secret     = random_string.cf_tunnel_secret.result
}

output "cf_tunnel_secret" {
  value = random_string.cf_tunnel_secret.result
}

resource "cloudflare_record" "cf_dns_record" {
  zone_id = var.cf_zone_id
  name    = "cr"
  content = cloudflare_zero_trust_tunnel_cloudflared.cf_tunnel.cname
  type    = "CNAME"
  ttl     = 1
  proxied = true
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "cf_tunnel_cfg" {
  account_id = var.cf_account_id
  tunnel_id  = cloudflare_zero_trust_tunnel_cloudflared.cf_tunnel.id
  config {
    ingress_rule {
      hostname = cloudflare_record.cf_dns_record.hostname
      service  = "http://localhost:8080"
    }
    ingress_rule {
      service = "http_status:404"
    }
  }
}
