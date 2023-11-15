package main

// https://leetcode.cn/problems/distribute-money-to-maximum-children/description/?envType=daily-question&envId=2023-09-22

func distMoney(money int, children int) int {
	if money < children {
		return -1
	} else if 8*children < money {
		return children - 1
	} else if 8*(children-1)+4 == money {
		return children - 2
	} else {
		return (money - children) / 7
	}
}
