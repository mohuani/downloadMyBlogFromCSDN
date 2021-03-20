package Model

type Blog struct {
	ArticleId    int
	Title        string
	Description  string
	Url          string
	Type         int
	Top          bool
	ForcePlan    bool
	ViewCount    int
	CommentCount int
	EditUrl      string
	PostTime     string
	DiggCount    int
}
