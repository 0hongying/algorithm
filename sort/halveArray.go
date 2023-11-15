package main

import (
	"container/heap"
)

type MaxHeap struct {
	nums []float64
}

func (m *MaxHeap) Len() int {
	return len(m.nums)
}

func (m *MaxHeap) Less(i, j int) bool {
	// 大顶堆： m.nums[i] > m.nums[j]
	// 小顶堆： m.nums[i] < m.nums[j]
	return m.nums[j] < m.nums[i]
}

func (m *MaxHeap) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
}

func (m *MaxHeap) Pop() interface{} {
	n := m.nums[m.Len()-1]
	m.nums = m.nums[:m.Len()-1]
	return n
}

func (m *MaxHeap) Push(n interface{}) {
	m.nums = append(m.nums, n.(float64))
}

// 给你一个正整数数组nums。每一次操作中，你可以从nums中选择任意一个数并将它减小到恰好一半
// 请你返回将 nums 数组和至少减少一半的最少操作数
// nums = [5,19,8,1]
// 19 / 2 = 9.5
// 9.5 / 2 = 4.75
// 8 / 2 = 4
// 5 + 4.75 + 4 + 1 = 14.75
// 33 - 14.75 = 18.25
// 18.25 >= 33/2 = 16.5

func halveArray(nums []int) int {
	sum := 0
	ret := 0
	maxHeap := &MaxHeap{}
	for _, n := range nums {
		sum += n
		maxHeap.nums = append(maxHeap.nums, float64(n))
	}

	// 总和的一半
	halfSum := float64(sum) / 2.0
	// 剩余数字之和
	leftSum := float64(sum)
	// 初始化堆
	heap.Init(maxHeap)

	for {
		m := (heap.Pop(maxHeap)).(float64)
		ret += 1
		leftSum -= m / 2.0
		if leftSum <= halfSum {
			break
		}
		heap.Push(maxHeap, m/2.0)
	}

	return ret
}

// func main() {
// 	nums := []int{
// 		5, 19, 8, 1,
// 	}
// 	ret := halveArray(nums)
// 	fmt.Printf("%v", ret)
// }
