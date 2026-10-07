package memorystore

import "todocli/entity"

type Category struct{
	categories []entity.Category
}

func NewCategoryStore() *Category {
	return &Category{
		categories: make([]entity.Category, 0),
	}
}

func (c Category) DoesThisUserHaveThisCategoryID(userID, categoryID int) bool {
	
	isfound := false
	for _, c := range c.categories {
		if c.Id == categoryID && c.UserID == userID {
			isfound = true

			break
		}
	}

	return isfound
}

func (c *Category) CreateNewCategory(cat entity.Category) (entity.Category,error){
	cat.Id = len(c.categories)+1
	c.categories= append(c.categories, cat)
	return cat,nil
}