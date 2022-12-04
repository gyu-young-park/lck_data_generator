package calibratordate

type DateCalibrator interface {
	Calibrate(string, string) string
}
