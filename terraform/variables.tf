variable "token" {
    description = "Linode API token"
}

variable "k8s_version" {
    description = "Kubernetes version"
    default = "1.26"
}

variable "label" {
    description = "Cluster's label"
    default = "python-test"
}

variable "region" {
    description = "Cluster resource region"
    default = "ap-southeast"
}

variable "tags" {
    description = "Tag your cluster"
    type = list(string)
    default = ["play"]
}

variable "pools" {
    description = "Node pool specifications"
    type = list(object({
        type = string
        count = number
    }))
    default = [
        {
            type = "g6-standard-1"
            count = 1
        }
    ]
}