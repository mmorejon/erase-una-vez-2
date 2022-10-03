# Once upon a time 2

Golang application used in the exercises in the book [Erase una vez Kubernetes](https://leanpub.com/erase-una-vez-kubernetes).

Translation: [Español](README.md)

## Description

The code is used to do communication examples between two applications deployed in the cluster: a client and a server. Both applications have been compiled on the same Docker image, but the start command is different.

### Server service features

Use the `server` command to start the service as a web server. The application provides the following access points:

|Access point|Description|
|-----|-----------|
|`/echo`|Displays the message `'erase una vez ...'` along with the name of the machine.|
|`/healthz`|Displays the OK message if the server is working properly.|

All access points print the information in JSON format.

|Env variable|Description|Type|Default|
|-----|-----------|------|---|
|`CHARACTER`| String added at the end of the message.| String | `"..."` |

### Client service features

Use the `client` command to start the system as a client service. The application makes requests towards the configured link by time intervals. To configure the client application use the following environment variables:

|Env variable|Description|Type|Default|
|-----|-----------|------|---|
|`SLEEP_TIME`| Time interval between requests.| String | `"1s"` |
|`SERVER_URL`| URL used to make the requests. e.g: `http://example.com`.| String | `""` |
|`ENDPOINT`| URL path used to make the request.| String |  `""` |

## App behavior at localhost

```bash
# starting web server
docker container run --rm \
  --name server \
  --entrypoint server \
  -p 8000:8000 \
  --detach \
  ghcr.io/mmorejon/erase-una-vez-2:v0.3.0

2020/01/20 23:02:24 Servidor iniciado
```

```bash
curl http://localhost:8000/echo

{
  "hostname": "da1ebc56480a",
  "message": "érase una vez ..."
}%
```

```bash
curl http://localhost:8000/healthz
{
  "status": "OK"
}%
```

```bash
# get server ip
docker container inspect \
  --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' \
  server

172.17.0.2
```

```bash
# starting web server
docker container run --rm \
  --name client \
  --entrypoint client \
  --detach \
  --env SERVER_URL="http://172.17.0.2:8000" \
  --env ENDPOINT="/echo" \
  ghcr.io/mmorejon/erase-una-vez-2:v0.3.0
```

```bash
# getting logs from client app
docker container logs client
{
  "hostname": "9e65a5deb2f7",
  "message": "érase una vez ..."
}
{
  "hostname": "9e65a5deb2f7",
  "message": "érase una vez ..."
}
```

## App behavior at Kubernetes cluster

```bash
kubectl apply -f kubernetes/

deployment.apps/client created
deployment.apps/server created
service/server created
```

```bash
kubectl get all -l app=server

NAME                        READY   STATUS    RESTARTS   AGE
pod/server-b4746684-wgfxm   1/1     Running   0          6m46s

NAME             TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/server   ClusterIP   10.109.117.210   <none>        8000/TCP   6m46s

NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/server   1/1     1            1           6m46s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/server-b4746684   1         1         1       6m46s
```

```bash
kubectl get all -l app=client

NAME                         READY   STATUS    RESTARTS   AGE
pod/client-7dd5bc4b4-j6qg4   1/1     Running   1          7m41s

NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/client   1/1     1            1           7m41s

NAME                               DESIRED   CURRENT   READY   AGE
replicaset.apps/client-7dd5bc4b4   1         1         1       7m41s
```
