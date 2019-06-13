package models

type NudityResponse struct {
	Output NudityOutputResponse `json:"output"`
	ID     string               `json:"id"`
}

type NudityOutputResponse struct {
	Detections []NudityDetection `json:"detections"`
	NSFWScore  float64         `json:"nsfw_score"`
}

type NudityDetection struct {
	Confidence  string `json:"confidence"`
	BoundingBox []int  `json:"bounding_box"`
	Name        string `json:"name"`
}
