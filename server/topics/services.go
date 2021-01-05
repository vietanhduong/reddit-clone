package topics

import (
	"fmt"
	"net/http"

	"reddit-clone/server/common"
)

type (
	TopicRepository interface {
		FindByID(id int) *Topic
		First10() []*Topic
		Insert(topic *Topic) *Topic
		UpVote(id int) *Vote
		DownVote(id int) *Vote
	}
	ServiceImpl struct {
		topicRepo TopicRepository
	}
)

func NewServiceImpl(topicRepo TopicRepository) *ServiceImpl {
	return &ServiceImpl{topicRepo: topicRepo}
}

func (s *ServiceImpl) GetTopicById(id int) (*Topic, error) {
	topic := s.topicRepo.FindByID(id)
	if common.IsNil(topic) {
		return nil, common.HttpError(http.StatusNotFound, "topic does not exist")
	}
	return topic, nil
}

func (s *ServiceImpl) CreateTopic(topic *Topic) (*Topic, error) {
	if common.IsEmpty(topic.Content) {
		return nil, common.HttpError(http.StatusBadRequest, "topic content cannot be empty")
	}
	if len(topic.Content) > 255 {
		return nil, common.HttpError(http.StatusBadRequest, "topic content cannot be exceed 255 characters")
	}
	topic = s.topicRepo.Insert(topic)
	return topic, nil
}

func (s *ServiceImpl) Fetch() []*Topic {
	return s.topicRepo.First10()
}

func (s *ServiceImpl) Vote(id int, up bool) (*Vote, error) {
	var vote *Vote
	v := "down"
	if up {
		vote = s.topicRepo.UpVote(id)
		v = "up"
	} else {
		vote = s.topicRepo.DownVote(id)
	}
	if common.IsNil(vote) {
		return nil, common.HttpError(http.StatusBadRequest, fmt.Sprintf("%svote failed maybe topic does not exist", v))
	}
	return vote, nil
}
