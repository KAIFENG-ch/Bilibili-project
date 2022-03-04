package serialize

import "Bilibili-project/model"

type Barrage struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Content string `json:"content"`
	Time uint `json:"time"`
}

func BuildBarrages(barrage []model.Barrage) *Base {
	return &Base{
		Status: 200,
		Msg: "ok",
		Data: Datalist{
			Item: barrage,
			Total: len(barrage),
		},
	}
}
