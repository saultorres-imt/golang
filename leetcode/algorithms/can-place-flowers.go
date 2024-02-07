package algorithms

// Every odd number of ceros together starting from 3, can place 1 plant
// 0-0-0-0-0-0
func CanPlaceFlowers(flowerbed []int, n int) bool {
	var count int
	var numOfFlowers int
	var ones int

	for i, n := range flowerbed {
		if n == 0 {
			count++
		} else {
			ones++
			if count == 2 && (i == 2 || i == len(flowerbed)-1) {
				numOfFlowers += 1
			} else if count%2 == 0 && count > 0 {
				if flowerbed[0] == 0 && ones <= 1 {
					numOfFlowers++
				}
				numOfFlowers += (count/2 - 1)
			} else {
				numOfFlowers += (count / 2)
			}
			count = 0
		}
	}

	if count > 2 {
		numOfFlowers += (count / 2)
		if ones == 0 && count%2 != 0 {
			numOfFlowers++
		}
	} else if count == 2 || count == 1 && len(flowerbed) == 1 {
		numOfFlowers += 1
	}

	return n <= numOfFlowers
}
