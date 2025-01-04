func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, v := range flowerbed {
        if (i == 0 || flowerbed[i-1] == 0) && 
        (len(flowerbed) - 1 == i || flowerbed[i+1] == 0) && 
        v == 0 {
            flowerbed[i] = 1
            n--
        }
	}
	return n <= 0
}
