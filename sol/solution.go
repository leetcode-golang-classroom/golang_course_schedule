package sol

type Courses []int

func canFinish(numCourses int, prerequisites [][]int) bool {
	preCourseMap := make(map[int]Courses, numCourses)
	visit := make(map[int]struct{})
	// init preCourseMap
	for _, dependency := range prerequisites {
		preCourseMap[dependency[0]] = append(preCourseMap[dependency[0]], dependency[1])
	}
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if _, ok := visit[course]; ok {
			return false
		}
		if len(preCourseMap[course]) == 0 {
			return true
		}
		visit[course] = struct{}{}
		for _, preCourse := range preCourseMap[course] {
			if !dfs(preCourse) {
				return false
			}
		}
		delete(visit, course)
		preCourseMap[course] = []int{}
		return true
	}
	for idx := 0; idx < numCourses; idx++ {
		if !dfs(idx) {
			return false
		}
	}
	return true
}
