package models

// User ユーザー情報のモデル
type User struct {
	ID       int    `json:"id" example:"1"`                    // ユーザーID
	Name     string `json:"name" example:"Leanne Graham"`      // フルネーム
	Username string `json:"username" example:"Bret"`           // ユーザー名（ログイン名など）
	Email    string `json:"email" example:"Sincere@april.biz"` // メールアドレス

	Address struct {
		Street  string `json:"street" example:"Kulas Light"` // 番地
		Suite   string `json:"suite" example:"Apt. 556"`     // 部屋番号やスイート名
		City    string `json:"city" example:"Gwenborough"`   // 市
		Zipcode string `json:"zipcode" example:"92998-3874"` // 郵便番号

		Geo struct {
			Lat string `json:"lat" example:"-37.3159"` // 緯度
			Lng string `json:"lng" example:"81.1496"`  // 経度
		} `json:"geo"` // 緯度経度
	} `json:"address"` // 住所情報

	Phone   string `json:"phone" example:"1-770-736-8031 x56442"` // 電話番号
	Website string `json:"website" example:"hildegard.org"`       // WebサイトURL

	Company struct {
		Name        string `json:"name" example:"Romaguera-Crona"`                               // 会社名
		CatchPhrase string `json:"catchPhrase" example:"Multi-layered client-server neural-net"` // キャッチフレーズ
		Bs          string `json:"bs" example:"harness real-time e-markets"`                     // 業務内容
	} `json:"company"` // 会社情報
}
