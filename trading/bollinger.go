package trading

import (
	"math"
	"math/big"
)

// Calculate the Bollinger index
func CalculateBollingerBands(closes []*big.Float, stdDevMultiplier *big.Float, stdDevWindowSize int) ([]*big.Float, []*big.Float, []*big.Float) {
	var (
		upperBand  []*big.Float
		middleBand []*big.Float
		lowerBand  []*big.Float
	)

	period := len(closes)

	for i := stdDevWindowSize; i < period; i++ {
		window := closes[i-stdDevWindowSize : i]
		mean := calculateMean(window)
		middle := mean
		stdDev := calculateStandardDeviation(window, mean)
		variable := (new(big.Float)).Mul(stdDevMultiplier, stdDev)
		upper := (new(big.Float)).Add(mean, variable)
		lower := (new(big.Float)).Sub(mean, variable)

		upperBand = append(upperBand, upper)
		middleBand = append(middleBand, middle)
		lowerBand = append(lowerBand, lower)
	}

	return upperBand, middleBand, lowerBand
}

// Calculated mean
func calculateMean(values []*big.Float) *big.Float {
	sum := big.NewFloat(0)
	for _, value := range values {
		sum.Add(sum, value)
	}
	return sum.Quo(sum, big.NewFloat(float64(len(values))))
}

// Calculated standard deviation
func calculateStandardDeviation(values []*big.Float, mean *big.Float) *big.Float {
	sum := big.NewFloat(0)
	for _, value := range values {
		diff := (new(big.Float)).Sub(value, mean)
		sum.Add(sum, diff.Mul(diff, diff))
	}
	variance := sum.Quo(sum, big.NewFloat(float64(len(values))))
	v, _ := variance.Float64()
	return big.NewFloat(math.Sqrt(v))
}
