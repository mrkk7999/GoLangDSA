package rotated_Sorted_Array_Search

// https://www.interviewbit.com/problems/rotated-sorted-array-search/

func search(A []int, B int) int {
	// return rotBinSearch(A,start,end,B)
	var start, end int = 0, len(A) - 1
	for start <= end {
		mid := start + (end-start)/2
		if A[mid] == B {
			return mid
		}
		if A[start] <= A[mid] {
			if (B >= A[start]) && (B < A[mid]) {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if (B <= A[end]) && (B > A[mid]) {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
