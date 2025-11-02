docker run -d -p 5000:5000 --name registry registry:2 --network bridge

#docker tag healthcheck:latest localhost:5000/healthcheck:latest
#docker push localhost:5000/healthcheck:latest

docker run --env-file ../.env  -d -p 5432:5432  --name db  postgres:16.9-alpine