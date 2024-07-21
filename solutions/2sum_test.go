package solutions

import (
    "go_lc/helper"
    "testing"
)

func TestTwoSum(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := TwoSum(nums, target)
    if result == nil {
        t.Fatalf("Not Expect nill")
    }
    if !helper.SlicesEqual(result, []int{0, 1}) && !helper.SlicesEqual(result, []int{1, 0}) {
        t.Fatalf("wrong ansower")
    }

    nums = []int{2, 7, 11, 15}
    target = 1
    result = TwoSum(nums, target)
    if result != nil {
        t.Fatalf("Expect nill")
    }
}
