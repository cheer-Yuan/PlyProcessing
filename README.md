##Environment configuration :


To install the gonum package : go get -u gonum.org/v1/gonum/mat

To run the under commandline and get some outputs : go run main.go

#Version 1.1
2 functions added in PointcloudCalculator.go, using least square method and RANSAC

#Version 1.0
2 planes : A and B which form an angle between 90` and 180`
3 critic zones : A' c  A , B' c B and C' c (A U B)
We test the affiliation of the points in the critic zones to the planes

Formal tests waited to be done 

#Version 0.1 : 
test if a chosen vertex in 2021-03-30-14:45:57-.ply  belongs to the given plane

