package stringmatching

func LCS(str string, substr string) int {
	var res int
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
    	dp[i] = make([]int, len(substr) + 1)
	}
	res = 0
	for i := 1; i <= len(str); i++ {
        for j := 1; j <= len(substr); j++ {
            if (str[i - 1] == substr[j - 1]) {
                dp[i % 2][j] = dp[(i - 1) % 2][j - 1] + 1
                if (dp[i % 2][j] > res) {
					res = dp[i % 2][j]
				}
            } else {
				dp[i % 2][j] = 0
			}
        }
    }
    return res;
}