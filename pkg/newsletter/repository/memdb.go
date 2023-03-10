package repository

import (
	"github.com/Eitol/newsletter-api/pkg/newsletter"
	"sync"
)

var inMemoryDB []newsletter.Subscription
var resultIdxPool []int
var mutex = &sync.Mutex{}
