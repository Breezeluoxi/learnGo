package main

// 定义发送到 OpenAI 的请求体结构
type OpenAIRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

func main() {
	print(findMedianSortedArrays([]int{2, 3}, []int{}))
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	for i := (m + n) - 1; i >= 0; i-- {
		if index1 >= 0 && index2 >= 0 {

			if nums1[index1] >= nums2[index2] {
				nums1[i] = nums1[index1]
				index1--
			} else {
				nums1[i] = nums2[index2]
				index2--
			}
		} else if index1 >= 0 {
			nums1[i] = nums1[index1]
			index1--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}
}

func removeDuplicates(nums []int) int {
	result := 1
	temp := nums[0]
	index := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != temp {
			temp = nums[i]
			result++
			nums[index] = nums[i]
			index++
		} else if count == 1 && nums[i] == temp {
			temp = nums[i]
			result++
			nums[index] = nums[i]
			index++
			count++
		} else {

			temp = nums[i]
			result++
			nums[index] = nums[i]
			index++

		}

	}

	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	index1 := len(nums1) - 1
	index2 := len(nums2) - 1
	toLen := len(nums1) + len(nums2)

	if toLen == 1 {
		if index1 == 0 {
			return float64(nums1[index1])
		} else {
			return float64(nums2[index2])
		}
	}

	for i := toLen - 1; i > toLen-1-(toLen-1)/2; i-- {
		if index1 >= 0 && index2 >= 0 {
			if nums1[index1] >= nums2[index2] {
				index1--
			} else {

				index2--
			}
		} else if index1 >= 0 {

			index1--
		} else {

			index2--
		}

		if i-1 == toLen-1-(toLen-1)/2 {
			if toLen%2 == 0 {
				if index1 >= 0 && index2 >= 0 {
					return (float64)(nums2[index2]+nums1[index1]) / 2
				} else if index1 >= 0 {
					return (float64)(nums1[index1]+nums1[index1-1]) / 2
				} else {
					return (float64)(nums2[index2]+nums2[index2-1]) / 2
				}
			} else {
				if index1 >= 0 && index2 >= 0 {
					if nums1[index1] >= nums2[index2] {
						return float64(nums1[index1])
					} else {
						return float64(nums2[index2])
					}
				} else if index1 >= 0 {
					return float64(nums1[index1])
				} else {
					return float64(nums2[index2])
				}
			}

		}
	}

	if toLen%2 == 0 {
		if index1 >= 0 && index2 >= 0 {
			return (float64)(nums2[index2]+nums1[index1]) / 2
		} else if index1 >= 0 {
			return (float64)(nums1[index1]+nums1[index1-1]) / 2
		} else {
			return (float64)(nums2[index2]+nums2[index2-1]) / 2
		}
	} else {
		if index1 >= 0 && index2 >= 0 {
			if nums1[index1] >= nums2[index2] {
				return float64(nums1[index1])
			} else {
				return float64(nums2[index2])
			}
		} else if index1 >= 0 {
			return float64(nums1[index1])
		} else {
			return float64(nums2[index2])
		}
	}
	return 0
}
