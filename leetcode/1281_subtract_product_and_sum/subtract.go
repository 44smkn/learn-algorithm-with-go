package subtract

// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
func SubtractProductAndSum(n int) int {
    product := 1
    sum := 0
    for n > 0 {
        product *= n % 10
        sum += n % 10
        n /= 10
    }
    return product - sum
}