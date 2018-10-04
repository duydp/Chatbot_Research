package fptai

type Meaning struct {
	Intents []struct{
		Name string `json:"label"`
		Confidence float64
	}
	Entities []struct{
		Name string `json:"type"`
		Value string `json:"value"`
		Confidence float64 `json:"confidence"`
	}
}