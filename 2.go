// func countOdds(low int, high int) int {
//     result:= 0
//     i:= low

//     for i <= high {
//         if i%2 == 1 {
//             result++
//         }
//         i++
//     }

//     return result
// }
//O( high - low )
//tuc la O(N)
//cach giai 2 => toi uu hon
// n = high - low + 1 .. so phan tử
// neu day là chẵn => 1 nửa chẵn 1 nửa lẻ .. chia 2 ra lẻ r
// 4 5 6 7 8 9 10 11 12  =>   9  / 2   .  3 4 5 6 7 -> 5
package main

func countOdds(low int, high int) int {
	result := 0
	n := high - low + 1
	if low%2 == 1 {
		result = n/2 + 1
	} else {
		result = n / 2
	}
	return result
}
