package dao

func (db *DataBase) CreateFieldTable() error {

	if err := db.createTable(gblog_contents); err != nil {
		return err
	}
	return nil
}
