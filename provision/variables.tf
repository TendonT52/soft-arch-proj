variable "proxmox_username" {
  description = "Proxmox username"
  type        = string
  sensitive   = true
}

variable "proxmox_password" {
  description = "Proxmox password"
  type        = string
  sensitive   = true
}

variable "proxmox_api_url" {
  description = "Proxmox API URL"
  type        = string
  default     = "https://192.168.0.200:8006/api2/json"
}

variable "ssh_key" {
  description = "SSH public key"
  type        = string
  sensitive = true
}

variable "template_name" {
  default = "ubuntu-2023-cloudinit-template"
}

variable "vm_count" {
  default = 1
}