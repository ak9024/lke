resource "linode_lke_cluster" "lke-cluster-sg" {
    label       = "lke-cluster-sg"
    k8s_version = "1.28"
    region      = "ap-south"

    pool {
        type  = "g6-standard-2"
        count = 1
    }
}
