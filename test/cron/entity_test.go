package cron_test

import (
	"container/heap"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tsubasa597/BILIBILI-HELPER/cron"
)

var (
	_now = time.Now()

	_entities = []*cron.Entity{
		{
			Task: cron.NewDynamic(1, 1, 4, nil),
			ID:   4,
			Prev: _now,
		},
		{
			Task: cron.NewDynamic(1, 1, 3, nil),
			ID:   3,
			Prev: _now.Add(time.Second * 1),
		},
		{
			Task: cron.NewDynamic(1, 1, 2, nil),
			ID:   2,
			Prev: _now.Add(time.Second * 2),
		},
		{
			Task: cron.NewDynamic(1, 1, 1, nil),
			ID:   1,
			Prev: _now.Add(time.Second * 3),
		},
	}
)

func TestEnties(t *testing.T) {
	assert := assert.New(t)

	var entities cron.Entities
	for _, entity := range _entities {
		heap.Push(&entities, entity)
	}

	entity := heap.Pop(&entities).(*cron.Entity)
	assert.Equal(int64(4), entity.ID)
	entity.Done()
	heap.Pop(&entities).(*cron.Entity).Done()
	heap.Pop(&entities).(*cron.Entity).Done()
	heap.Pop(&entities).(*cron.Entity).Done()

	entities.Reset()
	entity = heap.Pop(&entities).(*cron.Entity)
	assert.Equal(int64(1), entity.ID)
	entity.Done()
	heap.Pop(&entities).(*cron.Entity).Done()
	heap.Pop(&entities).(*cron.Entity).Done()
	heap.Pop(&entities).(*cron.Entity).Done()

	entities.Reset()
	heap.Pop(&entities)
	entities.Remove()

	heap.Pop(&entities).(*cron.Entity).Done()
	heap.Pop(&entities).(*cron.Entity).Done()
	heap.Pop(&entities)
	entities.Remove()

	entities.Reset()
	assert.Equal(int64(2), heap.Pop(&entities).(*cron.Entity).ID)
	assert.Equal(int64(3), heap.Pop(&entities).(*cron.Entity).ID)
}
