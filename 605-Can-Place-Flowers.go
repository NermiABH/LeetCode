func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0{
        return true
    }
	for i, v := range flowerbed {
        if (i == 0 || flowerbed[i-1] == 0) && 
        (len(flowerbed) - 1 == i || flowerbed[i+1] == 0) && 
        v == 0 {
            flowerbed[i] = 1
            n--
            if n == 0 {
                return true
            }
        }
	}
	return false
}
