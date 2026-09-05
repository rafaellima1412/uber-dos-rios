package dto

type CabinLayout struct {
	Andar         int    `json:"andar"`
	Lado          string `json:"lado"`
	Local         string `json:"local"`
	Identificador string `json:"identificador"`
}

type Cabin struct {
	Name        string  `json:"name"`
	BedType     string  `json:"bed_type"`
	Capacity    int     `json:"capacity"`
	Description *string `json:"description"`
}

type CreateCabinRequest struct {
	Cabins []Cabin `json:"cabins"`
}

type CabinDetailResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	BedType  string `json:"bed_type"`
	Capacity int    `json:"capacity"`

	Description *string `json:"description"`
}

type UpdateCabinRequest struct {
	Name        string  `json:"name"`
	BedType     string  `json:"bed_type"`
	Capacity    int     `json:"capacity"`
	Description *string `json:"description"`
}

type CabinListItemResponse struct {
	Cabins      []CabinDetailResponse `json:"cabins"`
	TotalCount  int                   `json:"total_count"`
	TotalPages  int                   `json:"total_pages"`
	CurrentPage int                   `json:"current_page"`
}
