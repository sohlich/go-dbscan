package dbscan

const (
	NOISE     = false
	CLUSTERED = true
)

type Clusterable interface {
	Distance(c interface{}) float64
}

type Cluster []Clusterable

func Clusterize(objects []Clusterable, minPts int, eps float64) []Cluster {
	clusters := make([]Cluster, 0)
	visited := map[Clusterable]bool{}
	for _, point := range objects {
		neighbours := findNeighbours(point, objects, eps)
		if len(neighbours) >= minPts {
			visited[point] = CLUSTERED
			cluster := make(Cluster, 1)
			cluster[0] = point
			cluster = expandCluster(cluster, neighbours, visited, minPts, eps)
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
func findNeighbours(point Clusterable, points []Clusterable, eps float64) []Clusterable {
	neighbours := make([]Clusterable, 0)
	for _, potNeigb := range points {
		if point != potNeigb && potNeigb.Distance(point) <= eps {
			neighbours = append(neighbours, potNeigb)
		}
	}
	return neighbours
}

//Try to expand existing clutser
func expandCluster(cluster Cluster, neighbours []Clusterable, visited map[Clusterable]bool, minPts int, eps float64) Cluster {
	seed := make([]Clusterable, len(neighbours))
	copy(seed, neighbours)
	for _, point := range seed {
		pointState, isVisited := visited[point]
		if !isVisited {
			currentNeighbours := findNeighbours(point, seed, eps)
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
	mergeMap := make(map[Clusterable]bool)
	putAll(mergeMap, one)
	putAll(mergeMap, two)
	merged := make([]Clusterable, 0)
	for key := range mergeMap {
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
