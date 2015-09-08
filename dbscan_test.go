package dbscan

import (
	"log"
	"math"
	"testing"
)

type SimpleClusterable struct {
	position float64
}

func (s *SimpleClusterable) Distance(c Clusterable) float64 {
	distance := math.Abs(c.GetParams()[0] - s.GetParams()[0])
	return distance
}

func (s *SimpleClusterable) GetParams() []float64 {
	return []float64{s.position}
}

func TestPutAll(t *testing.T) {
	testMap := make(map[Clusterable]bool)
	clusterList := []Clusterable{
		&SimpleClusterable{10},
		&SimpleClusterable{12},
	}
	putAll(testMap, clusterList)
	mapSize := len(testMap)
	if mapSize != 2 {
		t.Errorf("Map does not contain expected size 2 but was %d", mapSize)
	}
}

//Test find neighbour function
func TestFindNeighbours(t *testing.T) {
	log.Println("Executing TestFindNeighbours")
	clusterList := []Clusterable{
		&SimpleClusterable{0},
		&SimpleClusterable{1},
		&SimpleClusterable{2},
		&SimpleClusterable{2},
		&SimpleClusterable{2},
	}

	Eps = 1
	neighbours := findNeighbours(clusterList[0], clusterList)

	assertEquals(t, 1, len(neighbours))
}

func TestMerge(t *testing.T) {
	log.Println("Executing TestMerge")
	expected := 6
	one := []Clusterable{
		&SimpleClusterable{0},
		&SimpleClusterable{1},
		&SimpleClusterable{2},
		&SimpleClusterable{2},
		&SimpleClusterable{2},
	}

	two := []Clusterable{
		one[0],
		one[1],
		&SimpleClusterable{2},
	}

	output := merge(one, two)
	assertEquals(t, expected, len(output))
}

func TestExpandCluster(t *testing.T) {
	log.Println("Executing TestExpandCluster")
	expected := 3
	clusterList := []Clusterable{
		&SimpleClusterable{0},
		&SimpleClusterable{1},
		&SimpleClusterable{2},
		&SimpleClusterable{2},
		&SimpleClusterable{5},
	}

	Eps = 1
	visitMap := make(map[Clusterable]bool)
	cluster := make(Cluster, 0)
	cluster = expandCluster(cluster, clusterList, visitMap)
	assertEquals(t, expected, len(cluster))
}

func TestClusterize(t *testing.T) {
	log.Println("Executing TestExpandCluster")
	clusterList := []Clusterable{
		&SimpleClusterable{1},
		&SimpleClusterable{2},
		&SimpleClusterable{2},
		&SimpleClusterable{5},
		&SimpleClusterable{6},
		&SimpleClusterable{4},
		&SimpleClusterable{5},
	}
	clusters := Clusterize(clusterList)
	assertEquals(t, 2, len(clusters))
}

//Assert function. If  the expected value not equals result, function
//returns error.
func assertEquals(t *testing.T, expected, result int) {
	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
