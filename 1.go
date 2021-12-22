package main

// func numIdenticalPairs(nums []int) int {
//     result := 0
//     for i := 0 ; i < len(nums) -1;i++{
//         for j:= i+1 ; j< len(nums) ; j++{
//             if(nums[i] == nums[j] && i < j){
//                 result += 1;
//             }
//         }
//     }
//     return result
// }
// O(n*n)

func numIdenticalPairs(nums []int) int {
	result := 0
	//map key : value
	//key : phan tu trong mang
	//value : so lan xuat hien cua phan tu do
	numberCountMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if count, ok := numberCountMap[nums[i]]; ok {
			result = result + count
			numberCountMap[nums[i]] = count + 1
		} else {
			numberCountMap[nums[i]] = 1
		}
	}
	return result
}
func main() {

}
