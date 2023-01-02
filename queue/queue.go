package queue

import (
	"errors"
	"sync"

	"github.com/biganashvili/imdb-challenge/models"
)

type Queue struct {
	mu sync.Mutex
	q  []models.Movie
}

// FifoQueue
type FifoQueue interface {
	Insert()
	Remove()
}

// Insert inserts the item into the queue
func (q *Queue) Insert(item models.Movie) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q = append(q.q, item)
}

// Remove removes the oldest element from the queue
func (q *Queue) Remove() (models.Movie, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		item := q.q[0]
		q.q = q.q[1:]
		return item, nil
	}
	return models.Movie{}, errors.New("Queue is empty")
}

// Remove removes the oldest element from the queue
func (q *Queue) List() []models.Movie {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.q
}

// CreateQueue creates an empty queue with desired capacity
func CreateQueue() *Queue {
	return &Queue{
		q: make([]models.Movie, 0),
	}
}
