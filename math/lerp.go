package math

// Lerp performs a linear interpolation between two values, given the progress between them in the range [0,1].
func Lerp(v1, v2 float64, progress float64) float64 {
	return v1 + (v2-v1)*progress
}

// Lerp2D performs linear interpolation on both X and Y co-ordinates, given the progress between them in the
// range [0,1].
func Lerp2D(x1, y1 float64, x2, y2 float64, progress float64) (x, y float64) {
	return Lerp(x1, x2, progress), Lerp(y1, y2, progress)
}

// EaseIn is an easing function that slows progress at the start.
func EaseIn(progress float64) float64 {
	return progress * progress
}

// EaseOut is an easing function that slows progress at the end.
func EaseOut(progress float64) float64 {
	return 1 - ((1 - progress) * (1 - progress))
}

// EaseInOut is an easing function that slows progress at the start and end.
func EaseInOut(progress float64) float64 {
	return Lerp(EaseIn(progress), EaseOut(progress), progress)
}

// EaseOutBounce is an easing function that rapidly approaches the end and then bounces back, gradually settling in
// to place. Using the algorithm described at https://easings.net/#easeOutBounce.
func EaseOutBounce(progress float64) float64 {
	const n1 = 7.5625
	const d1 = 1 / 2.75

	if progress < d1 {
		return n1 * progress * progress
	} else if progress < 2*d1 {
		progress -= 1.5 * d1
		return n1*progress*progress + 0.75
	} else if progress < 2.5*d1 {
		progress -= 2.25 * d1
		return n1*progress*progress + 0.9375
	} else {
		progress -= 2.625 * d1
		return n1*progress*progress + 0.984375
	}
}
