package topics

import (
	"math"
	"sort"

	"github.com/vietanhduong/reddit-clone/server/common"
)

type (
	Repository struct {
		// Implement your database in here
		// In this project I just use seed data
		topics map[int]*Topic
	}
)

func NewRepository() *Repository {
	return &Repository{
		topics: make(map[int]*Topic),
	}
}

func (r *Repository) FindByID(id int) *Topic {
	if topic, found := r.topics[id]; found {
		return topic
	}
	return nil
}

func (r *Repository) First10() []*Topic {
	var topics []*Topic
	size := len(r.topics)
	pairs := make(common.Pairs, size)

	i := 0
	for k, v := range r.topics {
		pairs[i] = common.Pair{Key: k, Value: v.UpVote - v.DownVote}
		i++
	}
	sort.Sort(pairs)
	for _, p := range pairs[:int(math.Min(10, float64(size)))] {
		topics = append(topics, r.topics[p.Key])
	}
	return topics
}

func (r *Repository) Insert(topic *Topic) *Topic {
	topic.ID = len(r.topics) + 1
	r.topics[topic.ID] = topic
	return topic
}

func (r *Repository) UpVote(id int) bool {
	topic, found := r.topics[id]
	if !found {
		return false
	}
	topic.UpVote++
	return true
}

func (r *Repository) DownVote(id int) bool {
	topic, found := r.topics[id]
	if !found {
		return false
	}
	topic.DownVote++
	return true
}
