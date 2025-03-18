package configuration

import (
	"os"
)

var (
	JOB_QUEUE_DOMPDF_URL string
	LAMBDA_AWS_REGION    string
	JOB_QUEUE_FPDF_URL   string
	JOB_QUEUE_TCPDF_URL  string
	TARGETS_FILE_PATH    string
)

func LoadConfig() {
	JOB_QUEUE_DOMPDF_URL = getEnv("JOB_QUEUE_DOMPDF_URL", "")
	JOB_QUEUE_FPDF_URL = getEnv("JOB_QUEUE_FPDF_URL", "")
	JOB_QUEUE_TCPDF_URL = getEnv("JOB_QUEUE_TCPDF_URL", "")
	LAMBDA_AWS_REGION = getEnv("LAMBDA_AWS_REGION", "us-east-1")
	TARGETS_FILE_PATH = getEnv("TARGETS_FILE_PATH", "")
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
