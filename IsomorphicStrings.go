func isIsomorphic(s string, t string) bool {
	dictS, dictT := make(map[string]int), make(map[string]int)
	countS, countT := 0, 0
	equivalentS, equivalentT := "", ""
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if val, ok := dictS[string(rune(s[i]))]; ok {
			equivalentS = equivalentS + strconv.Itoa(val)
		} else {
			countS++
			dictS[string(rune(s[i]))] = countS
			equivalentS = equivalentS + strconv.Itoa(countS)
		}

		if val, ok := dictT[string(rune(t[i]))]; ok {
			equivalentT = equivalentT + strconv.Itoa(val)
		} else {
			countT++
			dictT[string(rune(t[i]))] = countT
			equivalentT = equivalentT + strconv.Itoa(countT)
		}
	}
	return equivalentS == equivalentT
}