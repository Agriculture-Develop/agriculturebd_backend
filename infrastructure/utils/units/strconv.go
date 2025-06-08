package units

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

// Duration
/**
 * @Description: 时间转化工具 ,支持 1s 1m 1h 1d
 * @param d
 * @return time.Duration , error
 */
func Duration(d string) time.Duration {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr
	}

	// 解析day
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)

		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr
		}
		return dr + ndr
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	zap.L().Error("Error parsing duration:", zap.Error(err))

	return time.Duration(dv)
}

// DurationWithErr
/**
 * @Description: 时间转化工具 ,支持 1s 1m 1h 1d
 * @param d
 * @return time.Duration , error
 */
func DurationWithErr(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}

	// 解析day
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)

		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)

	return time.Duration(dv), fmt.Errorf("error parsing duration:%v", err)
}
