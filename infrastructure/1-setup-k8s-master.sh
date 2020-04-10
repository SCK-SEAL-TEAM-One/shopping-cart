sudo ufw disable
sudo systemctl disable ufw
sudo kubeadm init --kubernetes-version v1.13.0 --ignore-preflight-errors=all