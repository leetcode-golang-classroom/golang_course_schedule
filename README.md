# golang_course_schedule

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you **must** take course `bi` first if you want to take course `ai`.

- For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return `true` if you can finish all courses. Otherwise, return `false`.

## Examples

**Example 1:**

```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.

```

**Example 2:**

```
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

```

**Constraints:**

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- All the pairs prerequisites[i] are **unique**.

## 解析

給定一個正整數 numCourses，代表有 0 到 numCourses - 1的課程

還有一個 矩陣 prerequisites 每個 entry [$a_i,$ $b_i$] 代表 要完成 $a_i$ 課程必須先完成 $b_i$ 課程

題目要求寫出一個演算法去判斷給定的 numCourses, 還有 prerequisites 能不能夠完成

根據 prerequisites 可以先畫出 dependency 關係圖

![](https://i.imgur.com/uG1RAT4.png)

根據 preMap 來做 DFS 並且紀錄每個走訪過的 node

當發現遇到重複拜訪的點代表出現了 dependency 循環 所以 return false

如果走訪完所有 dependency 都沒有重複出現點代表可以完成

![](https://i.imgur.com/yzYxNet.png)
## 程式碼
```go
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

```
## 困難點

1. 理解如何找出有循環的 dependency
2. 理解透過 adjacency map 就可以表示每個 course 的 dependency 順序
3. 理解透過 HashSet 儲存找到有走訪過代表無法完成

## Solve Point

- [x]  理解如何找出有循環的 dependency
- [x]  透過 adjacency map 就可以表示每個 course 的 dependency 順序
- [x]  透過 DFS 來走訪 dependency graph
- [x]  透過 hashSet 儲存已經走訪過的 課程，如果有遇到則代表有 cycle 代表無法完成