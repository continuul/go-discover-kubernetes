# Development

## Testing with Minikube

### Environment

### Starting Minikube

Below are instructions for testing a Docker image locally under a realistic test scenario. The instructions may be used for verifying a built image (release candidate) or for verifying changes before submitting a PR.

The primary benefit of this configuration is that you do not have to set up a private registry, that you don't have to push to Docker Hub before hand; this setup shares the registry namespace of the native Docker daemon with the running Kubernetes cluster, so built images show up as cached by the Docker daemon in the cluster - thus it doesn't have to download an image from elsewhere (clever, huh?).

So lets jump onto the steps...

First you will need to install Kubectl and Minikube:

```bash
curl -Lo kubectl https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/darwin/amd64/kubectl && chmod +x ./kubectl && sudo mv kubectl /usr/local/bin/kubectl

curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.23.0/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
```

Then you will need to start Minikube and configure it:

```bash
minikube stop
minikube delete

minikube start --memory 8192 --cpus 4 --vm-driver=virtualbox --extra-config=kubelet.CAdvisorPort=4194
# (optional) --extra-config=apiserver.Authorization.Mode=RBAC
sleep 30
minikube addons enable dashboard
minikube addons enable default-storageclass
minikube addons enable heapster
minikube addons enable kube-dns
sleep 60
```

To (optionally) open a proxy port to the API server:

```bash
kubectl proxy --address 0.0.0.0 --accept-hosts '.*'
```

To open the Kubernetes Dashboard:

```bash
minikube dashboard
```

To open the Kubernetes Heapster (Grafana):

```bash
minikube addons open heapster
```

### Building a Docker Image

To build a local image which can be used by Minikube
during deployment, run the following:

```bash
eval $(minikube docker-env)
make docker
docker tag continuul/on:0.1.2 continuul/on:latest
```

### Deployments

```bash
kubectl create namespace continuul
```

#### Single Node

```bash
kubectl create -f on-template.yaml
```

#### Multi Node

```bash
todo...
```

### Validating a Deployment

Open the Kubernetes Dashboard, and view the logs for the deployed pod, ensuring the nodes peer up
and no errors appear in the logs. Alternatively, you can issue view them at the command line:

```bash
kubectl logs -f <image-name>
```

Verify the correct version number is displayed; if it's not, check your environment.
