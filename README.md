# 🌐 stunning-octo-spork

Two services deployed with Kubernetes. Written in go.

### 🚀 Run the Kubernetes setup:

1. Create the namespaces

```bash
kubectl create ns proxy-ns
kubectl create ns internal0ns
```

2. Create the cluster (I used Kind, you can use other clusters too):

```bash
kubectl apply -f kubernetes/
```

3. Port forward from inside the cluster (just for testing purposes):

```bash
kubectl port-forward service/proxy-service 8080:5665
```

4. Test it with postman (make a get request to `http://localhost:8080/proxy`)
