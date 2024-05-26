package hash_maps

type RequestLogger struct {
	timeLimit      int
	lastRequestMap map[string]int
}

func (rl *RequestLogger) requestLoggerInit(timeLimit int) {
	rl.timeLimit = timeLimit
	rl.lastRequestMap = make(map[string]int)
}

func (rl *RequestLogger) messageRequestDecision(timestamp int, request string) bool {
	lastTimestamp, ok := rl.lastRequestMap[request]
	if ok {
		if timestamp-lastTimestamp >= rl.timeLimit {
			rl.lastRequestMap[request] = timestamp
			return true
		} else {
			return false
		}
	}

	rl.lastRequestMap[request] = timestamp
	return true
}
