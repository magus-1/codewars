package kata

import "fmt"

func Josephus(items []interface{}, k int) []interface{} {
	var result, winners []interface{}
	newItems := items
	i := 1

	if len(items) == 0 {
		return items
	}

	for {
		// outer loop: go through the entire slice once
		remaining := len(items) - len(result)

		if remaining == 0 {
			break
		}

		winners, newItems = SingleRound(newItems[:remaining], k)
		result = append(result, winners...)
		fmt.Printf("o: %v, newItems: %v \n", i, newItems)
		i++
	}
	return result
}

func SingleRound(items []interface{}, k int) ([]interface{}, []interface{}) {
	var winners []interface{}
	newItems := items
	i := 1

	for {
		// inner loop: pick winners, and remove them
		s := (i * (k - 1)) % len(items)
		winners = append(winners, items[s])
		newItems = append(newItems[0:s], newItems[s+1:]...)

		if ((i+1)*k-len(items)) > 0 || len(items)-i == 0 {
			fmt.Println("loop break!")
			// re-sorting so the next loop starts at the point we leave here
			sortedItems := append(newItems[s:], newItems[:s]...)
			return winners, sortedItems
		}
		i++
	}
}
