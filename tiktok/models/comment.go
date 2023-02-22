package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id         int64     `json:"id"`
	UserInfoId int64     `json:"-"` 
	VideoId    int64     `json:"-"` 
	User       UserInfo  `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"`
}

type CommentDAO struct {
}

var (
	commentDao CommentDAO
)

func NewCommentDAO() *CommentDAO {
	return &commentDao
}

func (c *CommentDAO) AddCommentAndUpdateCount(comment *Comment) error {
	if comment == nil {
		return errors.New("AddCommentAndUpdateCount comment空指针")
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		//添加评论数据
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		//增加count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count+1 WHERE v.id=?", comment.VideoId).Error; err != nil {
			return err
		}

		return nil
	})
}

func (c *CommentDAO) DeleteCommentAndUpdateCountById(commentId, videoId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		//删除评论
		if err := tx.Exec("DELETE FROM comments WHERE id = ?", commentId).Error; err != nil {
			return err
		}
		//减少count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count-1 WHERE v.id=? AND v.comment_count>0", videoId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (c *CommentDAO) QueryCommentById(id int64, comment *Comment) error {
	if comment == nil {
		return errors.New("QueryCommentById comment nullptr")
	}
	return DB.Where("id=?", id).First(comment).Error
}

func (c *CommentDAO) QueryCommentListByVideoId(videoId int64, comments *[]*Comment) error {
	if comments == nil {
		return errors.New("QueryCommentListByVideoId comments nullptr")
	}
	if err := DB.Model(&Comment{}).Where("video_id=?", videoId).Find(comments).Error; err != nil {
		return err
	}
	return nil
}
