#!/bin/sh

for i in 1 2 3
do
echo $i
done

let i=0
while(($i < 10))
do
i=$(($i+1))
echo $i*$i
done


