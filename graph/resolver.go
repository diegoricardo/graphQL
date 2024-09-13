package graph

import "github.com/diegoricardo/graphQL/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
