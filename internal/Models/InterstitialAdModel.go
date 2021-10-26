package Models

type InterstitialAdModel struct {
	ClientIdList           []string
	ProjectId              string
	IsAdvSettingsActive    bool
	AdvFrequencyStrategies map[string]int
}

type InterstitialAdDto struct {
	IsAdvSettingsActive    bool
	AdvFrequencyStrategies map[string]int
}
