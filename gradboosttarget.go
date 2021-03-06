package CloudForest

import ()

/*
GradBoostTarget wraps a numerical feature as a target for us in Adaptive Boosting (AdaBoost)
*/
type GradBoostTarget struct {
	NumFeature
	LearnRate float64
}

//BUG(ryan) does GradBoostingTarget need seperate residuals and values?
func (f *GradBoostTarget) Boost(leaves *[][]int) (weight float64) {
	for _, cases := range *leaves {
		f.Update(&cases)
	}
	return f.LearnRate

}

//Update updates the underlying numeric data by subtracting the mean*weight of the
//specified cases from the value for those cases.
func (f *GradBoostTarget) Update(cases *[]int) {
	m := f.Predicted(cases)
	for _, i := range *cases {
		if !f.IsMissing(i) {
			f.Put(i, f.Get(i)-f.LearnRate*m)
		}
	}
}
