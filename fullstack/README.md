# fullstack

```bash
# install wordpress
helm upgrade --install wordpress bitnami/wordpress \
--create-namespace --namespace fullstack
# and next edit svc/wordpress change type LoadBalancer to ClusterIP
# then create ingress for wp
create ingress wordpress --rule="wordpress.adiatma.tech/*=wordpress:80,tls=wordpress.adiatma.tech"
# add cert-manager with lets-encrypt
kubectl annotate ingress/wordpress cert-manager.io/cluster-issuer=letsencrypt-production
```
