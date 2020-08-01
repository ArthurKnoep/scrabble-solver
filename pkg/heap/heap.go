package heap

var i = 0

func Heap(k int, array []rune, handler func (permutation []rune)) {
	if k == 1 {
		i += 1
		handler(array)
	} else {
		Heap(k-1, array, handler)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				tmp := array[k-1]
				array[k-1] = array[i]
				array[i] = tmp
			} else {
				tmp := array[k-1]
				array[k-1] = array[0]
				array[0] = tmp
			}
			Heap(k-1, array, handler)
		}
	}
}
