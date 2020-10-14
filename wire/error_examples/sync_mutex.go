package error_examples

import "sync"

type Service struct {
	mutex sync.Mutex
}

