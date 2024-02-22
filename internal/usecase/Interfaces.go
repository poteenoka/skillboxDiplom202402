package usecase

type csvReader interface {
	//SetData(data []string) error
	SetData(data interface{}) error
	GetContent(path string) ([]string, error)
}

//type getFileContent interface {
//	getFileContent(path string) ([]string, error)
//}
