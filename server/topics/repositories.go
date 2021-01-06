package topics

import (
	"reddit-clone/server/common"
)

type (
	Repository struct {
		// Implement your database in here
		// In this project I just use seed data
		topics map[int]*Topic
		votes  PriorityQueue
	}
)

func NewRepository() *Repository {
	votes := New()
	return &Repository{
		topics: make(map[int]*Topic),
		votes:  votes,
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
	size := common.Min(r.votes.Len(), 10)
	for _, v := range (*r.votes.itemHeap)[:size] {
		topics = append(topics, r.topics[v.value.ID])
	}
	return topics
}

func (r *Repository) Insert(topic *Topic) *Topic {
	topic.ID = len(r.topics)
	r.topics[topic.ID] = topic
	vote := Vote{
		ID: topic.ID,
	}
	r.votes.Insert(vote, 0)
	return topic
}

func (r *Repository) UpVote(id int) *Vote {
	topic, found := r.topics[id]
	if !found {
		return nil
	}
	topic.UpVote++
	return r.updateVote(topic)
}

func (r *Repository) DownVote(id int) *Vote {
	topic, found := r.topics[id]
	if !found {
		return nil
	}
	topic.DownVote++
	return r.updateVote(topic)
}

func (r *Repository) updateVote(topic *Topic) *Vote {
	r.votes.UpdatePriority(topic.ID, topic.UpVote-topic.DownVote)
	return &Vote{
		ID:       topic.ID,
		UpVote:   topic.UpVote,
		DownVote: topic.DownVote,
	}
}
