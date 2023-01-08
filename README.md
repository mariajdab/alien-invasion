# alien-invasion

### assumptions
- If in the same city there are two aliens, then they will fight. If there are three or more is an assumption that the could have a good relationship.

- The file containing the names of cities in the non-existent world of X will be at the root of the project, if is needed to test with a new file please update the constant ```worldMap``` with the name of the file.

------------------------------------------------------------------------------------------------------------------------------------------------

The file will be read, and the data parsed in order to extract only the valuable information, each line of the file that contains the main city and its adjacent cities will be saved on a map of maps in Go. The directions will not be saved.

An example of this, consider that if the file ```cities_list.txt``` found in this repository is to be read by the program: 

A possible map of the firts cities (main cities) that appers in each line of the file ```cities_list.txt```, will be the following: 

![Screenshot from 2023-01-07 21-29-53](https://user-images.githubusercontent.com/48325352/211176741-5f574cb0-d299-40c7-be73-c5345e40b02b.png)

A possible map of the adjacents cities for the main city ```Foo```

![Screenshot from 2023-01-07 21-30-32](https://user-images.githubusercontent.com/48325352/211176743-85c5d064-9899-4682-9816-14b64daf912a.png)


#### Note: The order is not restricted



## Starting 🚀

To run the program please clone this repository:

```git clone https://github.com/mariajdab/alien-invasion.git```

run the file ```main.go``` it will need the argument of the number of aliens

`go run main.go <number_of_aliens>`
