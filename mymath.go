package mymath

func Sqrt(x float64) float64 {
	const eps floa64 = 1e-9;
    var x float64 = 1;
    while (abs(x * x - n) > eps) {
        x = (x + n / x) / 2;
    }
	return x
}

func Max(x, y float64) float64 {
	var diff float64 = x - y
	if diff < 0 {
		return y
	}
	return x
}

func Min(x, y float64) float64 {
	var diff float64 = x - y
	if diff < 0 {
		return x
	}
	return y
}

func Pow(x, y float64) float64 {
		
}

func Floor(x float64) float64 {
	return float64(int64(x))
}

func Ceil(x float64) float64 {
	return float64(int64(x) + 1)
}



