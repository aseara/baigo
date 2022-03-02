package store

import "errors"

var (
	// ErrNotFound error no found
	ErrNotFound = errors.New("not found")
	// ErrExist error exist
	ErrExist = errors.New("exist")
)

// Book ... book
type Book struct {
	ID      string   `json:"id"`      // 图书ISBN Id
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社
}

// Store store
type Store interface {
	Create(*Book) error       // 创建一个新图书条目
	Update(*Book) error       // 更新某图书条目
	Get(string) (Book, error) // 获取某图书信息
	GetAll() ([]Book, error)  // 获取所有图书信息
	Delete(string) error      // 删除某图书条目
}
