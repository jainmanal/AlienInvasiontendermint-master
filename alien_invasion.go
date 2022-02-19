package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

var direction [1000][4]int
var city_list [12]string
var k = 0
var total_city = k
var alien_no int
var l = 0
var j int
var alien_city [1000]int
var city_removed int
var movement_made = 0
var content string
var text []string
var city_direction [12]string

// This function is created to Simulate Movement Randomly
func simulate(alien_city [1000]int, direction [1000][4]int, city_list [12]string, city_removed int, city_direction [12]string) {
	var i int
	movement_made = movement_made + 1

	// On Each movement a new city is assigned to alien over here
	for i = 0; i < len(alien_city); i++ {
		var new_city int
		var generate_direction int
		generate_direction = rand.Intn(4)
		new_city = direction[i][generate_direction]
		if alien_city[i] != 9997 {
			if new_city != 0 {
				alien_city[i] = new_city
			}
		}
	}

	// Here the condition is checked when two aliens are in same city and city is destroyed
	for i = 0; i < len(alien_city); i++ {
		for j = i + 1; j < len(alien_city); j++ {
			if (alien_city[i] == alien_city[j]) && (alien_city[i] != 9997) && city_list[alien_city[i]] != "" {

				t := strconv.Itoa(i)
				t1 := strconv.Itoa(j)
				content = city_list[alien_city[i]] + " has been destroyed by Alien" + t + " and Alien" + t1 + "\n"
				f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
					fmt.Println(err)
					return
				}
				f.WriteString(content)
				if err != nil {
					fmt.Println(err)
					f.Close()
					return
				}

				err = f.Close()
				if err != nil {
					fmt.Println(err)
					return
				}
				city_list[alien_city[i]] = ""
				alien_city[i] = 9997
				alien_city[j] = 9997
				city_removed = city_removed + 2
				break
			}
		}
	}

	// Over here the condition is checked if all cities are destroyed or movement is made more than 10000 times
	// if condition is not satisfied remaining cities are printed
	if city_removed < len(alien_city) && movement_made < 10000 {
		simulate(alien_city, direction, city_list, city_removed, city_direction)

	} else {
		for i = 0; i < len(city_direction); i++ {
			for j = i; j < len(city_list); j++ {
				words := strings.Fields(city_direction[i])
				if words[0] == city_list[j] {
					fmt.Println(city_direction[i])
				}
			}
		}

	}

}

func main() {
	var alien_no int

	// Taking input from user
	fmt.Printf("Enter No. Of Aliens: ")
	fmt.Scanln(&alien_no)

	// Condition to erase text file whenver is the code is executed
	if err := os.Truncate("test.txt", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	//Reading map file from here
	file, err := os.Open("map.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	var i = 0
	var k = 0
	for _, each_ln := range text {
		words := strings.Fields(each_ln)
		city_list[k] = words[0]
		city_direction[k] = each_ln
		k++
	}

	// Initiallizing the cities via map file to a 2d array over here
	for _, each_ln := range text {
		words := strings.Fields(each_ln)
		var j int
		for j = 1; j < len(words); j++ {
			v := strings.Split(words[j], "=")
			if v[0] == "east" {

				direction[i][0] = sort.StringSlice(city_list[:]).Search(v[1])
			}
			if v[0] == "west" {
				direction[i][1] = sort.StringSlice(city_list[:]).Search(v[1])
			}
			if v[0] == "north" {
				direction[i][2] = sort.StringSlice(city_list[:]).Search(v[1])
			}
			if v[0] == "south" {
				direction[i][3] = sort.StringSlice(city_list[:]).Search(v[1])
			}

		}
		i++

	}
	var total_city = k

	// Initializing the random cities to each alien before they move for the first time
	for l = 0; l < alien_no; l++ {
		alien_city[l] = rand.Intn(total_city)
	}

	// function to simulate movements are first time initialized here
	simulate(alien_city, direction, city_list, 0, city_direction)

}
