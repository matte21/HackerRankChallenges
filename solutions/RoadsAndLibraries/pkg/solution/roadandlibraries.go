package solution

// Complete the RoadAndLibraries function below
// Time complexity: o(1), O(n*len(roads)).
// Space complexity: o(1), O(n).
func RoadsAndLibraries(n, cLib, cRoad int32, roads [][]int32) int64 {
	// if building a library is cheaper than building a street the best solution
	// is to build a library in every city.
	if cLib <= cRoad {
		return int64(n) * int64(cLib)
	}

	// repToIsland associates multi-city island representatives to their island.
	// A multi-city island representative is the first city belonging to that
	// island to be found. A connected sub-graph of the input graph made up by
	// more than one city which is disconnected by the rest of the cities of the
	// input graph is called an island.
	repToIsland := make(map[int32]map[int32]struct{})

	// nodeToIslandRep associates nodes belonging to multi-node islands to the
	// representative of such island.
	nodeToIslandRep := make(map[int32]int32)

	for _, aRoad := range roads {
		rep1, rep1Found := nodeToIslandRep[aRoad[0]]
		rep2, rep2Found := nodeToIslandRep[aRoad[1]]
		switch {
		case rep1Found && rep2Found && (rep1 != rep2):
			mergeIslands(rep1, rep2, nodeToIslandRep, repToIsland)
		case rep1Found && !rep2Found:
			addToIsland(rep1, aRoad[1], nodeToIslandRep, repToIsland)
		case !rep1Found && rep2Found:
			addToIsland(rep2, aRoad[0], nodeToIslandRep, repToIsland)
		case !rep1Found && !rep2Found:
			newIsland(aRoad[0], aRoad[1], nodeToIslandRep, repToIsland)
		}
	}

	roadsCost := int64(0)
	singleCityIslands := n
	for _, anIsland := range repToIsland {
		singleCityIslands -= int32(len(anIsland)) + int32(1)
		roadsCost += int64(len(anIsland)) * int64(cRoad)
	}

	return roadsCost + int64(cLib)*(int64(len(repToIsland))+int64(singleCityIslands))
}

func mergeIslands(rep1, rep2 int32,
	nodeToIslandRep map[int32]int32,
	repToIsland map[int32]map[int32]struct{}) {

	// Merge the smallest island to the biggest one because it's more efficient.
	if len(repToIsland[rep1]) >= len(repToIsland[rep2]) {
		mergeIslandsInOrder(rep1, rep2, nodeToIslandRep, repToIsland)
	} else {
		mergeIslandsInOrder(rep2, rep1, nodeToIslandRep, repToIsland)
	}
}

func addToIsland(islandRep, toAdd int32,
	nodeToIslandRep map[int32]int32,
	repToIsland map[int32]map[int32]struct{}) {

	repToIsland[islandRep][toAdd] = struct{}{}
	nodeToIslandRep[toAdd] = islandRep
}

func newIsland(city1, city2 int32,
	nodeToIslandRep map[int32]int32,
	repToIsland map[int32]map[int32]struct{}) {

	nodeToIslandRep[city1] = city1
	nodeToIslandRep[city2] = city1
	repToIsland[city1] = map[int32]struct{}{city2: struct{}{}}
}

func mergeIslandsInOrder(rep1, rep2 int32,
	nodeToIslandRep map[int32]int32,
	repToIsland map[int32]map[int32]struct{}) {

	island1, island2 := repToIsland[rep1], repToIsland[rep2]
	for aNodeInIsland2 := range island2 {
		island1[aNodeInIsland2] = struct{}{}
		nodeToIslandRep[aNodeInIsland2] = rep1
	}
	island1[rep2] = struct{}{}
	nodeToIslandRep[rep2] = rep1
	delete(repToIsland, rep2)
}
