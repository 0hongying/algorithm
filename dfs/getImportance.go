package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// employees:[[1, 5, [2, 3]], [2, 3, []], [3, 3, []]]
// id: 1
// 11
// 员工 1 自身的重要度是 5 ，他有两个直系下属 2 和 3 ，而且 2 和 3 的重要度均为 3 。因此员工 1 的总重要度是 5 + 3 + 3 = 11

func getImportance(employees []*Employee, id int) int {
	mp := map[int]*Employee{}
	for _, employee := range employees {
		mp[employee.Id] = employee
	}

	ans := 0
	var dfs func(id int)
	dfs = func(id int) {
		e := mp[id]
		ans += e.Importance
		sub := e.Subordinates
		for _, v := range sub {
			dfs(v)
		}
	}

	dfs(id)
	return ans
}
