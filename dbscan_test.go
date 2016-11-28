package dbscan

import (
	"fmt"
	"log"
	"math"
	"testing"
)

type SimpleClusterable struct {
	position float64
}

func (s SimpleClusterable) Distance(c interface{}) float64 {
	distance := math.Abs(c.(SimpleClusterable).position - s.position)
	return distance
}

func (s SimpleClusterable) GetID() string {
	return fmt.Sprint(s.position)
}

func TestPutAll(t *testing.T) {
	testMap := make(map[string]Clusterable)
	clusterList := []Clusterable{
		SimpleClusterable{10},
		SimpleClusterable{12},
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
		SimpleClusterable{0},
		SimpleClusterable{1},
		SimpleClusterable{-1},
		SimpleClusterable{1.5},
		SimpleClusterable{-0.5},
	}

	eps := 1.0
	neighbours := findNeighbours(clusterList[0], clusterList, eps)

	assertEquals(t, 3, len(neighbours))
}

func TestMerge(t *testing.T) {
	log.Println("Executing TestMerge")
	expected := 6
	one := []Clusterable{
		SimpleClusterable{0},
		SimpleClusterable{1},
		SimpleClusterable{2.1},
		SimpleClusterable{2.2},
		SimpleClusterable{2.3},
	}

	two := []Clusterable{
		one[0],
		one[1],
		SimpleClusterable{2.4},
	}

	output := merge(one, two)
	assertEquals(t, expected, len(output))
}

func TestExpandCluster(t *testing.T) {
	log.Println("Executing TestExpandCluster")
	expected := 4
	clusterList := []Clusterable{
		SimpleClusterable{0},
		SimpleClusterable{1},
		SimpleClusterable{2},
		SimpleClusterable{2.1},
		SimpleClusterable{5},
	}

	eps := 1.0
	minPts := 3
	visitMap := make(map[string]bool)
	cluster := make(Cluster, 0)
	cluster = expandCluster(cluster, clusterList, visitMap, minPts, eps)
	assertEquals(t, expected, len(cluster))
}

func TestClusterize(t *testing.T) {
	log.Println("Executing TestClusterize")
	clusterList := []Clusterable{
		SimpleClusterable{1},
		SimpleClusterable{0.5},
		SimpleClusterable{0},
		SimpleClusterable{5},
		SimpleClusterable{4.5},
		SimpleClusterable{4},
	}
	eps := 1.0
	minPts := 2
	clusters := Clusterize(clusterList, minPts, eps)
	assertEquals(t, 2, len(clusters))
	if 2 == len(clusters) {
		assertEquals(t, 3, len(clusters[0]))
		assertEquals(t, 3, len(clusters[1]))
	}
}

func TestClusterizeNoData(t *testing.T) {
	log.Println("Executing TestClusterizeNoData")
	clusterList := []Clusterable{}
	eps := 1.0
	minPts := 3
	clusters := Clusterize(clusterList, minPts, eps)
	assertEquals(t, 0, len(clusters))
}

//Assert function. If  the expected value not equals result, function
//returns error.
func assertEquals(t *testing.T, expected, result int) {
	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
