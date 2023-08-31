package entities

type User struct {
	ID int64 `db:"user_id"`
}

type AddedDeleted struct {
	Added   []string
	Deleted []string
}
