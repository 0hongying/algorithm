package main

import "sort"

// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
// 给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数

func findMinArrowShots(points [][]int) int {
	// 以右边界排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	ret := 0
	for i := 0; i < len(points); i++ {
		right := points[i][1]
		// 如果以基准的右边界大于其左边界，则说明可以被一支箭引爆
		for j := i + 1; j < len(points) && right >= points[j][0]; {
			j++
			i++
		}
		ret++
	}

	return ret
}
