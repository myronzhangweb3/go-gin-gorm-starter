#!/bin/bash
version=$1
echo 'build version: '${version}

sudo make build

repo_name="w3ifun/training_collector_serv"
image_name="${repo_name}:${version}"
latest="${repo_name}:latest"
sudo docker rmi -f $(docker images -q ${repo_name})
sudo docker build --pull -t "$latest" -t "$image_name" .
sudo docker push "$repo_name" -a