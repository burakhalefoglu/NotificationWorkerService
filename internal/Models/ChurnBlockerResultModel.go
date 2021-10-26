package Models

type ChurnBlockerResultModel struct {
	ClientId                string
	ProjectId               string
	CenterOfDifficultyLevel int
	RangeCount              int
}

type ChurnBlockerResultDto struct {
	CenterOfDifficultyLevel int
	RangeCount              int
}

