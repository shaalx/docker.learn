#!/bin/sh
for i in `docker images -q`;do
docker inspect --format="{{.RepoTags}}" $i | grep :test && docker rmi -f $i
echo 'delete:'+$i
done

