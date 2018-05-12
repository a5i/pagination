package pagination

import (
	"strconv"
)

func min(l, r int) int {
	if l < r {
		return l
	}
	return r
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

// GeneratePagination generates a footer pagination so the user can navigate through the web app.
// 	* currentPage - actual page
//	* totalPages - total available pages
//	* boundaries - how many pages we want to link in the beginning and in the end
//	* around - how many pages we want to link before and after the actual page
// It checks and fixes wrong values. Why? We want to generate anyway in any case.
// I think we can add error value or log wrong parameters but stay it for future.
// It returns an array of page numbers and skipped pages "...". Why array? We want to render result in application and style the pagination.
func GeneratePagination(currentPage, totalPages, boundaries, around int) []string {
	boundaries = max(boundaries, 1)
	around = max(around, 0)
	totalPages = max(totalPages, 0)
	result := make([]string, 0)

	var i = 1
	for ; i <= min(boundaries, totalPages); i++ {
		result = append(result, strconv.Itoa(i))
	}
	if (currentPage - around) > i {
		result = append(result, "...")
		i = currentPage - around
	}
	for ; i <= min(currentPage+around, totalPages); i++ {
		result = append(result, strconv.Itoa(i))
	}
	if (totalPages - boundaries) > i {
		result = append(result, "...")
		i = totalPages - boundaries + 1
	}
	for ; i <= totalPages; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}
