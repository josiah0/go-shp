package shp

import "testing"

func pointsEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func getShapesFromFile(prefix string, t *testing.T) (shapes []Shape) {
	filename := prefix + ".shp"
	file, err := Open(filename)
	if err != nil {
		t.Fatal("Failed to open shapefile: " + filename + " (" + err.Error() + ")")
	}
	defer file.Close()

	for file.Next() {
		_, shape := file.Shape()
		shapes = append(shapes, shape)
	}
	if file.Err() != nil {
		t.Errorf("Error while getting shapes for %s: %v", prefix, file.Err())
	}

	return shapes
}

type shapeGetterFunc func(string, *testing.T) []Shape

type identityTestFunc func(*testing.T, [][]float64, []Shape)

func test_Point(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*Point)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		if !pointsEqual([]float64{p.X, p.Y}, points[n]) {
			t.Error("Points did not match.")
		}
	}
}

func test_PolyLine(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*PolyLine)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_Polygon(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*Polygon)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_MultiPoint(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*MultiPoint)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_PointZ(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*PointZ)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		if !pointsEqual([]float64{p.X, p.Y, p.Z}, points[n]) {
			t.Error("Points did not match.")
		}
	}
}

func test_PolyLineZ(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*PolyLineZ)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y, p.ZArray[k]}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_PolygonZ(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*PolygonZ)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y, p.ZArray[k]}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_MultiPointZ(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*MultiPointZ)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y, p.ZArray[k]}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_PointM(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*PointM)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		if !pointsEqual([]float64{p.X, p.Y, p.M}, points[n]) {
			t.Error("Points did not match.")
		}
	}
}

func test_PolyLineM(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*PolyLineM)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y, p.MArray[k]}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_PolygonM(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*PolygonM)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y, p.MArray[k]}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_MultiPointM(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*MultiPointM)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y, p.MArray[k]}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_MultiPatch(t *testing.T, points [][]float64, shapes []Shape) {
	for n, s := range shapes {
		p, ok := s.(*MultiPatch)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			if !pointsEqual(points[n*3+k], []float64{point.X, point.Y, p.ZArray[k]}) {
				t.Error("Points did not match.")
			}
		}
	}
}

func test_shapeIdentity(t *testing.T, prefix string, getter shapeGetterFunc) {
	shapes := getter(prefix, t)
	d := dataForReadTests[prefix]
	if len(shapes) != d.count {
		t.Errorf("Number of shapes for %s read was wrong. Wanted %d, got %d.", prefix, d.count, len(shapes))
	}
	d.tester(t, d.points, shapes)
}

func TestReadBBox(t *testing.T) {
	tests := []struct {
		filename string
		want     Box
	}{
		{"test_files/multipatch.shp", Box{0, 0, 10, 10}},
		{"test_files/multipoint.shp", Box{0, 5, 10, 10}},
		{"test_files/multipointm.shp", Box{0, 5, 10, 10}},
		{"test_files/multipointz.shp", Box{0, 5, 10, 10}},
		{"test_files/point.shp", Box{0, 5, 10, 10}},
		{"test_files/pointm.shp", Box{0, 5, 10, 10}},
		{"test_files/pointz.shp", Box{0, 5, 10, 10}},
		{"test_files/polygon.shp", Box{0, 0, 5, 5}},
		{"test_files/polygonm.shp", Box{0, 0, 5, 5}},
		{"test_files/polygonz.shp", Box{0, 0, 5, 5}},
		{"test_files/polyline.shp", Box{0, 0, 25, 25}},
		{"test_files/polylinem.shp", Box{0, 0, 25, 25}},
		{"test_files/polylinez.shp", Box{0, 0, 25, 25}},
	}
	for _, tt := range tests {
		r, err := Open(tt.filename)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if got := r.BBox().MinX; got != tt.want.MinX {
			t.Errorf("got MinX = %v, want %v", got, tt.want.MinX)
		}
		if got := r.BBox().MinY; got != tt.want.MinY {
			t.Errorf("got MinY = %v, want %v", got, tt.want.MinY)
		}
		if got := r.BBox().MaxX; got != tt.want.MaxX {
			t.Errorf("got MaxX = %v, want %v", got, tt.want.MaxX)
		}
		if got := r.BBox().MaxY; got != tt.want.MaxY {
			t.Errorf("got MaxY = %v, want %v", got, tt.want.MaxY)
		}
	}
}

type testCaseData struct {
	points [][]float64
	tester identityTestFunc
	count  int
}

