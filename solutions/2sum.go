package solutions

func TwoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        if _, ok := m[target-nums[i]]; ok {
            return []int{i, m[target-nums[i]]}
        } else {
            m[nums[i]] = i
        }
    }
    return nil
}
