package ngos

// Ngos struct
type Ngos struct {
	Reader Reader
	Writer Writer
}

// New function return *Ngos
func New() *Ngos {
	csvReader := CSVReader{}
	csvWriter := CSVWriter{}
	return &Ngos{
		Reader: csvReader,
		Writer: csvWriter,
	}
}

// Run function, will Run Ngos
func (n *Ngos) Run() {

}

func (n *Ngos) compare(a, b []string) [][]string {
	for i := len(a) - 1; i >= 0; i-- {
		for _, vD := range b {
			if a[i] == vD {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}

	var result [][]string

	for _, v := range a {
		result = append(result, []string{v})
	}

	return result
}
