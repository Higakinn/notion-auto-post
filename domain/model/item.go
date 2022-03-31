package model

type Item struct {
    ID        string    `json:"id"`         //TaskのID
    Title      string `json:"title"`       //Task自体
    Url        string `json:url`
    CreateDate string `json:"create_date"`  //Taskの完了期限
}
