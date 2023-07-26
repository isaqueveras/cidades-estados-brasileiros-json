package script

// Lista modela os dados das cidades e estados
type Lista struct {
	TotalCidades int64    `json:"total_cidades,omitempty"`
	Estados      []Estado `json:"estados,omitempty"`
}

// Estado modela os dados para cada estado
type Estado struct {
	Sigla   string   `json:"sigla,omitempty"`
	Nome    string   `json:"nome,omitempty"`
	Cidades []string `json:"cidades,omitempty"`
}

// NovaLista modela os dados das cidades e estados
type NovaLista struct {
	TotalCidades int64                 `json:"total_cidades"`
	Estados      []NovaEstruturaEstado `json:"estados,omitempty"`
}

// NovaEstruturaEstado modela os dados para uma nova estrutura de cidades por estado
type NovaEstruturaEstado struct {
	ID      string   `json:"id"`
	Sigla   string   `json:"sigla,omitempty"`
	Nome    string   `json:"nome,omitempty"`
	Cidades []Cidade `json:"cidades,omitempty"`
}

// Cidade modela os dados para cada cidade
type Cidade struct {
	ID       string `json:"id"`
	EstadoID string `json:"estado_id"`
	Nome     string `json:"nome,omitempty"`
	Slug     string `json:"slug,omitempty"`
}
