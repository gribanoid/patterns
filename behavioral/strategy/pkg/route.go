package pkg

import "fmt"

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJim := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJim
	fmt.Printf("Road A:[%d] to B:[%d] Avg speed:[%d] Traffic jim:[%d] Total:[%d] Total time:[%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJim, total, totalTime)
}
