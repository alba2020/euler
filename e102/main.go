package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Point3D struct {
	x float64
	y float64
	z float64
}

type Vector3D struct {
	x float64
	y float64
	z float64
}

func (src Point3D) To(dest Point3D) Vector3D {
	return Vector3D{
		x: dest.x - src.x,
		y: dest.y - src.y,
		z: dest.z - src.z,
	}
}

func (a Vector3D) CrossProduct(b Vector3D) Vector3D {
	i := a.y*b.z - a.z*b.y
	j := -(a.x*b.z - a.z*b.x)
	k := a.x*b.y - a.y*b.x
	return Vector3D{i, j, k}
}

func (v Vector3D) Module() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

type Triangle struct {
	a Point3D
	b Point3D
	c Point3D
}

func (t Triangle) Area() float64 {
	v1 := t.a.To(t.b)
	v2 := t.a.To(t.c)
	return v1.CrossProduct(v2).Module() / 2
}

func (t Triangle) Contains(p Point3D) bool {
	s1 := Triangle{t.a, t.b, p}.Area()
	s2 := Triangle{t.b, t.c, p}.Area()
	s3 := Triangle{t.c, t.a, p}.Area()
	return s1+s2+s3 <= t.Area()
}

func NewTriangleFromSlice(coords []string) Triangle {
	cs := make([]float64, 6)
	for i, c := range coords {
		cs[i], _ = strconv.ParseFloat(c, 64)
	}
	return Triangle{Point3D{cs[0], cs[1], 0}, Point3D{cs[2], cs[3], 0}, Point3D{cs[4], cs[5], 0}}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	// t1 := Triangle{Point3D{1, 2, 0}, Point3D{5, 2, 0}, Point3D{-5, -5, 0}}
	// fmt.Println(t1.Contains(Point3D{1, 10, 0}))
	records := readCsvFile("triangles.txt")
	found := 0
	for _, rec := range records {
		t := NewTriangleFromSlice(rec)
		// fmt.Println(t, t.Area())
		if t.Contains(Point3D{0, 0, 0}) {
			found++
		}
	}
	fmt.Println("found=", found)
}
