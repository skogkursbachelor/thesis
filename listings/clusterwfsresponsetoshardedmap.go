func (wfsResponse WFSResponse) ClusterWFSResponseToShardedMap() *ShardedMap {
	featureMap := NewShardedMap(runtime.NumCPU())
	semaphore := make(chan struct{}, runtime.NumCPU())
	var wg sync.WaitGroup

	for _, feature := range wfsResponse.Features {
		semaphore <- struct{}{}
		wg.Add(1)

		go func(feature ForestRoad) {
			defer wg.Done()
			defer func() { <-semaphore }()
			
			middleIndex := len(feature.Geometry.Coordinates) / 2
			coordinates := feature.Geometry.Coordinates[middleIndex]
			
			roundedX := utils.RoundToNearest500(int(math.Round(coordinates[0])))
			roundedY := utils.RoundToNearest500(int(math.Round(coordinates[1])))
			roundedCoordinates := fmt.Sprintf("%d,%d", roundedX, roundedY)
			
			featureMap.Set(roundedCoordinates, feature)
		}(feature)
	}

	wg.Wait()
	return featureMap
}