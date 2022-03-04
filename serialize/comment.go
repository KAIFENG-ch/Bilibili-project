package serialize

import "Bilibili-project/model"

type Comment struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Like uint `json:"like"`
	Dislike uint `json:"dislike"`
	CreatedAt int64 `json:"created_at"`
}

func BuildComments(comments []model.Comments) *Base {
	return &Base{
		Status: 200,
		Msg: "ok",
		Data: Datalist{
			Item: comments,
			Total: len(comments),
		},
	}
}

func BuildReply(reply model.Replies) *Comment {
	return &Comment{
		ID:        reply.ID,
		Username:  reply.Username,
		Like:      reply.Like,
		Dislike:   reply.Dislike,
		CreatedAt: reply.CreatedAt.Unix(),
	}
}

func BuildReplies(replies []model.Replies) *Base {
	return &Base{
		Status: 200,
		Msg: "ok",
		Data: Datalist{
			Item: replies,
			Total: len(replies),
		},
	}
}