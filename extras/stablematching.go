package extras

func StableMatching(group, items [][]int) map[int]int {
	numOfGroups, numOfItems := len(group), len(items)
	groupMap := make(map[int]int)
	itemMap := make(map[int]int)
	for i := 0; len(itemMap) <= numOfGroups && i < numOfItems; i++ {
		for g := 0; len(groupMap) <= numOfItems && g < numOfGroups; g++ {
			if _, grpOk := groupMap[g]; !grpOk {
				curItem := group[g][i]
				if oldGrp, itemOk := itemMap[curItem]; !itemOk {
					groupMap[g] = curItem
					itemMap[curItem] = g
				} else {
					preferredGroup := oldGrp
					for _, grp := range items[curItem] {
						if grp == g || grp == oldGrp {
							preferredGroup = grp
							break
						}
					}
					delete(itemMap, curItem)
					delete(groupMap, oldGrp)
					groupMap[preferredGroup] = curItem
					itemMap[curItem] = preferredGroup
				}
			}
		}
	}
	return groupMap
}
