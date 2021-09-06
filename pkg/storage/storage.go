package storage

// Post - публикация.
type Post struct {
	ID      int    `xml:"-" json:"ID"`                // номер записи
	Title   string `xml:"title" json:"Title"`         // заголовок публикации
	Content string `xml:"description" json:"Content"` // содержание публикации
	PubDate string `xml:"pubDate" json:"-"`           // время публикации из RSS
	PubTime int64  `xml:"-" json:"PubTime"`           //время публикации для БД и фронта
	Link    string `xml:"link" json:"Link"`           // ссылка на источник
}

//Модель полной формы новости
type NewsFullDetailed struct {
	ID      int    `xml:"-" json:"ID"`                // номер записи
	Title   string `xml:"title" json:"Title"`         // заголовок публикации
	Content string `xml:"description" json:"Content"` // содержание публикации
	PubDate string `xml:"pubDate" json:"-"`           // время публикации из RSS
	PubTime int64  `xml:"-" json:"PubTime"`           //время публикации для БД и фронта
	Link    string `xml:"link" json:"Link"`           // ссылка на источник
}

//Модель короткой формы новости
type NewsShortDetailed struct {
	ID      int    `xml:"-" json:"ID"`                // номер записи
	Title   string `xml:"title" json:"Title"`         // заголовок публикации
	Content string `xml:"description" json:"Content"` // содержание публикации
	PubDate string `xml:"pubDate" json:"-"`           // время публикации из RSS
	PubTime int64  `xml:"-" json:"PubTime"`           //время публикации для БД и фронта
	Link    string `xml:"link" json:"Link"`           // ссылка на источник
}

//Модель комментария
type Comment struct {
	ID      int    `xml:"-" json:"ID"`                // номер записи
	Title   string `xml:"title" json:"Title"`         // заголовок публикации
	Content string `xml:"description" json:"Content"` // содержание публикации
	PubDate string `xml:"pubDate" json:"-"`           // время публикации из RSS
	PubTime int64  `xml:"-" json:"PubTime"`           //время публикации для БД и фронта
	Link    string `xml:"link" json:"Link"`           // ссылка на источник
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error)     // получение всех публикаций
	PostsN(int) ([]Post, error) // получение n последних публикаций
	PostsByFilter(string, string,
		int, int) ([]Post, error) // получение публикаций по фильтру
	AddPost(Post) error    // создание новой публикации
	UpdatePost(Post) error // обновление публикации
	DeletePost(Post) error // удаление публикации по ID
	Close()                // освобождение ресурса
	//DropDB() error              //удаление БД
}
