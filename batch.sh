#!/bin/bash

for i in {1..10}
do 
    echo "$i"
    make pub & 
done