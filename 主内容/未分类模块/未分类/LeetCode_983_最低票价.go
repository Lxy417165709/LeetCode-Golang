func mincostTickets(days []int, costs []int) int {
    const INF = 1000000000
    dp := make([]int,366)
    
    isExist := make(map[int]bool)
    for i:=0;i<len(days);i++{
        isExist[days[i]]=true
    }

    for day:=1;day<=365;day++{
        if isExist[day]==false{
            dp[day] = dp[day-1]
            continue
        }
        dp[day] = INF
        if day>=1{
            dp[day] = min(dp[day],dp[day-1] + costs[0])
        }else{
            dp[day] = min(dp[day],costs[0])
        }
        if day>=7{
            dp[day] = min(dp[day],dp[day-7] + costs[1])
        }else{
            dp[day] = min(dp[day],costs[1])
        }
        if day>=30{
            dp[day] = min(dp[day],dp[day-30] + costs[2])
        }else{
            dp[day] = min(dp[day],costs[2])
        }
    }
    return dp[365]
}



func min(a,b int) int{
    if a>b{
        return b
    }
    return a
}

/*
	总结：
		1. 第一次看题的时候以为是区间的题，然后没想出，看了解答后发现是动态规划的..
		2. 本例的DP数组意义是：旅游到第 i 天的最少花费。
*/