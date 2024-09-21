// Removes values from index1 to index2 inclusively.
// Given: [“M”, “N”, “O”, “P”, “Q”, “R”] and index1 = 2 and index2 = 4
// Result: [M N Q R]
func removeFrom(slice []string, index1, index2 int) []string {
    result := make([]string, len(slice) - (index2 - index1))
    at := copy(result, slice[:index1])
    copy(result[at:], slice[index2:])
    return result
}
