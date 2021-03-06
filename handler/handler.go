package handler

func OK() map[string]interface{} {
	return map[string]interface{}{
		"message": "OK",
	}
}

func Err(err error) map[string]interface{} {
	return map[string]interface{}{
		"message": err.Error(),
	}
}