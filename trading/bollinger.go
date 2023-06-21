package trading

import "math"

// Calculate the Bollinger index
func CalculateBollingerBands(closes []float64, stdDevMultiplier float64, stdDevWindowSize int) ([]float64, []float64, []float64) {
	var (
		upperBand  []float64
		middleBand []float64
		lowerBand  []float64
	)

	period := len(closes)
	// stdDevMultiplier := 2.0
	// stdDevWindowSize := 20

	for i := stdDevWindowSize; i < period; i++ {
		window := closes[i-stdDevWindowSize : i]
		mean := calculateMean(window)
		stdDev := calculateStandardDeviation(window, mean)
		upper := mean + (stdDevMultiplier * stdDev)
		middle := mean
		lower := mean - (stdDevMultiplier * stdDev)

		upperBand = append(upperBand, upper)
		middleBand = append(middleBand, middle)
		lowerBand = append(lowerBand, lower)
	}

	return upperBand, middleBand, lowerBand
}

// Calculated mean
func calculateMean(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

// Calculated standard deviation
func calculateStandardDeviation(values []float64, mean float64) float64 {
	sum := 0.0
	for _, value := range values {
		diff := value - mean
		sum += diff * diff
	}
	variance := sum / float64(len(values))
	return math.Sqrt(variance)
}