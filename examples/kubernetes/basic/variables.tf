variable "prefix" {
  description = "A prefix used for all resources in this example"
  default = "acctestkt"
}

variable "location" {
  description = "The Azure Region in which all resources in this example should be provisioned"
  default = "west europe"
}

variable "kubernetes_client_id" {
  description = "The Client ID for the Service Principal to use for this Managed Kubernetes Cluster"
  default = "480276d8-7cfc-4c0c-9655-ee8001e7eaf8"
}

variable "kubernetes_client_secret" {
  description = "The Client Secret for the Service Principal to use for this Managed Kubernetes Cluster"
  default = "oq3LmilqTu3jnN8ljNdfy9GJp2dPU9ITNnHThh9RrYc="
}
