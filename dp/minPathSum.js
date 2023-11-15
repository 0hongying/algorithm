var minPathSum = function(grid) {
    const m = grid.length
    const n = grid[0].length
    const dp = new Array(m).fill(0).map(() => new Array(n).fill(0))

    dp[0][0] = grid[0][0]


    for (let i = 1; i < grid.length; i++) {
        dp[i][0] = grid[i][0] + dp[i-1][0]
    }

    for (let i = 1; i < grid[0].length; i++) {
        dp[0][i] = grid[0][i] + dp[0][i-1]
    }


    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[0].length; j++) {
            if (i > 0 && j > 0) {
                dp[i][j] = Math.min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
            }
        }
    }

    return dp[grid.length-1][grid[0].length-1]
};

minPathSum([[1,3,1],[1,5,1],[4,2,1]])