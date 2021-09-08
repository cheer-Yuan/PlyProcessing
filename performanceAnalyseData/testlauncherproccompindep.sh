#!/bin/bash
date

#clean the log file
rm ~/Desktop/Stages/gogs/output/*

#run the test
cd ~/Desktop/Stages/gogs/datasaving && ~/Desktop/Stages/gogs/datasaving/datasaving -n -1 -i 0 -f 0 -r 0 -d 0 &
cd ~/Desktop/Stages/gogs/src/dataprocessing && ~/Desktop/Stages/gogs/src/dataprocessing/dataprocessing &

date > ./autotest.txt
#length of the test equals to the product of the 2 number
for i in {1..10}
do
  # let the process run for the period of time
    sleep 20

  #read the number of files processed
  numsaved=$(($(tail ~/Desktop/Stages/gogs/datasaving/parameters.txt -n 1) + 1))
  numprocessed=$(tail ~/Desktop/Stages/gogs/src/dataprocessing/performanceAnalyse/readnum.txt -n 1)
  echo $i >> ./autotest.txt
  echo $numsaved >> ./autotest.txt
  echo $numprocessed >> ./autotest.txt
  echo -e '\n\n\n' >> ./autotest.txt

done

#stop the test and remove the process
killall -s 9 datasaving
killall -s 9 dataprocessing

date