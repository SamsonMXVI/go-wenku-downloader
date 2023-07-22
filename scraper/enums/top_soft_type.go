package enums

type TopSortType int

const (
	TopSoftAnime TopSortType = iota
	TopSoftUpdateToday
	TopSoftNewBook
	TopSoftTotalFavorites
	TopSoftWordCount

	TopSoftOverallLeardboard
	TopSoftMonthlyLeaderboard
	TopSoftWeeklyLeaderboard
	TopSoftDailyLeaderboard

	TopSoftOverallRecommendation
	TopSoftMonthlyRecommendation
	TopSoftWeeklyRecommendation
	TopSoftDailyRecommendation
)

var TopSoftTextReq = []string{
	"anime",
	"lastupdate",
	"postdate",
	"goodnum",
	"size",

	"allvisit",
	"monthvisit",
	"weekvisit",
	"dayvisit",

	"allvote",
	"monthvote",
	"weekvote",
	"dayvote",
}

var TopSoftText = []string{
	"动画化作品",
	"今日更新",
	"(新书一览/最新入库)",
	"总收藏榜",
	"字数排行",

	"(总排行榜/热门轻小说)",
	"月排行榜",
	"周排行榜",
	"日排行榜",

	"总推荐榜",
	"月推荐榜",
	"周推荐榜",
	"日推荐榜",
}
