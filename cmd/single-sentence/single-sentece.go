package main

import (
	"single-sentence/routes"
)

type T struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUid int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func main() {
	r := routes.NewRouter()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}

/*for i := 0; i < 100; i++ {
	var t T
	var xyq data.Xyq777
	resp, err := http.Get("https://v1.hitokoto.cn")
	if err != nil {
		fmt.Println(err)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonData, &t)
	if err != nil {
		fmt.Println(err)
	}
	xyq.Xyq777 = t.Hitokoto
	xyq.Type = t.Type
	xyq.From = t.From
	data.PostSentenceSql(&xyq)
	data.PostSentenceRedis(&xyq)
	time.Sleep(1 * time.Second)
	resp.Body.Close()
}*/
