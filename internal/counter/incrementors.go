package counter

func PushInRequestsCounter() {
	inRequests.Add()
}

func PushOutRequestsCounter() {
	outRequests.Add()
}

func PushSuccessRequestsCounter() {
	successRequests.Add()
}

func PushFailedRequestsCounter() {
	failedRequests.Add()
}

func PushHitRedisCounter() {
	hitRedis.Add()
}

func PushMissRedisCounter() {
	missRedis.Add()
}
