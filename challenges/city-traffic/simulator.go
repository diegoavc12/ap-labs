package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	 "github.com/golang-collections/collections/stack"
	//"stack"
)

var numberOfCars int
var numberOfStoplights int
var r *rand.Rand
var city [][]sector
var paths [][]sector
var cars []car
var streets []sector
var crosses [][]sector
var outputFile *os.File
var err error

// Map matrix values
const LEFT = 0
const RIGHT = 1
const DOWN = 2
const UP = 3
const LEFT_DOWN = 4
const RIGHT_DOWN = 5
const LEFT_UP = 6
const RIGHT_UP = 7
const STREET = 0
const BUILDING = 1
const WIDTH = 16


type point struct {
	xPosition int
	yPosition int
}

type stoplight struct {
	sectors []sector
	index int
	speed int
}

type car struct {
	id       int
	position point
	speed    int
	path     []sector
	inmobile int
}

type sectorQueue struct {
	previous *sectorQueue
	c        sector
}

type sector struct {
	position   point
	cellType   int
	direction  int
	isOcuppied bool
	greenLight bool
}

func main() {
	outputFile, err = os.Create("traffic-output.txt")
	if err != nil {
		log.Fatalf("Error trying to create file")
	}
	defer outputFile.Close()

	getParameters()
	
	initializeProgram()
	
	ch := make(chan int, numberOfCars)
	createCars(streets, &ch)
	showSimulation(&ch)
	fmt.Println("Simulation finished. Results can be found in the output file")
}

func getParameters() {
	fmt.Print("Enter number of cars (Less than 16): ")
	var cars int
	fmt.Scan(&cars)
	numberOfCars = cars
	
	fmt.Print("Enter number of stoplights: ")
	var sl int
	fmt.Scan(&sl)
	numberOfStoplights = sl
}

func initializeProgram() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	city, streets = populateCityMatrix()
	
	verifyCarParameter()
	
	crosses = createCrosses()
	
	verifyStoplightParameter(len(crosses))

	createStoplights(crosses)
}

func verifyStoplightParameter(numberOfCrosses int) {
	if numberOfStoplights < 0 {
		log.Fatalf("Please use a positive number for the number of stoplights")
	}
	if numberOfStoplights > numberOfCrosses {
		log.Fatalf("There's no need for that many stoplights, we only have %v crosses", numberOfCrosses)
	}

}

func verifyCarParameter() {
	if numberOfCars > WIDTH {
		log.Fatalf("Please use less than 16 cars, try again")
	}
	if numberOfCars < 0 {
		log.Fatalln("Please use a positive number for the number of cars")
	}
}

func populateCityMatrix() ([][]sector, []sector) {
	var city = make([][]sector, 0)
	var streets = make([]sector, 0)
	for i := 0; i < WIDTH; i++ {
		var path = make([]sector, 0)
		x := i % 7

		for j := 0; j < WIDTH; j++ {
			y := j % 7

			if x < 2 || y < 2 {
				dir := -1

				switch {
				case x == 0 && y == 0:
					dir = LEFT_DOWN
				case x == 1 && y == 0:
					dir = RIGHT_DOWN
				case x == 0 && y == 1:
					dir = LEFT_UP
				case x == 1 && y == 1:
					dir = RIGHT_UP
				case x == 0:
					dir = LEFT
				case x == 1:
					dir = RIGHT
				case y == 0:
					dir = DOWN
				case y == 1:
					dir = UP
				}

				sectorA := sector{point{i, j}, STREET, dir, false, true}
				if dir == LEFT || dir == RIGHT || dir == UP || dir == DOWN {
					streets = append(streets, sectorA)
				}
				path = append(path, sectorA)
			} else {
				sectorA := sector{point{i, j}, BUILDING, -1, false, true}
				path = append(path, sectorA)
			}
		}
		city = append(city, path)
	}

	// print(city)
	return city, streets

}

