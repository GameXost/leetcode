package hard

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArray := make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0
	cnt := 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				mergedArray[cnt] = nums1[i]
				i++
			} else {
				mergedArray[cnt] = nums2[j]
				j++
			}
		} else if i < len(nums1) {
			mergedArray[cnt] = nums1[i]
			i++
		} else if j < len(nums2) {
			mergedArray[cnt] = nums2[j]
			j++
		}
		cnt++
	}
	lenMerged := len(mergedArray)
	if lenMerged%2 == 1 {
		return float64(mergedArray[lenMerged/2])
	}
	return (float64(mergedArray[lenMerged/2]) + float64(mergedArray[lenMerged/2-1])) / 2
}
