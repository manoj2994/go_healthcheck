#``
kubectl apply -f configMap.yaml
kubectl apply -f secrets.yaml
kubectl apply -f db/psql-pv.yaml
kubectl apply -f db/psql-pvc.yaml
kubectl apply -f db/psql.yaml
kubectl apply -f health_app.yaml
kubectl apply -f health_app-service.yaml