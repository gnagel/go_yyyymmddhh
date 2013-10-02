package time_utils

import "fmt"
import "time"

func NowUTC() *time.Time {
	t := time.Now().UTC()
	return &t
}

func YYYYMMDDHH() string {
	utc := NowUTC()
	return fmt.Sprintf("%04d%02d%02d%02d", utc.Year(), utc.Month(), utc.Day(), utc.Hour())
}

func YYYYMMDD() string {
	utc := NowUTC()
	return fmt.Sprintf("%04d%02d%02d", utc.Year(), utc.Month(), utc.Day())
}

func YYYYMM() string {
	utc := NowUTC()
	return fmt.Sprintf("%04d%02d", utc.Year(), utc.Month())
}

func YYYY() string {
	utc := NowUTC()
	return fmt.Sprintf("%04d", utc.Year())
}
