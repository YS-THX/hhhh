// Package holidays supply holidays query.
package holidays

import (
	"errors"
	"fmt"
	"time"
)

var b *book

func init() {
	events, err := loadData()
	if err != nil {
		fmt.Println(err)
	}

	bk, err := newBookfromEvents(events)
	if err != nil {
		fmt.Println(err)
	} else {
		b = &bk
	}
}

// IsHoliday checks given date is holiday or not.
func IsHoliday(date time.Time) (bool, error) {
	err := checkInitBook()
	if err != nil {
		return false, err
	}

	return b.isHoliday(date), nil
}

// IsWorkingDay checks given date is working day or not.
func IsWorkingDay(date time.Time) (bool, error) {
	err := checkInitBook()
	if err != nil {
		return false, err
	}

	return b.isWorkingday(date), nil
}

// checkInitBook check init book
func checkInitBook() error {
	if b == nil {
		return fmt.Errorf("book initialize failed")
	}
	return nil
}

// GetTNthWorkingDay Get the Nth working day--(time params)
// eg GetTNthWorkingDay("2021-09-30 00:00:00",3) ==> 2021-10-11 00:00:00
func GetTNthWorkingDay(date time.Time, Nth int32) (time.Time, error) {
	if Nth <= 0 {
		return date, errors.New("not support such Nth")
	}
	// Countdown days、the number of days to be added counter（倒计时天数、需新增的天数计数器）
	countNth, needAddNth := int(0), int(0)
	for {
		if int32(countNth) == Nth {
			break
		}
		needAddNth++
		isWorkday := b.isWorkingday(date.AddDate(0, 0, needAddNth))
		if isWorkday {
			countNth++
		}
	}
	return date.AddDate(0, 0, needAddNth), nil
}

// GetSNthWorkingDay get nth working day-(string params)
// eg GetSNthWorkingDay("2021-09-30 00:00:00",3) ==> 2021-10-11 00:00:00
func GetSNthWorkingDay(timeStr string, Nth int32) (time.Time, error) {
	if timeStr == "" {
		return time.Now(), errors.New("not support such timeStr and return now")
	}
	// str ==> time
	date, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return time.Now(), errors.New(err.Error() + ": " + "not support such timeStr and return now")
	}
	return GetTNthWorkingDay(date, Nth)
}
