package merge

/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output Integer array.
 */
func solve(A []int, B []int) []int {
	var res []int
	var i, j = 0, 0
	for i < len(A) && j < len(B) {
		if A[i] <= B[j] {
			res = append(res, A[i])
			i++
		} else {
			res = append(res, B[j])
			j++
		}
	}

	for i < len(A) {
		res = append(res, A[i])
		i++
	}

	for j < len(B) {
		res = append(res, B[j])
		j++
	}

	return res
}
