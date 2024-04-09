package main

import (
	"time"
	"strings"
	"fmt"
)

type RedisEntry struct {
	value string
	expiry time.Duration
	canExpire bool
	expired bool
	addedAt time.Time
}

var dataStore = make(map[string]RedisEntry)

func set(args []string) {
	redisEntry := RedisEntry{
		value: args[1],
		addedAt: time.Now(),
		expired: false,
	}

	if ( len(args) > 2 && strings.ToUpper(args[2]) == EXPIRYARG ) {
		millisecondsExpiry := fmt.Sprintf("%sms", args[3])
		redisEntry.expiry, _ = time.ParseDuration(millisecondsExpiry)
		redisEntry.canExpire = true
	}

	dataStore[args[0]] = redisEntry
}

func get(key string) (string) {
	redisEntry, ok := dataStore[key]

	if (!ok || redisEntry.expired) {
		return ""
	}

	if (redisEntry.canExpire) {
		elapsedTime := time.Since(dataStore[key].addedAt)

		if elapsedTime > redisEntry.expiry {
			redisEntry.expired = true
			return ""
		}
	}

	return redisEntry.value
}