package database

func GetCounter() int {
	return database.Counter
}

func IncrementCounter(n int) {
	database.Counter += n

	syncDatabase()
}

func DecrementCounter(n int) {
	database.Counter -= n

	syncDatabase()
}

func SetCounter(n int) {
	database.Counter = n

	syncDatabase()
}
