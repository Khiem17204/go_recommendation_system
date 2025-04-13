Simple setup to get all the services running locally.

- install kubectl 
- install minikube
*for development, I am using OrbStack k8s cluster instead of minikube.* I have update scripts to standardize applying images to minikube.

minikube start

`./k8s-update.sh`

`kubectl create job manual-card-crawler --from=cronjob/card-crawler-job`

`kubectl create job manual-deck-crawler --from=cronjob/deck-crawler-job`

create deck embedding, this takes quite a while

`kubectl apply -f config/deck-embedding-worker.yaml`