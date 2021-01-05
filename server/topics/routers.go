package topics

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/vietanhduong/reddit-clone/server/common"
	"github.com/vietanhduong/reddit-clone/server/handler"
)

type (
	TopicService interface {
		GetTopicById(id int) (*Topic, error)
		CreateTopic(topic *Topic) (*Topic, error)
		Fetch() []*Topic
		Vote(id int, up bool) error
	}

	ServerImpl struct {
		topicSrv TopicService
	}
)

func RegisterAPI(api *echo.Group) {
	topicRepo := NewRepository()
	topicSrv := NewServiceImpl(topicRepo)
	server := &ServerImpl{topicSrv: topicSrv}

	// Init topic endpoints
	topicEndpoints := api.Group("/topics")
	topicEndpoints.GET("", server.home)
	topicEndpoints.POST("", server.create, handler.IsLoggedIn, handler.IsAdmin)
	topicEndpoints.GET("/:id", server.detail)
	topicEndpoints.GET("/:id/upvote", server.upvote, handler.IsLoggedIn)
	topicEndpoints.GET("/:id/downvote", server.downvote, handler.IsLoggedIn)

}

func (s *ServerImpl) home(ctx echo.Context) error {
	topics := s.topicSrv.Fetch()

	return ctx.JSON(http.StatusOK, &common.Response{
		Code:    http.StatusOK,
		Content: topics,
	})
}

func (s *ServerImpl) detail(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	topic, err := s.topicSrv.GetTopicById(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		Code:    http.StatusOK,
		Content: topic,
	})
}

func (s *ServerImpl) create(ctx echo.Context) error {
	topic := &Topic{}
	if err := ctx.Bind(&topic); err != nil {
		return err
	}
	topic, err := s.topicSrv.CreateTopic(topic)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, &common.Response{
		Code:    http.StatusCreated,
		Content: topic,
	})
}

func (s *ServerImpl) upvote(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return s.vote(ctx, id, true)
}

func (s *ServerImpl) downvote(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return s.vote(ctx, id, false)
}

func (s *ServerImpl) vote(ctx echo.Context, id int, up bool) error {
	if err := s.topicSrv.Vote(id, up); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
