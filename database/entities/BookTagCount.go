package entities

type BookTagCount struct {
	Tag          	string			`json:"title"`
	Count			int				`json:"count"`

	BookId			int				`json:"book_id"`
	Book			*Book			`json:"book,omitempty" gorm:"ForeignKey:BookId"`

	//housekeeping attributes
	ID        		int				`json:"id"`
}

//for GORM
func (BookTagCount) TableName() string {
	return "book_tags_count"
}