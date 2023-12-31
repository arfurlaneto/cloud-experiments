## Description

Kind cluster with a simple Go server featuring:

- Environment variables set from `ConfigMap`
- Environment variables set from `Secret`
- File mounted from `ConfigMap`
- Folder mounted from `PersistentVolumeClaim`
- HPA and metrics-server
- Startup probes
- Ingress
- `Deployment` with custom `ServiceAccount`

## Instructions

Starting:

```bash
make kind-create
make kind-load-images
make k8s-setup
```

Testing:

```bash
curl http://localhost:80/ -H "Host: kubernetes.local"
curl http://localhost:80/configmapenv -H "Host: kubernetes.local"
curl http://localhost:80/configmapfile -H "Host: kubernetes.local"
curl http://localhost:80/secret -H "Host: kubernetes.local"
```

Stopping:

```bash
make kind-destroy
```
