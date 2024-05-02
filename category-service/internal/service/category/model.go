package category

type Category struct {
	ID   uint64
	Name string
}

type Categories []*Category

func (c Categories) FilterByID(id uint64) *Category {
	for _, category := range c {
		if category.ID == id {
			return category
		}
	}
	return nil
}