func writeCity() {
	speedStr := ""
	for i := 0; i < WIDTH; i++ {
		path := ""
		for j := 0; j < WIDTH; j++ {
			if city[i][j].isOcuppied {
				path += " ■ "
			} else if !city[i][j].greenLight {
				d := city[i][j].direction
				if d == LEFT || d == RIGHT {
					path += "==="
				}
				if d == DOWN || d == UP {
					path += " ║ "
				}
			} else {
				switch city[i][j].cellType {
				case STREET:
					switch city[i][j].direction {
					case LEFT:
						path += "==="
					case RIGHT:
						path += "==="
					case DOWN:
						path += " ║ "
					case UP:
						path += " ║ "
					case LEFT_DOWN:
						path += "  ┼"
					case RIGHT_DOWN:
						path += "  ┼"
					case LEFT_UP:
						path += "┼  "
					case RIGHT_UP:
						path += "┼  "
					default:
						path += "   "
					}
					break
				case BUILDING:
					path += "···"
				}
			}
		}
		
		if i < numberOfCars {
			cIndex := strconv.Itoa(i)
			if cars[i].speed != 0 {
				if cars[i].speed > 240 {
					speedStr += "Car " + cIndex + " speed: 0 km/h"
				} else {
					speedStr += "Car " + cIndex + " speed: " + strconv.Itoa(1250 / cars[i].speed) + " km/h \n"
				}
			} 
		}
		
		outputFile.WriteString(path + "\n")
		outputFile.Sync()
	}
	outputFile.WriteString("\n" + speedStr + " \n")
}

func createStoplights(crosses [][]sector) {
	stoplights := make([]stoplight, 0)
	for i := 0; i < numberOfStoplights; i++ {
		n := r.Intn(len(crosses))

		cross := crosses[n]
		crosses[len(crosses)-1], crosses[n] = crosses[n], crosses[len(crosses)-1]

		crosses = crosses[:len(crosses)-1]
		speed := r.Intn(1200-800) + 800
		s := stoplight{cross, 0, speed}

		initialState(&s)

		stoplights = append(stoplights, s)
	}
	// creating stop light threads
	for i := 0; i < len(stoplights); i++ {
		index := i
		go func() {
			for {
				changeState(&stoplights[index])
				time.Sleep(time.Duration(stoplights[index].speed) * time.Millisecond)
			}
		}()
	}
}




func createCrosses() [][]sector {
	crosses := make([][]sector, 0)
	for i := 0; i < WIDTH; i += 7 {
		for j := 0; j < WIDTH; j += 7 {
			cross := make([]sector, 0)
			if i > 0 {
				cross = append(cross, city[i-1][j])
			}
			if j > 0 {
				cross = append(cross, city[i+1][j-1])
			}
			if i < WIDTH-2 {
				cross = append(cross, city[i+2][j+1])
			}
			if j < WIDTH-2 {
				cross = append(cross, city[i][j+2])
			}
			crosses = append(crosses, cross)
		}
	}
	return crosses
}

func createCars(streets []sector, ch *chan int) {
	cars = make([]car, 0)
	for i := 0; i < numberOfCars; i++ {
		n := r.Intn(len(streets))
		sectorA := streets[n]

		streets[len(streets)-1], streets[n] = streets[n], streets[len(streets)-1]
		streets = streets[:len(streets)-1]

		n2 := r.Intn(len(streets))
		sectorB := streets[n2]

		speed := r.Intn(250-50) + 50

		path := getPath(sectorA, sectorB)
		paths = append(paths, path)

		carA := car{i, point{sectorA.position.xPosition, sectorA.position.yPosition}, speed, path, 0}

		cars = append(cars, carA)
		addCar(carA)
	}
	// creating car threads
	for i := 0; i < len(cars); i++ {
		index := i
		go func() {
			for len(cars[index].path) > 0 {
				time.Sleep(time.Duration(cars[index].speed) * time.Millisecond)
				moveCar(&cars[index])
			}
			*ch <- cars[index].id
			cars[index].speed = 0
			deleteCar(&cars[index])
		}()
	}
}

func initialState(stopA *stoplight) {
	for i := 0; i < len(stopA.sectors); i++ {
		x := stopA.sectors[i].position.xPosition
		y := stopA.sectors[i].position.yPosition

		city[x][y].greenLight = false
	}
}

func changeState(stopA *stoplight) {
	length := len(stopA.sectors)

	secAX := stopA.sectors[stopA.index].position.xPosition
	secAY := stopA.sectors[stopA.index].position.yPosition

	city[secAX][secAY].greenLight = false
	stopA.index = (stopA.index + 1) % length

	secBX := stopA.sectors[stopA.index].position.xPosition
	secBY := stopA.sectors[stopA.index].position.yPosition

	city[secBX][secBY].greenLight = true
}



func addCar(carA car) {
	i := carA.position.xPosition
	j := carA.position.yPosition
	if !city[i][j].isOcuppied {
		city[i][j].isOcuppied = true
	}
}

