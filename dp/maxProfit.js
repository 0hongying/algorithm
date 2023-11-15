// 买股票的最佳时机（只能一次买卖）
function maxProfit(prices) {
    let min = prices[0]
    let max = 0
    for (let i = 1; i < prices.length; i++) {
        // 假设第i天卖出，寻找前i天的最低点
        min = Math.min(min, prices[i])
        max = Math.max(max, prices[i]-min)
    }
    return max
}

const prices = [];
const ret = maxProfit(prices);
console.log(ret);

// 买彩票的最佳时机（可以多次买卖）
function maxProfit2(prices) {
    let ret = 0
    for (let index = 1; index < prices.length; index++) {
        // 把所有上坡的
        if (prices[index] > prices[index -1]) {
            ret +=(prices[index] - prices[index-1])
        }
    }
    return ret
}

const prices2 = [];
const ret2 = maxProfit2(prices2);
console.log(ret2);