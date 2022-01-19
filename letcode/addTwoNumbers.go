package letcode

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/* 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitList(s string) *ListNode {
	var firstNode, lastNode *ListNode

	for i := len(s) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(string(s[i]))
		node := &ListNode{
			Val: val,
		}

		if firstNode == nil {
			firstNode = node
			lastNode = node
		} else {
			lastNode.Next = node
			lastNode = node
		}
	}

	return firstNode
}

func GetNumber(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head  = &ListNode{}
		tail  = head
		carry = 0
	)
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		val := n1 + n2 + carry
		val, carry = val%10, val/10

		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head.Next
}

/* 给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
func KSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}

	h.data = append(h.data, pair{0, 0})
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}

		if j == 0 && i+1 < m {
			heap.Push(&h, pair{i + 1, 0})
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }

/* 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	var (
		m, n       = len(nums1), len(nums2)
		start, end = 0, len(nums1)
		num1       = (start + end) / 2
		num2       = (m+n)/2 - num1
	)

	for {
		if num1 > 0 && nums1[num1-1] > nums2[num2] {
			end = num1 - 1
		} else if num1 < m && nums1[num1] < nums2[num2-1] {
			start = num1 + 1
		} else {
			break
		}

		num1 = (start + end) / 2
		num2 = (m+n)/2 - num1
	}

	var big int
	if num1 == m || (num2 < n && nums1[num1] > nums2[num2]) {
		big = nums2[num2]
	} else {
		big = nums1[num1]
	}

	if (m+n)%2 == 1 {
		return float64(big)
	}

	if (m+n)%2 == 1 {
		return float64(big)
	}

	var small int
	if num1 == 0 || (num2 > 0 && nums1[num1-1] < nums2[num2-1]) {
		small = nums2[num2-1]
	} else {
		small = nums1[num1-1]
	}

	return float64(big+small) / 2.0
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/* 给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。*/
func FindMinDifference(timePoints []string) int {
	times := make([]int, 0, len(timePoints)+1)
	for i := 0; i < len(timePoints); i++ {
		times = append(times, getIntTime(timePoints[i]))
	}

	sort.Ints(times)
	times = append(times, times[0]+24*60)
	minDelta := 24 * 60
	for i := 0; i < len(timePoints); i++ {
		minDelta = min(minDelta, times[i+1]-times[i])
	}
	return minDelta
}

func getIntTime(strTime string) int {
	strs := strings.Split(strTime, ":")
	hours, _ := strconv.Atoi(strs[0])
	minutes, _ := strconv.Atoi(strs[1])
	return hours*60 + minutes
}
