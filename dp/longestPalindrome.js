var longestPalindrome = function(s) {
    const n = s.length
    const dp = new Array(n).fill(false).map(() => new Array(n).fill(false))

    let ans = ''

    // 中心扩散法
    for (let i = n-1; i >= 0; i--) {
        for (let j = i; j < n; j++) {
            dp[i][j] = (j - i < 2 || dp[i+1][j-1]) && s[i] == s[j]
            if (ans.length < j - i + 1 && dp[i][j]) {
                ans = s.slice(i, j+1)
            }
        }   
    }

    return ans
};

longestPalindrome("aaaa")