package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type SparseColorMethod int

const (
	INTERPOLATE_UNDEFINED_COLOR   SparseColorMethod = C.UndefinedColorInterpolate
	INTERPOLATE_BARYCENTRIC_COLOR SparseColorMethod = C.BarycentricColorInterpolate
	INTERPOLATE_BILINEAR_COLOR    SparseColorMethod = C.BilinearColorInterpolate
	INTERPOLATE_POLYNOMIAL_COLOR  SparseColorMethod = C.PolynomialColorInterpolate
	INTERPOLATE_SHEPARDS_COLOR    SparseColorMethod = C.ShepardsColorInterpolate
	INTERPOLATE_VORONOI_COLOR     SparseColorMethod = C.VoronoiColorInterpolate
	INTERPOLATE_INVERSE_COLOR     SparseColorMethod = C.InverseColorInterpolate
)