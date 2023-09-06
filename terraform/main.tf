terraform {
  required_providers {
    linode = {
        source = "linode/linode"
        version = "2.5.2"
    }
  }
}

provider "linode" {
    # token = var.linode_token
}

resource "linode_lke_cluster" "play" {
    k8s_version = var.k8s_version
    label = var.label
    region = var.region
    tags = var.tags

    dynamic "pool" {
        for_each = var.pools
        content {
            type = pool.value["type"]
            count = pool.value["count"]
        }
    }
}

output "kubeconfig" {
    value = linode_lke_cluster.play.kubeconfig
    sensitive = true
}

output "api_endpoints" {
    value = linode_lke_cluster.play.api_endpoints
}

output "status" {
    value = linode_lke_cluster.play.status
}

output "id" {
    value = linode_lke_cluster.play.id
}

output "pool" {
    value = linode_lke_cluster.play.pool
}