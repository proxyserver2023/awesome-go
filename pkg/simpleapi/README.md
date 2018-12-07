1. Run MYSQL in docker-container:

``` bash
mkdir -p /host # for accessing from localhost
docker run --name my-mysql                               \
           -e     MYSQL_ROOT_PASSWORD=my-secret-password \
           -v     /host:/shared                          \
           mysql
```
