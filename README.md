##Environment configuration :


To install the gonum package : go get -u gonum.org/v1/gonum/mat

To run the under commandline and get some outputs : go test -v -run Exp_ main_test.go (see the instructions here below)

#Exp4

This is the test for function PlaneConsecRANSAC, an improved version of the function in the previous test . A process of batch composition is added to the beginning of each iteration. RANSAC will be applied to the vertices in the batch. Then the plane obtained will be passed to the generalization and re-adjustment. Another process of redundancy detection is also added to check whether the new plane is the same as another one.

This test is run on the point cloud of 3 perpendicular planes (the same as the last test) and also on the point cloud of a desk, which is more complicated (./data/desk_static_complex_L515)

The results are very exciting : during several repeated tests, all of 3 perpendicular angles are successfully detected. As for the second point cloud, we detected many angles, including many perpendicular angles (may come from the desk, the wall, the boxes on the desk, and etc.)

We observed that by changing the parameter minVforPlane (the minimal number of the vertices required to compose a plane)

The next step may be the visualization of these results to make it observable more directly

To run under the commandline and get some outputs : go test -v -run Exp4 main_test.go


#Exp3

This is the test for function PlaneSeqRANSAC in PlaneExtractor.go. It is the first version of an automatic plane detector. 


#Exp2

In this experiment we do the same thing as the Exp1 but on the data captured by de depth camera D455. The relevant data is located in ./data/c_p_g_ofD455_highview.

The position of the camera is changed since D455 is not good at detecting a very short distance as the L515. A capture from the above avoid the floor at a short distance.

To run under the commandline and get some outputs : go test -v -run Exp1 main_test.go

Result : All 3 perpendicular angles are rightly detected and measured

#Exp1

In this experiment we define 3 critical zones (zones d'intérêts) from the .ply file in ./data/angle_of_board_and_paper_card_on_table. From each of these zones we fit a surface equation and measure the angles formed by them

The data is captured by the LiDAR camera Intel L515

To run under the commandline and get some outputs : go test -v -run Exp2 main_test.go

Result : All 3 perpendicular angles are rightly detected and measured

#Version 1.1
2 functions added in PointcloudCalculator.go, using least square method and RANSAC

#Version 1.0
2 planes : A and B which form an angle between 90` and 180`
3 critic zones : A' c  A , B' c B and C' c (A U B)
We test the affiliation of the points in the critic zones to the planes

Formal tests waited to be done 

#Version 0.1 : 
test if a chosen vertex in 2021-03-30-14:45:57-.ply  belongs to the given plane

