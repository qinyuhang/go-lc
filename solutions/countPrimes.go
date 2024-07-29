package solutions

func countPrimes(n int) int {
    if n <= 1 {
        return 0
    }
    arr := make([]bool, n+1)
    for i := range n {
        arr[i] = true
    }
    arr[0] = false
    arr[1] = false
    for i := range n {
        if arr[i] {
            for j := i; i*j < n; j++ {
                arr[i*j] = false
            }
        }
    }
    res := 0
    for i := range n {
        if arr[i] {
            res++
        }
    }
    return res
}
