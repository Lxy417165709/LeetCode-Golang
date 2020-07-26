package 数学

import "math"

const circleDegree = 360.0
const countOfMinuteInOneHour = 60
const constOfHourInOneCircle = 12
const oneHourDegree = circleDegree / constOfHourInOneCircle
const oneMinuteDegree = circleDegree / countOfMinuteInOneHour

func angleClock(hour int, minutes int) float64 {
	minuteDegree := float64(minutes * oneMinuteDegree)
	hourDegree := float64((hour % constOfHourInOneCircle) * oneHourDegree) + oneHourDegree * (float64(minutes) / countOfMinuteInOneHour)
	degreeDiff := math.Abs(hourDegree - minuteDegree)
	return math.Min(degreeDiff,circleDegree - degreeDiff)
}
