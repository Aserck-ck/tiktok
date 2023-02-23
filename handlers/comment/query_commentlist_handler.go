package comment

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Aserck-ck/tiktok/handlers/video"
	"github.com/Aserck-ck/tiktok/models"
	"github.com/Aserck-ck/tiktok/service/comment"
	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	models.CommonResponse
	*comment.List
}

func QueryCommentListHandler(c *gin.Context) {
	NewProxyCommentListHandler(c).Do()
}

type ProxyCommentListHandler struct {
	*gin.Context

	videoId int64
	userId  int64
}

func NewProxyCommentListHandler(context *gin.Context) *ProxyCommentListHandler {
	return &ProxyCommentListHandler{Context: context}
}

func (p *ProxyCommentListHandler) Do() {
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	commentList, err := comment.QueryCommentList(p.userId, p.videoId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	p.SendOk(commentList)
}

func (p *ProxyCommentListHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId went wrong")
	}
	p.userId = userId

	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	return nil
}

func (p *ProxyCommentListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, video.FavorVideoListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyCommentListHandler) SendOk(commentList *comment.List) {
	p.JSON(http.StatusOK, ListResponse{CommonResponse: models.CommonResponse{StatusCode: 0},
		List: commentList,
	})
}
