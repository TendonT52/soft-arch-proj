
wget https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img
apt install libguestfs-tools -y
virt-customize -a jammy-server-cloudimg-amd64.img --install qemu-guest-agent

qm create 9001 --name "ubuntu-2023-cloudinit-template" --memory 2048 --cores 2 --net0 virtio,bridge=vmbr0
qm importdisk 9001 jammy-server-cloudimg-amd64.img local-lvm
qm set 9001 --scsihw virtio-scsi-pci --scsi0 local-lvm:vm-9001-disk-0
qm set 9001 --boot c --bootdisk scsi0
qm set 9001 --ide2 local-lvm:cloudinit
qm set 9001 --serial0 socket --vga serial0
qm set 9001 --agent enabled=1

qm template 9001