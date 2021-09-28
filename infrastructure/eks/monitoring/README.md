# Provision EKS Cluster for Monitoring Project

## Preparation

### Install Terraform

### Create AWS Account

### Install AWS Cli

### Install Helm

## Provision EKS Cluster for Shoppping-Cart

### Initial Terraform

```cmd
terraform init

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

### Provision the EKS cluster

```cmd
terraform apply
```

### Configure kubectl

```cmd
aws eks --region $(terraform output -raw region) update-kubeconfig --name $(terraform output -raw cluster_name)
```

### Deploy Kubernetes Metrics Server

The Kubernetes Metrics Server, used to gather metrics such as cluster CPU and memory usage over time, is not deployed by default in EKS clusters.

Download and unzip the metrics server by running the following command.

```cmd
wget -O v0.3.6.tar.gz https://codeload.github.com/kubernetes-sigs/metrics-server/tar.gz/v0.3.6 && tar -xzf v0.3.6.tar.gz
```

Deploy the metrics server to the cluster by running the following command.

```cmd
kubectl apply -f metrics-server-0.3.6/deploy/1.8+/
```

Verify that the metrics server has been deployed. If successful, you should see something like this.

```cmd
kubectl get deployment metrics-server -n kube-system
NAME             READY   UP-TO-DATE   AVAILABLE   AGE
metrics-server   1/1     1            1           4s
```

### Deploy Kubernetes Dashboard

The following command will schedule the resources necessary for the dashboard.

```cmd
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml

namespace/kubernetes-dashboard created
serviceaccount/kubernetes-dashboard created
service/kubernetes-dashboard created
secret/kubernetes-dashboard-certs created
secret/kubernetes-dashboard-csrf created
secret/kubernetes-dashboard-key-holder created
configmap/kubernetes-dashboard-settings created
role.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrole.rbac.authorization.k8s.io/kubernetes-dashboard created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
deployment.apps/dashboard-metrics-scraper created
```

Now, create a proxy server that will allow you to navigate to the dashboard from the browser on your local machine. This will continue running until you stop the process by pressing CTRL + C.

```cmd
kubectl proxy
Starting to serve on 127.0.0.1:8001
```

You should be able to access the Kubernetes dashboard [here](http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/)

### Authenticate the dashboard

This step is very important. If you lack this step, you will report a lot of forbidden after opening the dashboard.

To use the Kubernetes dashboard, you need to create a ClusterRoleBinding and provide an authorization token. This gives the cluster-admin permission to access the kubernetes-dashboard. Authenticating using kubeconfig is not an option. You can read more about it in the Kubernetes documentation.

In another terminal (do not close the kubectl proxy process), create the ClusterRoleBinding resource.

```cmd
kubectl apply -f kubernetes-dashboard-admin.rbac.yaml 
```

Then, generate the authorization token.

```cmd
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')

eyJhbGciOiJSUzI1NiIsImtpZCI6ImE3UElQajY2Vlc0V1FmamhkSS1ia0o0U3BwYjRfeXBrVFlhbFFIWjBHb2MifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi11c2VyLXRva2VuLXdneGd6Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI0Mzc5OGM2OC05ZmU4LTQxMmQtYjUyYS0wMzRkZTJiYjkxOGUiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06YWRtaW4tdXNlciJ9.dSmHDRpulWYrGV4u_Wf3bm0MQL67oSWJrZQPZtaJhz-eZuFYJ9ObDC7C2OH_GMwTQARvTLvv_M1KR-TzkQifs145On1WRZOX_kJUd4c-ssd8Ii8C3lWfEH5cuW27uO0iD_5F3S5xge-uzA6N1xiFvtALhd4Pg0cNCQ6DkOEFDSxmQ7Qg65SVSpTR570BvZ75pCEIk2WRMeoTdecYRxNKaxVm0G52ASbwTz-vNdCmqdM1vCOLSmkQlgCpUlbWD7QnL3m9j7Jcq-Eavs2-3auDhqyh_KJ1nFoD4IJf5WKgKfAl2NGx8D6x1QLsX6AQ_p2q-9-mg72NM27r1WFuxMO-5Q
```

---

## Install Monitoring

## Create New namespace

```cmd
kubectl create namespace monitoring
```

### Set Current Namespace to Shopping-Cart Namespace

```cmd
kubectl config set-context --current --namespace=monitoring
```

### Install Prometeus

```cmd
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add stable https://charts.helm.sh/stable
helm repo update
```

```cmd
helm install shopping-cart prometheus-community/kube-prometheus-stack
```

- Verify that the pods and services are installed correctly.

    ```cmd
    kubectl get pods -l "release=shopping-cart"
    kubectl get service -l "release=shopping-cart"
    ```

- Updated Grafana Service to LoadBalancer

    converted the Prometheus service into a LoadBalancer by patching it with the below command:

    ```cmd
    kubectl patch svc shopping-cart-grafana -p '{"spec": {"type": "LoadBalancer"}}'
    ```

- Login with Grafana

    get URL for Grafana

    ```cmd
    kubectl get service | grep shopping-cart-grafana| awk '{print $4}'

    adc790c671ed64f528456e7f6750e716-621529730.ap-southeast-1.elb.amazonaws.com
    ```

    go to grafana url and login with user/password -> admin / prom-operator

    [Default Grafana Password for ube-prometheus-stack Helm Chart](https://www.google.com/search?q=kube-prometheus-stack+grafana+password&oq=kube-prometheus-stack+&aqs=chrome.2.69i57j0i512j0i20i263i512j0i512l4j69i60.2763j0j4&sourceid=chrome&ie=UTF-8)

## Destroy

อย่าลืมลบ Load Balancer ก่อน EC2 -> Load Balancing -> Load Balancers

## Reference

### Learn Terraform - Provision an EKS Cluster

This repo is a companion repo to the [Provision an EKS Cluster learn guide](https://learn.hashicorp.com/terraform/kubernetes/provision-eks-cluster), containing
Terraform configuration files to provision an EKS cluster on AWS.

### Reference Link

- [How do I expose the Kubernetes services running on my Amazon EKS cluster? - Create a LoadBalancer service](https://aws.amazon.com/th/premiumsupport/knowledge-center/eks-kubernetes-services-cluster/)
- [Provisioning Kubernetes clusters on AWS with Terraform and EKS](https://learnk8s.io/terraform-eks)
- [The eksctl command line utility](https://docs.aws.amazon.com/eks/latest/userguide/eksctl.html)
- [Kubernetes in Production: The Ultimate Guide to Monitoring Resource Metrics with Prometheus](https://www.replex.io/blog/kubernetes-in-production-the-ultimate-guide-to-monitoring-resource-metrics)
- [How do I resolve the "Your current user or role does not have access to Kubernetes objects on this EKS cluster" error in Amazon EKS?](https://aws.amazon.com/th/premiumsupport/knowledge-center/eks-kubernetes-object-access-error/)
- [Kubernetes 1.13 cluster install dashboard 1.10.1](https://blog.titanwolf.in/a?ID=01400-c3606239-eb96-4373-8315-d9a19bf36f65)
