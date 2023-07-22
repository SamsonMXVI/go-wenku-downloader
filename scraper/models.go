package scraper

const (
	BASE_URL           = "https://www.wenku8.net/book/"
	DEFAULT_USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36"
	SEARCH_URL         = "https://www.wenku8.net/modules/article/search.php"
	TOP_URL            = "https://www.wenku8.net/modules/article/toplist.php"
)

type Novel struct {
	NovelId        int    `json:"novel_id"`
	NovelName      string `json:"novel_name"`
	Cover          string `json:"cover"`
	Library        string `json:"library"`
	Author         string `json:"author"`
	Status         string `json:"status"`
	LastUpdateTime string `json:"last_update_time"`
	Length         string `json:"length"`
	Tag            string `json:"tag"`
	RecentChapter  string `json:"recent_chapter"`
	Desc           string `json:"desc"`
	CatalogueUrl   string `json:"catalogue_url"`
}

type Volume struct {
	Index        int    `json:"index"`
	Name         string `json:"name"`
	RowNumber    int    `json:"-"`
	EndRow       int    `json:"-"`
	CatalogueUrl string `json:"catalogue_url"`
	ChapterCount int    `json:"chapter_count"`
}

type Chapter struct {
	Index   int             `json:"index"`
	Title   string          `json:"title"`
	Url     string          `json:"url"`
	Content *ChapterContent `json:"content"`
}

type CommandOptions struct {
	Epub       bool
	Ext        string
	OnlyImages bool
	OutDir     string
	Verbose    bool
	Strict     bool
}

type ChapterContent struct {
	Article string   `json:"article"`
	Images  []string `json:"images"`
}

type PageResult struct {
	TotalPage  string   `json:"total_page"`
	NovelArray []*Novel `json:"result"`
}
