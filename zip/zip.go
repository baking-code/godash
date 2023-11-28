package zip

import "fmt"

/*
Creates an array of grouped elements, the first of which contains the first elements of the given arrays, the second of which contains the second elements of the given arrays, and so on.
*/
func Zip[T interface{}](arrays ...[]T) [][]T {
	zipLength := len(arrays)
	zipped := make([][]T, zipLength)
	for i := 0; i < zipLength; i++ {
		zipped[i] = make([]T, len(arrays[i]))
		for j := 0; j < len(arrays[i]); j++ {
			zipped[i][j] = arrays[j][i]
		}

	}
	return zipped
}

/*
This method is like _.zip except that it accepts an array of grouped elements and creates an array regrouping the elements to their pre-zip configuration.
*/
func Unzip[T interface{}](zipped [][]T) [][]T {
	zipLength := len(zipped)
	unzipped := make([][]T, zipLength)
	for i := 0; i < zipLength; i++ {
		unzipped[i] = make([]T, len(zipped[i]))
		for j := 0; j < len(zipped[i]); j++ {
			fmt.Println(zipped[i], zipped[i][j], i, j)
			unzipped[i][j] = zipped[j][i]
		}

	}
	return unzipped
}
