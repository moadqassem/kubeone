/*
Copyright 2019 The KubeOne Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

variable "cluster_name" {
  description = "Name of the cluster"
  type        = string
}

variable "apiserver_alternative_names" {
  description = "subject alternative names for the API Server signing cert."
  default     = []
  type        = list(string)
}

variable "worker_os" {
  description = "OS to run on worker machines"

  # valid choices are:
  # * ubuntu
  # * centos
  default = "ubuntu"
  type    = string
}

variable "ssh_public_key_file" {
  description = "SSH public key file"
  default     = "~/.ssh/id_rsa.pub"
  type        = string
}

variable "ssh_port" {
  description = "SSH port to be used to provision instances"
  default     = 22
  type        = number
}

variable "ssh_username" {
  description = "SSH user, used only in output"
  default     = "root"
  type        = string
}

variable "ssh_private_key_file" {
  description = "SSH private key file used to access instances"
  default     = ""
  type        = string
}

variable "ssh_agent_socket" {
  description = "SSH Agent socket, default to grab from $SSH_AUTH_SOCK"
  default     = "env:SSH_AUTH_SOCK"
  type        = string
}

# provider specific settings

variable "dc_name" {
  default     = "dc-1"
  description = "datacenter name"
  type        = string
}

variable "datastore_name" {
  default     = "datastore1"
  description = "datastore name"
  type        = string
}

variable "datastore_cluster_name" {
  default     = ""
  description = "datastore cluster name"
  type        = string
}

variable "resource_pool_name" {
  default     = ""
  description = "cluster resource pool name"
  type        = string
}

variable "folder_name" {
  default     = "kubeone"
  description = "folder name"
  type        = string
}

variable "network_name" {
  default     = "public"
  description = "network name"
  type        = string
}

variable "compute_cluster_name" {
  default     = "cl-1"
  description = "internal vSphere cluster name"
  type        = string
}

variable "template_name" {
  default     = "ubuntu-18.04"
  description = "template name"
  type        = string
}

variable "disk_size" {
  default     = 50
  description = "disk size"
  type        = number
}

variable "control_plane_memory" {
  default     = 2048
  description = "memory size of each control plane node in MB"
  type        = number
}

variable "worker_memory" {
  default     = 2048
  description = "memory size of each worker node in MB"
  type        = number

}

variable "worker_disk" {
  default     = 10
  description = "disk size of each worker node in GB"
  type        = number
}

variable "api_vip" {
  default     = ""
  description = "virtual IP address for Kubernetes API"
  type        = string
}

variable "vrrp_interface" {
  default     = "ens192"
  description = "network interface for API virtual IP"
  type        = string
}

variable "vrrp_router_id" {
  default     = 42
  description = "vrrp router id for API virtual IP. Must be unique in used subnet"
  type        = number
}
