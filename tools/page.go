package tools

import "github.com/tyr-tech-team/hawk/env"

// ConvertMongoSizeAndLimit -
func ConvertMongoSizeAndLimit(page, size int64) (skip int64, limit int64) {
	if page < 1 {
		page = 1
	}

	if size == 0 {
		size = env.DefaultPageSize
	}

	skip = size * (page - 1)
	limit = size

	return
}
