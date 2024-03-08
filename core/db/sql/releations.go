package mysqldb

///releations///

// CreateTables creates both users and addresses tables and establishes a relationship between them
func CreateTables() error {
	err := CreateUserTable()
	if err != nil {
		return err
	}

	err = CreateAddressTable()
	if err != nil {
		return err
	}

	err = AddRelation("users", "id", "addresses", "user_id")
	if err != nil {
		return err
	}

	return nil
}
