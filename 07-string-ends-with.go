package kata

func solution(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str)-len(ending):] == ending
}
