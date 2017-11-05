# Setting up kubernetes with minikube

Simpliest way to install kubernetes is [minikube][]. After installation just
run `minikube start` it will start a cluster and will configure kubectl for
you. To stop cluster run `minikube` stop.

On arch it is (install docker-machine-kvm because virtualbox is super slow):

  yaourt minikube
  yaourt docker-machine-kvm

# Installing istio

  source setup-istio.sh
  kubectl apply -f $ISTIO_DIR/install/istio.yaml
  kubectl get svc -n istio-system

then wait until all pods are deployed, to ispect the status:

  kubectl get pod -n istio-system

# Building httpecho container

  images/httpecho/build.sh

# Deploying httpecho to kubernetes without istio

  kubectl create -f httpecho.yaml

Then check if that service is up and running

  http $(minikube service httpecho --url)/hello/world

Then simply delete it, because we will process with istio

  kubectl delete -f httpecho.yaml

# Deploying httpecho to kubernetes with istio

  kubectl create -f <(istioctl kube-inject -f httpecho.yaml)


[minikube]: https://github.com/kubernetes/minikube
