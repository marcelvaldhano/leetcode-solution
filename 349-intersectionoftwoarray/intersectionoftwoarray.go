package main

func intersection(nums1 []int, nums2 []int) []int {
	dict := make(map[int]bool, 1001)
	res := []int{}
	for i := range nums1 {
		dict[nums1[i]] = true
	}
	for i := range nums2 {
		if dict[nums2[i]] == true {
			res = append(res, nums2[i])
			dict[nums2[i]] = false
		}
	}
	return res
}
