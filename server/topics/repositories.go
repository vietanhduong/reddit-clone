package topics

import (
	"reddit-clone/server/common"
)

type (
	Repository struct {
		// Implement your database in here
		// In this project I just use seed data
		topics PriorityQueue
	}
)

func NewRepository() *Repository {
	topics := New()
	return &Repository{topics: topics}
}

func (r *Repository) FindByID(id int) *Topic {
	if item, found := r.topics.lookup[id]; found {
		return item.topic
	}
	return nil
}

func (r *Repository) First10() []*Topic {
	var topics []*Topic
	size := common.Min(r.topics.Len(), 10)
	for size > 0 {
		topic := r.topics.Pop()
		topics = append(topics, topic)
		size--
	}
	for _, t := range topics {
		r.topics.Insert(t, t.UpVote-t.DownVote)
	}
	return topics
}

func (r *Repository) Insert(topic *Topic) *Topic {
	topic.ID = len(r.topics.lookup) + 1
	r.topics.Insert(topic, 0)
	return topic
}

func (r *Repository) UpVote(id int) *Vote {
	item, found := r.topics.lookup[id]
	if !found {
		return nil
	}
	item.topic.UpVote++
	return r.updateVote(item)
}

func (r *Repository) DownVote(id int) *Vote {
	item, found := r.topics.lookup[id]
	if !found {
		return nil
	}
	item.topic.DownVote++
	return r.updateVote(item)
}

func (r *Repository) updateVote(i *item) *Vote {
	r.topics.UpdatePriority(i.topic.ID, i.topic.UpVote-i.topic.DownVote)
	return &Vote{
		ID:       i.topic.ID,
		UpVote:   i.topic.UpVote,
		DownVote: i.topic.DownVote,
	}
}
