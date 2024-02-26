package usecase

type csvReader interface {
	//SetData(data []string) error
	SetData(data []byte) error
	GetContent(path string) ([]byte, error)
}

type EntityData interface {
	SetCountry(new string)
	GetCountry() string
	GetProvider() string
}
