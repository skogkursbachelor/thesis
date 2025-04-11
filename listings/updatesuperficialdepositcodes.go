func UpdateSuperficialDepositCodes(roads *[]models.ForestRoad) error {
	semaphore := make(chan struct{}, runtime.NumCPU())
	var wg sync.WaitGroup

	for i := 0; i < len(*roads); i++ {
		wg.Add(1)

		semaphore <- struct{}{}

		go func(road *models.ForestRoad) {
			defer wg.Done()
			defer func() { <-semaphore }()

			codes, err := getSuperficialDepositCodesForRoad(*road)
			if err != nil {
				log.Warn().Msg("Failed to get superficial deposit codes: " + err.Error())
				return
			}

			road.SuperficialDepositCodes = codes
		}(&(*roads)[i])
	}

	wg.Wait()
	return nil
}