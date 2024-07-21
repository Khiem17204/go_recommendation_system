package database

type DatabaseCardManager struct {
	*DatabaseManager
}

func NewDatabaseCardManager() (*DatabaseCardManager, error) {
	dm, err := NewDatabaseManager()
	if err != nil {
		return nil, err
	}
	return &DatabaseCardManager{dm}, nil
}
