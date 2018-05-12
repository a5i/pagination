package pagination

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T) {
	require.Equal(t, "1 ... 4 5", strings.Join(GeneratePagination(4, 5, 1, 0), " "))
	require.Equal(t, "1 2 3 4 5 6 ... 9 10", strings.Join(GeneratePagination(4, 10, 2, 2), " "))
	require.Equal(t, "1 2", strings.Join(GeneratePagination(1, 2, 2, 2), " "))
	require.Equal(t, "1 2", strings.Join(GeneratePagination(1, 2, 2, 2), " "))
	require.Equal(t, "1", strings.Join(GeneratePagination(1, 1, 2, 2), " "))
	require.Equal(t, "", strings.Join(GeneratePagination(0, 0, 2, 2), " "))
	require.Equal(t, "1 ... 5 ... 10", strings.Join(GeneratePagination(5, 10, 0, 0), " "))
	require.Equal(t, "1 2 3 4 5 6 7 8 9 10", strings.Join(GeneratePagination(5, 10, 100, 100), " "))
	require.Equal(t, "1 2 3 4 5 6 7 8 9 10", strings.Join(GeneratePagination(5, 10, 0, 100), " "))
	require.Equal(t, "1 2 3 4 5 6 7 8 9 10", strings.Join(GeneratePagination(5, 10, 100, 0), " "))

	// incorrect parameters
	require.Equal(t, "1 ... 4 5", strings.Join(GeneratePagination(4, 5, -3, -4), " "))
	require.Equal(t, "1 ... 5", strings.Join(GeneratePagination(-1, 5, 1, 0), " "))
	require.Equal(t, "", strings.Join(GeneratePagination(0, -1, 2, 2), " "))
	require.Equal(t, "", strings.Join(GeneratePagination(-1, -1, -2, -2), " "))
}
