terraform {
  required_providers {
    proxmox = {
      source  = "Telmate/proxmox"
      version = "2.9.14"
    }
  }
}

provider "proxmox" {
  pm_user         = var.proxmox_username
  pm_password     = var.proxmox_password
  pm_api_url      = var.proxmox_api_url
  pm_tls_insecure = true
}

resource "proxmox_vm_qemu" "kube-controller" {
  count = var.vm_count
  name  = "vm-kube-controller"
  target_node = "laptop"
  clone = var.template_name
  agent    = 1
  os_type  = "cloud-init"
  cores    = 2
  sockets  = 1
  cpu      = "host"
  memory   = 2048
  scsihw   = "virtio-scsi-pci"
  bootdisk = "scsi0"
  disk {
    slot = 0
    size     = "50G"
    type     = "scsi"
    storage  = "local-lvm"
  }
  network {
    model  = "virtio"
    bridge = "vmbr0"
  }
  lifecycle {
    ignore_changes = [
      network,
    ]
  }
  ipconfig0 = "ip=192.168.0.210/24,gw=192.168.0.1"
  sshkeys = <<EOF
  ${var.ssh_key}
  EOF
}

resource "proxmox_vm_qemu" "kube-worker-1" {
  count = var.vm_count
  name  = "vm-kube-worker-1"
  target_node = "laptop"
  clone = var.template_name
  agent    = 1
  os_type  = "cloud-init"
  cores    = 6
  sockets  = 1
  cpu      = "host"
  memory   = 6144
  scsihw   = "virtio-scsi-pci"
  bootdisk = "scsi0"
  disk {
    slot = 0
    ssd = 1
    size     = "200G"
    type     = "scsi"
    storage  = "local-lvm"
  }
  network {
    model  = "virtio"
    bridge = "vmbr0"
  }
  lifecycle {
    ignore_changes = [
      network,
    ]
  }
  ipconfig0 = "ip=192.168.0.211/24,gw=192.168.0.1"
  sshkeys = <<EOF
  ${var.ssh_key}
  EOF
}

resource "proxmox_vm_qemu" "kube-worker-2" {
  count = var.vm_count
  name  = "vm-kube-worker-2"
  target_node = "pve"
  clone = var.template_name
  agent    = 1
  os_type  = "cloud-init"
  cores    = 4
  sockets  = 1
  cpu      = "host"
  memory   = 16384
  scsihw   = "virtio-scsi-pci"
  bootdisk = "scsi0"
  disk {
    slot = 0
    size     = "300G"
    type     = "scsi"
    storage  = "local-lvm"
    iothread = 1
  }
  network {
    model  = "virtio"
    bridge = "vmbr0"
  }
  lifecycle {
    ignore_changes = [
      network,
    ]
  }
  ipconfig0 = "ip=192.168.0.212/24,gw=192.168.0.1"
  sshkeys = <<EOF
  ${var.ssh_key}
  EOF
}

resource "proxmox_vm_qemu" "kube-worker-3" {
  count = var.vm_count
  name  = "vm-kube-worker-3"
  target_node = "pve"
  clone = var.template_name
  agent    = 1
  os_type  = "cloud-init"
  cores    = 4
  sockets  = 1
  cpu      = "host"
  memory   = 16384
  scsihw   = "virtio-scsi-pci"
  bootdisk = "scsi0"
  disk {
    slot = 0
    size     = "300G"
    type     = "scsi"
    storage  = "another"
    iothread = 1
  }
  network {
    model  = "virtio"
    bridge = "vmbr0"
  }
  lifecycle {
    ignore_changes = [
      network,
    ]
  }
  ipconfig0 = "ip=192.168.0.213/24,gw=192.168.0.1"
  sshkeys = <<EOF
  ${var.ssh_key}
  EOF
}
