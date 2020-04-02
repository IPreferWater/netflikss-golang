package organizer


type Episode struct {
	Number int `json:"number"`
	Label string `json:"label"`
}

type Season struct {
	Number int `json:"number"`
	Label string `json:"label"`
	Episodes []Episode `json:"episodes"`
}

type Info struct {
	Label          string `json:"label"`
	Seasons []Season `json:"seasons"`
}