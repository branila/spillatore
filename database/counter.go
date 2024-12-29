package database

func GetCounter() int {
	return database.Counter
}

func IncrementCounter(n int) {
	database.Counter += n

	syncDatabase()
}
