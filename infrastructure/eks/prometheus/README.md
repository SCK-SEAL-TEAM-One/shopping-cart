# Prometheus into Shopping-Cart Cluster

## Deploy Mysql NodeExporter

```cmd
kubectl apply -f mysql-node-exporter-deployment.yml
kubectl apply -f mysql-node-exporter-service.yml
```

## Deploy Nginx NodeExporter

```cmd
kubectl apply -f nginx-node-exporter-deployment.yml
kubectl apply -f nginx-node-exporter-service.yml
```

## Deploy Prometheus

```cmd
kubectl apply -f prometheus-configmap.yml
kubectl apply -f prometheus-deployment.yml
kubectl apply -f prometheus-service.yml
```
