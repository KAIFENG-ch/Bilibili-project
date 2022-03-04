package serialize

import "Bilibili-project/model"

type Video struct {
	Author string `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	Video string `json:"video"`
	Click uint `json:"click"`
	CreatedAt int64 `json:"created_at"`
}

func BuildVideos(videos []model.Video) *Base {
	return &Base{
		Status: 200,
		Msg: "OK",
		Data: Datalist{
			Item: videos,
			Total: len(videos),
		},
	}
}
