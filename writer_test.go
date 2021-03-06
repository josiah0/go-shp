package shp

import (
	"os"
	"testing"
)

var filename_prefix string = "test_files/write_"

func removeShapefile(filename string) {
	os.Remove(filename + ".shp")
	os.Remove(filename + ".shx")
	os.Remove(filename + ".dbf")
}

func pointsToFloats(points []Point) [][]float64 {
	floats := make([][]float64, len(points))
	for k, v := range points {
		floats[k] = make([]float64, 2)
		floats[k][0] = v.X
		floats[k][1] = v.Y
	}
	return floats
}

func TestWritePoint(t *testing.T) {
	filename := filename_prefix + "point"
	defer removeShapefile(filename)

	points := [][]float64{
		{0.0, 0.0},
		{5.0, 5.0},
		{10.0, 10.0},
	}

	shape, err := Create(filename+".shp", POINT)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range points {
		shape.Write(&Point{p[0], p[1]})
	}
	shape.Close()

	shapes := getShapesFromFile(filename, t)
	if len(shapes) != len(points) {
		t.Error("Number of shapes read was wrong")
	}
	test_Point(t, points, shapes)
}

func TestWritePolyLine(t *testing.T) {
	filename := filename_prefix + "polyline"
	defer removeShapefile(filename)

	points := [][]Point{
		{Point{0.0, 0.0}, Point{5.0, 5.0}},
		{Point{10.0, 10.0}, Point{15.0, 15.0}},
	}

	shape, err := Create(filename+".shp", POLYLINE)
	if err != nil {
		t.Log(shape, err)
	}

	shape.Write(NewPolyLine(points))
	shape.Close()

	shapes := getShapesFromFile(filename, t)
	if len(shapes) != 1 {
		t.Error("Number of shapes read was wrong")
	}
	test_PolyLine(t, pointsToFloats(flatten(points)), shapes)
}
