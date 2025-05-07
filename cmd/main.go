package main

import "fmt"

func MapAny[T any, V any](arr []T, f func(T) V) []V {
	res := make([]V, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}

	return res
}

func filterNew[T any](arr []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}

	return res
}

func reduceNew[T any, V any](arr []T, f func(V, T) V, initial V) V {
	acc := initial

	for _, v := range arr {
		acc = f(acc, v)
	}
	return acc
}

func main() {
	arr := make([]map[string]interface{}, 0)
	obj := make(map[string]interface{})

	obj["id"] = 1
	obj["name"] = "John Doe"
	obj["age"] = 30
	obj["isEmployed"] = true
	obj["skills"] = []string{"Go", "Python", "JavaScript"}
	obj["address"] = map[string]interface{}{
		"street": "123 Main St",
		"city":   "Anytown",
		"state":  "CA",
		"zip":    "12345",
	}

	obj["projects"] = map[string]interface{}{
		"name":         "Project A",
		"description":  "A sample project",
		"technologies": []string{"Go", "Docker"},
		"status":       "completed",
	}

	arr = append(arr, obj)

	arr = MapAny(arr, func(m map[string]interface{}) map[string]interface{} {
		m["salary"] = 1000
		if m["id"] == 1 {
			m["salary"] = 5000
		}

		return m
	})

	arr = filterNew(arr, func(m map[string]interface{}) bool {
		if salary, ok := m["salary"].(int); ok && salary == 1000 {
			return true
		}

		return false
	})

	fmt.Print("arr: ", arr, "\n")

	arrNum := []int{1, 2, 3, 4, 5}
	newNum := reduceNew(arrNum, func(acc int, value int) int {
		return acc + value
	}, 0)

	fmt.Println(newNum)
}
