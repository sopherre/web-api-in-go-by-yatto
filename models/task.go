package models

// Task タスクモデル
type Task struct {
	ID        uint   `gorm:"primaryKey" example:"1"`    // タスクID
	Title     string `json:"title" example:"買い物に行く"`    // タスクのタイトル
	Completed bool   `json:"completed" example:"false"` // 完了状態
}
