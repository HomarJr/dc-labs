package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	area := 0.0
	for i := 0; i < len(points); i++ {
		x1 := points[i].X
		y1 := points[i].Y
		x2 := 0.0
		y2 := 0.0
		if i == len(points)-1 {
			x2 = points[0].X
			y2 = points[0].Y
		} else {
			x2 = points[i+1].X
			y2 = points[i+1].Y
		}
		area += ((y1 + y2) / 2) * (x2 - x1)
	}
	return area
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	perimeter := 0.0
	for i := 0; i < len(points); i++ {
		x1 := points[i].X
		y1 := points[i].Y
		x2 := 0.0
		y2 := 0.0
		if i == len(points)-1 {
			x2 = points[0].X
			y2 = points[0].Y
		} else {
			x2 = points[i+1].X
			y2 = points[i+1].Y
		}
		perimeter += math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	}
	return perimeter
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome to the Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	if len(vertices) < 3 {
		response += "ERROR - Your shape is not compliying with the minimum number of vertices.\n"
		fmt.Fprintf(w, response)
		return
	}
	response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
	response += fmt.Sprintf(" - Perimeter       : %.1f\n", perimeter)
	response += fmt.Sprintf(" - Area            : %.1f\n", area)

	// Send response to client
	fmt.Fprintf(w, response)
}
