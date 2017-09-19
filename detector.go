package triangulator

import (
	"image"
	"math/rand"
	"time"
)

const (
	POINT_RATE = 0.075
	POINT_MAX_NUM = 3500
)

func GetEdgePoints(img *image.NRGBA, threshold int)[]point {
	rand.Seed(time.Now().UTC().UnixNano())
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y

	var (
		points []point
		sum, total int
		x, y, row, col, sx, sy, step int
		dpoints []point
	)

	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			sum, total = 0, 0

			for row = -1; row <= 1; row++ {
				sy = y + row
				step = sy * width

				if sy >= 0 && sy < height {
					for col = -1; col <= 1; col++ {
						sx = x + col
						if sx >= 0 && sx < width {
							sum += int(img.Pix[(sx + step) << 2])
							total++
						}
					}
				}
			}
			if total > 0 {
				sum /= total
			}
			if sum > threshold {
				points = append(points, point{x: x, y: y})
			}
		}
	}
	ilen := len(points)
	tlen := ilen
	limit := int(float64(ilen) * POINT_RATE)

	if limit > POINT_MAX_NUM {
		limit = POINT_MAX_NUM
	}

	for i := 0; i < limit && i < ilen; i++ {
		j := int(float64(tlen) * rand.Float64())
		dpoints = append(dpoints, points[j])
		// Remove points
		points = append(points[:j], points[j+1:]...)
		tlen--
	}

	return dpoints
}