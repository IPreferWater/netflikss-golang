package organizer

type Serie struct {
	Label          string `json:"label"`
	Seasons []Season `json:"seasons"`
}

type Season struct {
	Number int `json:"number"`
	Label string `json:"label"`
	Episodes []Episode `json:"episodes"`
}

type Episode struct {
	Number int `json:"number"`
	Label string `json:"label"`
}

type BundleMovies struct {
	Label   string `json:"label"`
	Movies []Movie `json:"movie"`
}

type Movie struct {
	Label   string `json:"label"`
}