package db

const SetTaskPrioritiesTableValuesQuery = `INSERT INTO task_priorities (priority)
										 	VALUES ('Urgent'),
													('High'),
													('Medium'),
													('Low'),
													('Backlog');`

func SetSeeds() error {
	_, err := GetDBConnection().Exec(SetTaskPrioritiesTableValuesQuery)
	if err != nil {
		return err
	}
	return nil
}
