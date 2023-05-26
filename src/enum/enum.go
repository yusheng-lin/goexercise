package enum

type Sort byte

const (
	ASC Sort = iota
	Desc
)

type UserSort byte

const (
	Account UserSort = iota
	Fullname
	CreatedAt
)

func (sort UserSort) String() string {
	switch sort {
	case Fullname:
		return "fullname"
	case CreatedAt:
		return "created_at"
	default:
		return "acct"
	}
}

func (sort Sort) String() string {
	if sort == Desc {
		return "DESC"
	}
	return "ASC"
}