var dataForReadTests = map[string]testCaseData{
	"test_files/polygonm": testCaseData{
		points: [][]float64{
			{0, 0, 0},
			{0, 5, 5},
			{5, 5, 10},
			{5, 0, 15},
			{0, 0, 0},
		},
		tester: test_PolygonM,
		count:  1,
	},
	"test_files/multipointm": testCaseData{
		points: [][]float64{
			{10, 10, 100},
			{5, 5, 50},
			{0, 10, 75},
		},
		tester: test_MultiPointM,
		count:  1,
	},
	"test_files/multipatch": testCaseData{
		points: [][]float64{
			{0, 0, 0},
			{10, 0, 0},
			{10, 10, 0},
			{0, 10, 0},
			{0, 0, 0},
			{0, 10, 0},
			{0, 10, 10},
			{0, 0, 10},
			{0, 0, 0},
			{0, 10, 0},
			{10, 0, 0},
			{10, 0, 10},
			{10, 10, 10},
			{10, 10, 0},
			{10, 0, 0},
			{0, 0, 0},
			{0, 0, 10},
			{10, 0, 10},
			{10, 0, 0},
			{0, 0, 0},
			{10, 10, 0},
			{10, 10, 10},
			{0, 10, 10},
			{0, 10, 0},
			{10, 10, 0},
			{0, 0, 10},
			{0, 10, 10},
			{10, 10, 10},
			{10, 0, 10},
			{0, 0, 10},
		},
		tester: test_MultiPatch,
		count:  1,
	},
	"test_files/point": testCaseData{
		points: [][]float64{
			{10, 10},
			{5, 5},
			{0, 10},
		},
		tester: test_Point,
		count:  3,
	},
	"test_files/polyline": testCaseData{
		points: [][]float64{
			{0, 0},
			{5, 5},
			{10, 10},
			{15, 15},
			{20, 20},
			{25, 25},
		},
		tester: test_PolyLine,
		count:  2,
	},
	"test_files/polygon": testCaseData{
		points: [][]float64{
			{0, 0},
			{0, 5},
			{5, 5},
			{5, 0},
			{0, 0},
		},
		tester: test_Polygon,
		count:  1,
	},
	"test_files/multipoint": testCaseData{
		points: [][]float64{
			{10, 10},
			{5, 5},
			{0, 10},
		},
		tester: test_MultiPoint,
		count:  1,
	},
	"test_files/pointz": testCaseData{
		points: [][]float64{
			{10, 10, 100},
			{5, 5, 50},
			{0, 10, 75},
		},
		tester: test_PointZ,
		count:  3,
	},
	"test_files/polylinez": testCaseData{
		points: [][]float64{
			{0, 0, 0},
			{5, 5, 5},
			{10, 10, 10},
			{15, 15, 15},
			{20, 20, 20},
			{25, 25, 25},
		},
		tester: test_PolyLineZ,
		count:  2,
	},
	"test_files/polygonz": testCaseData{
		points: [][]float64{
			{0, 0, 0},
			{0, 5, 5},
			{5, 5, 10},
			{5, 0, 15},
			{0, 0, 0},
		},
		tester: test_PolygonZ,
		count:  1,
	},
	"test_files/multipointz": testCaseData{
		points: [][]float64{
			{10, 10, 100},
			{5, 5, 50},
			{0, 10, 75},
		},
		tester: test_MultiPointZ,
		count:  1,
	},
	"test_files/pointm": testCaseData{
		points: [][]float64{
			{10, 10, 100},
			{5, 5, 50},
			{0, 10, 75},
		},
		tester: test_PointM,
		count:  3,
	},
	"test_files/polylinem": testCaseData{
		points: [][]float64{
			{0, 0, 0},
			{5, 5, 5},
			{10, 10, 10},
			{15, 15, 15},
			{20, 20, 20},
			{25, 25, 25},
		},
		tester: test_PolyLineM,
		count:  2,
	},
}

func TestReadPoint(t *testing.T) {
	test_shapeIdentity(t, "test_files/point", getShapesFromFile)
}

func TestReadPolyLine(t *testing.T) {
	test_shapeIdentity(t, "test_files/polyline", getShapesFromFile)
}

func TestReadPolygon(t *testing.T) {
	test_shapeIdentity(t, "test_files/polygon", getShapesFromFile)
}

func TestReadMultiPoint(t *testing.T) {
	test_shapeIdentity(t, "test_files/multipoint", getShapesFromFile)
}

func TestReadPointZ(t *testing.T) {
	test_shapeIdentity(t, "test_files/pointz", getShapesFromFile)
}

func TestReadPolyLineZ(t *testing.T) {
	test_shapeIdentity(t, "test_files/polylinez", getShapesFromFile)
}

func TestReadPolygonZ(t *testing.T) {
	test_shapeIdentity(t, "test_files/polygonz", getShapesFromFile)
}

func TestReadMultiPointZ(t *testing.T) {
	test_shapeIdentity(t, "test_files/multipointz", getShapesFromFile)
}

func TestReadPointM(t *testing.T) {
	test_shapeIdentity(t, "test_files/pointm", getShapesFromFile)
}

func TestReadPolyLineM(t *testing.T) {
	test_shapeIdentity(t, "test_files/polylinem", getShapesFromFile)
}

func TestReadPolygonM(t *testing.T) {
	test_shapeIdentity(t, "test_files/polygonm", getShapesFromFile)
}

func TestReadMultiPointM(t *testing.T) {
	test_shapeIdentity(t, "test_files/multipointm", getShapesFromFile)
}

func TestReadMultiPatch(t *testing.T) {
	test_shapeIdentity(t, "test_files/multipatch", getShapesFromFile)
}
