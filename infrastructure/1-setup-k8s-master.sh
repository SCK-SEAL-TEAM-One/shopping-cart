export KUBE_MASTER=ec2-13-229-235-183.ap-southeast-1.compute.amazonaws.com
ssh -i sck_default.pem ubuntu@$KUBE_MASTER
sudo ufw disable
sudo systemctl disable ufw
sudo kubeadm init --kubernetes-version v1.13.0 --ignore-preflight-errors=all