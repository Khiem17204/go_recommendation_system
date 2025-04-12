Simple setup to get all the services running locally.

- install kubectl 
- install minikube
*for development, I am using OrbStack k8s cluster instead of minikube.* I have update scripts to standardize applying images to minikube.

minikube start

`./k8s-update.sh`

`kubectl create job manual-card-crawler --from=cronjob/card-crawler-job`
