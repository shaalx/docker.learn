#!/bin/sh
for d in `sudo docker ps -a -q`;do
#sudo docker start $d
sudo docker kill $d
sudo docker rm $d
echo $d
done

for i in `sudo docker images -a -q`;do
sudo docker rmi -f $i
echo $i
done
