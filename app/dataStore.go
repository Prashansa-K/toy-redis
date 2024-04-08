package main

var dataStore = make(map[string]string)

func set(key string, value string) {
	dataStore[key] = value
}

func get(key string) (string) {
	return dataStore[key]
}