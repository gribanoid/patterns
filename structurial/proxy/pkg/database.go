package pkg

type Database struct {
}

func (db *Database) GetData(user string) ([]string, error) {
	return []string{"Строка 1", "Последняя строка"}, nil
}
