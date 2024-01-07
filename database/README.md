# database

```bash
# install postgresql
helm upgrade --install postgresql bitnami/postgresql \
--create-namespace --namespace database \
--set global.postgresql.auth.postgresPassword=<password>
```

```bash
# install pgadmin4
helm upgrade --install pgadmin4 runix/pgadmin4 \
--create-namespace --namespace database \
--set env.email=adiatma.mail@gmail.com \
--set env.password=<password>
```
