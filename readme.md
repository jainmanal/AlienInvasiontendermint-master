ALien Invasion Tendermint Solution

In the solution i provided it contains three main files
1. alien_invasion.go is the main file which contains all the logic created. There are two functions in the file
one contains the main initialization and second contains the simulation created.
2. map.txt is the file used to add city names and directions.
3. test.txt is the output file which stores the destroyed city names.


Assumption made - I have used array to store values hence city_name array has a size of 12. For aliens I have made test till 1000 aliens at once. 

Logic -
After reading the question i created the following logic.
1. In the main function i created logic to read the map.txt at first.
2. After that i seprated each line of map.txt on the basis of city name and direction and stored them in array of city and 2d array of direction.
3. After that i assigned each alien taken as input from user a random city and assigned them there city ids.
4. After this the main function calls the simulation function.
5. In the first step in simulation function a random direction is assigned to each alien and on the basis of the alien walks to a new city based on the direction.
6. After this condition to check which aliens are in same city and on the basis of that i remove the alien from alien list and destroy city from city list.
7. Cities destroyed on each iteration is updated in the test.txt file.
8. In final step the remaining cities are printed if there are any if 10000 simulations are completed.


Command to Run Code
golang run alien_invasion.go

After this you need to input the number of aliens