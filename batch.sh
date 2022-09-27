#!/bin/bash

start=`date +%s.%N`

for i in {1..10000}
do 
    make pub & 
done

end=`date +%s.%N`
runtime=$( echo "time: $end - $start" | bc -l )
