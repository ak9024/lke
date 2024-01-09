# LKE (Linode Kubernetes Engine)

### Prerequisite

- `kubectl` (https://kubernetes.io/docs/tasks/tools/)
- `kubectx` (https://github.com/ahmetb/kubectx)
- `helm` (https://helm.sh/docs/intro/install/)

### Getting Started

Setup kubernetes config.

```bash
# download new config from LKE and copy to .kube/config
KUBECONFIG=~/Downloads/config.yaml kubectl config view --raw > ~/.kube/config
# merge old config with a new config
mv .kube/config .kube/config.old
KUBECONFIG=.kube/config.old:~/path/config.yaml kubectl config view —raw > .kube/config
rm -rf .kube/config.old
```

Setup **external-dns** and **LINODE_API_TOKEN** to usage external-dns in LKE cluster.

```bash
# add helm repo
helm add bitnami https://charts.bitnami.com/bitnami
# install external-dns
helm upgrade --install external-dns bitnami/external-dns \
--namespace external-dns --create-namespace \
--set provider=linode \
--set linode.apiToken=$LINODE_API_TOKEN # create api token for domain in linode
```

Setup **traefik** for HTTP reverse proxy and load balancer.

```bash
# add helm repo
helm repo add traefik https://traefik.github.io/charts
# install traefik
helm upgrade --install traefik traefik/traefik \
--create-namespace --namespace traefik \
--set "ports.websecure.tls.enable.enabled=true" \
--set "providers.kubernetesIngress.publishedService.enabled=true"
# force redirect http to https
kubectl edit deploy/traefik
# and add
# spec.containers[].args
- --entrypoints.web.http.redirections.entryPoint.to=:443

# traefik dashboard
# https://github.com/traefik/traefik-helm-chart/blob/master/EXAMPLES.md
kubectl port-forward deploy/traefik 9000:9000
# access
# http://localhost:9000/dashboard
```

Setup **metrics-server** to enabled API Metrics.

```bash
# add helm repo
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
# install
helm upgrade --install metrics-server metrics-server/metrics-server \
--create-namespace --namespace metrics-server

# next
kubectl edit deployment/metrics-server
# set config https://github.com/kubernetes-sigs/metrics-server
--kubelet-preferred-address-types=InternalIP # set the priority the address type
--kubelet-insecure-tls # Do not verify the CA of serving certificates presented by Kubelets. For testing purposes only.

# check metrics-server
kg apiservices | grep metrics
v1beta1.metrics.k8s.io                 metrics-server/metrics-server   True  
```

Setup **monitoring** with `kube-prometheus-stack`

```bash
# add repo
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
# install kube-prometheus-stack
helm upgrade --install kube-prometheus-stack prometheus-community/kube-prometheus-stack \
--create-namespace --namespace kube-prometheus-stack
```

Setup **SSL/TLS** with `cert-manager`

```bash
# add repo
helm repo add cert-manager https://charts.jetstack.io
# helm install
helm upgrade --install cert-manager cert-manager/cert-manager \
> --create-namespace --namespace cert-manager \
> --set installCRDs=true
```

Create `cluster-issuer.yaml` and `kubectl apply -f cluster-issuer.yaml`

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-production
spec:
  acme:
    email: adiatma.mail@gmail.com
    # to use staging environment 
    # server: https://acme-staging-v02.api.letsencrypt.org/directory
    # to use the production environment
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: issuer-letsencrypt-production
    solvers:
    - http01:
       ingress:
         class: traefik
```

Expose the service to ingress

```bash
kubectl create ingress <ingress-name> --rule="<domain.com>/*=<service:port>,tls=<domain.com>"
kubectl annotate ing/<ingress-name> cert-manager.io/cluster-issuer=letsencrypt-production
```
