func myPow(x float64, n int) float64 {
    res := 1.0
    if n == 0 {
        return res
    } else if n < 0 {
        return myPow(1.0/x, -n)
    } else {
        res = x
        k := 2
        for ; k < n; k=k*2 {
            res *= res
        }
        return res * myPow(x, n - k/2)

    }
}