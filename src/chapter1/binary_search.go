package chapter1

// 二分查找
// arr必须是有序的
// @return -1:searchKey在arr中不存在;或searchKey在arr中的索引
func binarySearch(searchKey int, arr []int) int {
	low, mid, high := 0, -1, len(arr)-1
	for low <= high {
		//mid = (low + high) / 2
		mid = low + (high-low)/2
		if searchKey == arr[mid] {
			return mid
		}
		if searchKey < arr[mid] {
			high = mid - 1
			continue
		}
		low = mid + 1
	}
	return mid
}
