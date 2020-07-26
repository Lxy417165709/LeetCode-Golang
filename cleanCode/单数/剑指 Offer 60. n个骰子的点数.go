package 单数

func twoSum(n int) []float64 {
	dices := NewDices()
	return dices.GetProbabilities(n)
}

// ------------- Dices -------------- (递归超时，采用计次的递归)
type Dices struct {
	probabilities                 []float64
	combinationsOfDicesSum        []int
	countOfProbabilityCombination int
}

func NewDices() *Dices {
	return &Dices{}
}

func (ds *Dices) GetProbabilities(countOfDice int) []float64 {
	ds.formCombinationsOfDicesSum(countOfDice)
	ds.formProbabilities(countOfDice)
	return ds.probabilities[countOfDice : 6*countOfDice+1]
}

func (ds *Dices) formProbabilities(countOfDice int) {
	ds.probabilities = make([]float64, 6*countOfDice+1)
	for i := countOfDice; i <= countOfDice*6; i++ {
		ds.probabilities[i] = float64(ds.combinationsOfDicesSum[i]) / float64(ds.countOfProbabilityCombination)
	}
}

func (ds *Dices) formCombinationsOfDicesSum(countOfDice int) {
	ds.countOfProbabilityCombination = 0
	ds.combinationsOfDicesSum = make([]int, 6*countOfDice+1)
	ds.formCombinationsOfDicesSumThere(countOfDice, 0)
}

func (ds *Dices) formCombinationsOfDicesSumThere(countOfDice int, nowSum int) {
	if countOfDice == 0 {
		ds.combinationsOfDicesSum[nowSum]++
		ds.countOfProbabilityCombination++
		return
	}
	for i := 1; i <= 6; i++ {
		ds.formCombinationsOfDicesSumThere(countOfDice-1, nowSum+i)
	}
}

// ------------- Dices -------------- (动态规划)
type Dices struct {
}

func NewDices() *Dices {
	return &Dices{}
}

func (ds *Dices) GetProbabilities(countOfDice int) []float64 {
	// 2个状态：一个是被骰子的骰子数、一个是骰子们被投掷出的点数和
	// probabilitiesOfDicesSum[i][sum] =
	//		probabilitiesOfDicesSum[i-1][sum-1]*probabilitiesOfDicesSum[1][1] +
	//		probabilitiesOfDicesSum[i-1][sum-2]*probabilitiesOfDicesSum[1][2] +
	//		probabilitiesOfDicesSum[i-1][sum-3]*probabilitiesOfDicesSum[1][3] +
	//		probabilitiesOfDicesSum[i-1][sum-4]*probabilitiesOfDicesSum[1][4] +
	//		probabilitiesOfDicesSum[i-1][sum-5]*probabilitiesOfDicesSum[1][5] +
	//		probabilitiesOfDicesSum[i-1][sum-6]*probabilitiesOfDicesSum[1][6]
	const MinPoint = 1
	const MaxPoint = 6
	const countOfPoint = MaxPoint - MinPoint + 1
	var minSumOfDices = MinPoint * countOfDice
	var maxSumOfDices = MaxPoint * countOfDice
	probabilitiesOfDicesSum := get2DSlice(countOfDice+1, maxSumOfDices+1)
	for point := MinPoint; point <= MaxPoint; point++ {
		probabilitiesOfDicesSum[1][point] = 1.0 / float64(countOfPoint)
	}
	for _countOfDice := 2; _countOfDice <= countOfDice; _countOfDice++ {
		_minSumOfDices := MinPoint * _countOfDice
		_maxSumOfDices := MaxPoint * _countOfDice
		for sumOfDices := _minSumOfDices; sumOfDices <= _maxSumOfDices; sumOfDices++ {
			for point := MinPoint; point <= MaxPoint && sumOfDices-point >= MinPoint; point++ {
				probabilitiesOfDicesSum[_countOfDice][sumOfDices] += probabilitiesOfDicesSum[_countOfDice-1][sumOfDices-point] * probabilitiesOfDicesSum[1][point]
			}
		}
	}
	return probabilitiesOfDicesSum[countOfDice][minSumOfDices : maxSumOfDices+1]
}

func get2DSlice(rows, column int) [][]float64 {
	slice := make([][]float64, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]float64, column)
	}
	return slice
}

/*
	题目链接：https://leetcode-cn.com/submissions/detail/84481360/
	总结
		1. dp 要找出状态。
		2. dp的题目，+1推，+2推都可以。 (意思就是按最后一个骰子所得的点数或最后两个骰子所得的点数都能dp出正确结果)
			(这里保留+1推的结果)
		3. 可以采用策略模式进行重构
*/
