package dto

type ResponseDto struct {
	NumFound int      `json:"numFound"`
	Docs     HotelsDto `json:"docs"`
}

type SolrResponseDto struct {
	Response ResponseDto `json:"response"`
}

type DocDto struct {
	Doc HotelDto `json:"doc"`
}

type AddDto struct {
	Add DocDto `json:"add"`
}

type DeleteDoc struct {
	Query string `json:"query"`
}

type DeleteDto struct {
	Delete DeleteDoc `json:"delete"`
}

type SolrResponsesDto []SolrResponseDto