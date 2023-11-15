package main

// 对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。
// 找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	if len(asteroids) == 0 {
		return stack
	}

	for i := 0; i < len(asteroids); {
		// 负数向左永远不会碰撞
		if asteroids[i] > 0 || len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, asteroids[i])
			i++
		} else {
			s := stack[len(stack)-1]
			if s < -asteroids[i] {
				// 不加1，继续比对
				stack = stack[:len(stack)-1]
			} else if s == -asteroids[i] {
				stack = stack[:len(stack)-1]
				i++
			} else {
				// 该元素被碰撞
				i++
			}
		}
	}
	return stack
}

// func main() {
// 	asteroids := []int{-2, -1, 1, 2}
// 	asteroidCollision(asteroids)
// }
