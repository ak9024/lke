resource "linode_domain" "adiatma_tech" {
  type = "master"
  domain = "adiatma.tech"
  soa_email = "adiatma.mail@gmail.com"
}

resource "linode_domain" "malascoding_com" {
  type = "master"
  domain = "malascoding.com"
  soa_email = "adiatma.mail@gmail.com"
} 
