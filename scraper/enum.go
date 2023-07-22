package scraper

type SearchType int

const (
	SearchArticleName SearchType = iota
	SearchAuthor
)

var SearchTypeTextReq = []string{
	"articlename",
	"author",
}

var SearchTypeText = []string{
	"小说标题",
	"作者",
}
