package script

// Lista modela os dados das cidades e estados
type Lista struct {
	Estados []Estado
}

// Estado modela os dados para cada estado
type Estado struct {
	Sigla   string   `json:"sigla"`
	Nome    string   `json:"nome"`
	Cidades []string `json:"cidades"`
}

// NovaEstruturaEstado modela os dados para uma nova estrutura de cidades por estado
type NovaEstruturaEstado struct {
	Sigla   string   `json:"sigla"`
	Nome    string   `json:"nome"`
	Cidades []Cidade `json:"cidades"`
}

// Cidade modela os dados para cada cidade
type Cidade struct {
	Nome string `json:"nome"`
	Slug string `json:"slug"`
}