func showSimulation(ch *chan int) {
	for {
		writeCity()
		outputFile.WriteString("Cars finished:" + strconv.Itoa(len(*ch)) + "/" + strconv.Itoa(numberOfCars) + "\n")
		outputFile.WriteString("----------------------------------------------------------------\n")
		outputFile.Sync()
		if len(*ch) >= numberOfCars {
			break
		}
	}
	pathsStr := writePaths()
	outputFile.WriteString(pathsStr)
}

func moveCar(carA *car) {
	carPosX := carA.position.xPosition
	carPosY := carA.position.yPosition

	nextSector := carA.path[0]

	secPosX := nextSector.position.xPosition
	secPosY := nextSector.position.yPosition

	if city[carPosX][carPosY].greenLight && !city[secPosX][secPosY].isOcuppied {
		city[carPosX][carPosY].isOcuppied = false

		carA.position.xPosition = secPosX
		carA.position.yPosition = secPosY

		city[secPosX][secPosY].isOcuppied = true
		carA.path = carA.path[1:]

		if carA.speed > 50 {
			carA.speed -= 10
		}

		carA.inmobile = 0

	} else {
		if carA.inmobile <= 2 {
			carA.inmobile++

			if carA.speed < 250 {
				carA.speed += 10
			}

		} else {
			carA.speed = 250
		}
	}
}

func deleteCar(carA *car) {
	secX := carA.position.xPosition
	secY := carA.position.yPosition
	
	if city[secX][secY].isOcuppied {
		city[secX][secY].isOcuppied = false
	}
}

func getPath(source sector, destination sector) []sector {
	q := sectorQueue{nil, source}
	vis := make([]sector, 0)
	queue := make([]sectorQueue, 0)
	queue = append(queue, q)

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.c == destination {
			return buildPath(&curr)
		}
		vis = append(vis, curr.c)
		negh := findNegh(vis, curr.c)
		for i := 0; i < len(negh); i++ {
			q2 := sectorQueue{&curr, negh[i]}
			if !visitedAlr(vis, q2.c) {
				queue = append(queue, q2)
			}
		}
	}
	return nil
}

func findNegh(vis []sector, src sector) []sector {
	var negh = make([]sector, 0)
	dir := src.direction
	secPosX := src.position.xPosition
	secPosY := src.position.yPosition

	if dir == LEFT_DOWN || dir == LEFT || dir == LEFT_UP {
		if secPosY > 0 {
			cityA := city[secPosX][secPosY-1]
			if !visitedAlr(vis, cityA) && cityA.cellType != BUILDING {
				negh = append(negh, cityA)
			}
		}
	}

	if dir == DOWN || dir == LEFT_DOWN || dir == RIGHT_DOWN {
		if secPosX < WIDTH-1 {
			cityA := city[secPosX+1][secPosY]
			if !visitedAlr(vis, cityA) && cityA.cellType != BUILDING {
				negh = append(negh, cityA)
			}
		}
	}

	if dir == RIGHT || dir == RIGHT_DOWN || dir == RIGHT_UP {
		if secPosY < WIDTH-1 {
			cityA := city[secPosX][secPosY+1]
			if !visitedAlr(vis, cityA) && cityA.cellType != BUILDING {
				negh = append(negh, cityA)
			}
		}
	}

	if dir == UP || dir == LEFT_UP || dir == RIGHT_UP {
		if secPosX > 0 {
			cityA := city[secPosX-1][secPosY]
			if !visitedAlr(vis, cityA) && cityA.cellType != BUILDING {
				negh = append(negh, cityA)
			}
		}
	}

	return negh
}

func visitedAlr(vis []sector, sectorA sector) bool {
	length := len(vis)
	for i := 0; i < length; i++ {
		if vis[i] == sectorA {
			return true
		}
	}
	return false
}

func buildPath(queue *sectorQueue) []sector {
	path := make([]sector, 0)
	s := stack.New()
	curr := queue
	for curr != nil {
		s.Push(curr.c)
		curr = curr.previous
	}
	s.Pop()
	for s.Len() > 0 {
		path = append(path, s.Pop().(sector))
	}
	return path
}

func writePaths() string{
	allPathsStr := ""

	for i := 0; i < len(paths); i++ {
		pathStr := "path of car " + strconv.Itoa(i + 1) + ":\n"

		for j := 0; j < len(paths[i]); j++ {
			pathStr = pathStr + strconv.Itoa(j) +".- (" + strconv.Itoa(paths[i][j].position.xPosition) + ", " + strconv.Itoa(paths[i][j].position.yPosition) + ")\n"
		}

		allPathsStr = allPathsStr + pathStr
	}
	return allPathsStr
}
