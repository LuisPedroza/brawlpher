package bs

type Power struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Brawler struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	StarPowers []Power `json:"starPowers"`
	Gadgets    []Power `json:"gadgets"`
}

type Brawlers struct {
	Items []Brawler `json:"items"`
}
