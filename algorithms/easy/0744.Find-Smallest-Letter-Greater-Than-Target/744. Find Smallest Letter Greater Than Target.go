package main

import (
	"fmt"
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	size := len(letters)
	idx := sort.Search(size, func(i int) bool {
		return letters[i] > target
	})

	if idx >= size {
		return letters[0]
	}

	if letters[idx] <= target {
		return letters[0]
	}
	return letters[idx]
}

// func nextGreatestLetter(letters []byte, target byte) byte {
//     size := len(letters)
//     l,r := 0, size-1
//     for l < r {
//         mid := l + (r-l)/2
//         if letters[mid] > target {
//             r = mid
//         } else {
//             l = mid + 1
//         }
//     }
//     if letters[l] <= target {
//         return letters[0]
//     }
//     return letters[l]
// }

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')))
}
