package dbscan

import (
	"log"
)

var Eps = 0.2
var minPts = 3

const (
	NOISE     = false
	CLUSTERED = true
)

type Clusterable interface {
	Distance(c Clusterable) float64
	GetParams() []float64
}

type Cluster []Clusterable

func Clusterize(objects []Clusterable) []Cluster {
	clusters := make([]Cluster, 0)
	visited := map[Clusterable]bool{}
	for _, point := range objects {
		neighbours := findNeighbours(point, objects)
		if len(neighbours) >= minPts {
			visited[point] = CLUSTERED
			cluster := make(Cluster, 1)
			cluster[0] = point
			cluster = expandCluster(cluster, neighbours, visited)
			clusters = append(clusters, cluster)
		} else {
			visited[point] = NOISE
		}
	}
	return clusters
}

//Finds the neighbours from given array
//depends on Eps variable, which determines
//the distance limit from the point
func findNeighbours(point Clusterable, points []Clusterable) []Clusterable {
	neighbours := make([]Clusterable, 0)
	for _, potNeigb := range points {
		if point != potNeigb && potNeigb.Distance(point) <= Eps {
			neighbours = append(neighbours, potNeigb)
		}
	}
	return neighbours
}

//Try to expand existing clutser
func expandCluster(cluster Cluster, neighbours []Clusterable, visited map[Clusterable]bool) Cluster {
	seed := make([]Clusterable, len(neighbours))
	copy(seed, neighbours)
	for _, point := range seed {
		pointState, isVisited := visited[point]
		if !isVisited {
			currentNeighbours := findNeighbours(point, seed)
			if len(currentNeighbours) >= minPts {
				cluster = merge(cluster, currentNeighbours)
			}
		}

		if isVisited && pointState == NOISE {
			visited[point] = CLUSTERED
			cluster = append(cluster, point)
		}
	}

	return cluster
}

func merge(one []Clusterable, two []Clusterable) []Clusterable {
	merMap := make(map[Clusterable]bool)
	putAll(merMap, one)
	putAll(merMap, two)

	merged := make([]Clusterable, 0)
	for key := range merMap {
		merged = append(merged, key)
	}

	return merged
}

//Function to add all values from list to map
//map keys is then the unique collecton from list
func putAll(m map[Clusterable]bool, list []Clusterable) {
	for _, val := range list {
		m[val] = true
	}
}
