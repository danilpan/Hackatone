package model

type Establishment struct {
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	Type         string   `json:"type"`
	AverageCheck int      `json:"averageCheck"`
	Rating       int      `json:"rating"`
	ImagesURLs   []string `json:"imagesURLs"`
	Tables       []int    `json:"tables,omitempty"`
}
