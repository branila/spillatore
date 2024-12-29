package database

import "github.com/branila/spillatore/types"

func GetStats() []types.UserStats {
	return database.Stats
}

func AddStat(stat types.UserStats) {
	database.Stats = append(database.Stats, stat)
}
