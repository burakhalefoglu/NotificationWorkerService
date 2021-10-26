package Models


type ChurnPredictionResultModel struct {
	ClientId                string
	ProjectId               string
	CenterOfDifficultyLevel int
	RangeCount              int
}

type ChurnPredictionResultDto struct {
	CenterOfDifficultyLevel int
	RangeCount              int
}
