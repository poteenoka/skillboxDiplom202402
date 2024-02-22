package usecase

type csvReader interface {
	//SetData(data []string) error
	SetData(data []byte) error
	GetContent(path string) ([]byte, error)
}

//type getFileContent interface {
//	getFileContent(path string) ([]string, error)
//}
