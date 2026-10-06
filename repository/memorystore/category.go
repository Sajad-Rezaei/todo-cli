package memorystore

import "todocli/entity"

type Category struct{
	categories []entity.Category
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