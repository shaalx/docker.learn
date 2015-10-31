#!/bin/sh

for i in 1 2 3
do
echo $i
done

for i in $(seq 1 10);do
echo $i
done

for i in $[1..10];do
echo $i
done

i=0
while [ $i -lt 10 ]
do
i=$(($i+1))
echo $i*$i
done


