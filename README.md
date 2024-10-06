# Odd API Server - Steps to run locally

## Install required tools

### Kind, Helm, Kubectl

```bash
brew install kind
brew install helm
```

for `kubectl`,see here: https://kubernetes.io/docs/tasks/tools/#kubectl

## Create the kind cluster locally

```bash
kind create cluster --name test-cluster
```

Switch to the cluster context:

```bash
kubectl config use-context kind-test-cluster
```

## Install metrics-server

_**Note: the metrics-server is required for the HPA (Horizontal Pod Autoscaler) to work.**_

```bash
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

Edit its deployment and add to `--args`:

```bash
kubectl edit deployment metrics-server -n kube-system
```

```
- --kubelet-insecure-tls
- --kubelet-preferred-address-types=InternalIP
```

## Install ingress-nginx controller

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace
```

## Deploy the application

```bash
helm install odd-api-server ./charts/odd-api-server
```

## Configure port forwarding for the nginx ingress controller service

```bash
kubectl port-forward svc/nginx-ingress-ingress-nginx-controller 8080:80 -n ingress-nginx
```

## Add host to /etc/hosts to immitate real domain

Edit `/etc/hosts`:

```bash
sudo vim /etc/hosts
```

Add:

```
127.0.0.1 odd-api-server.local
```

## Test the application endpoints

```bash
curl http://odd-api-server.local:8080/odd
curl http://odd-api-server.local:8080/ready
```

## Install Prometheus Operator (Bonus)

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install odd-prometheus prometheus-community/kube-prometheus-stack
```
