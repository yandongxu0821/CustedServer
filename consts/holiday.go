package consts

import "github.com/LollipopKit/custed-server/model"

var (
	// TODO: 2025节假日信息，上次更新 2025-08-26
	Holidays = []model.HolidaySkip{
		{ // 元旦
			From: &model.YearMonthDay{2025, 1, 1},
			To:   &model.YearMonthDay{2025, 1, 1},
		},
		{ // 清明
			From: &model.YearMonthDay{2025, 4, 4},
			To:   &model.YearMonthDay{2025, 4, 6},
		},
		{ // 劳动节
			From: &model.YearMonthDay{2025, 5, 1},
			To:   &model.YearMonthDay{2025, 5, 5},
		},
		{ // 端午
			From: &model.YearMonthDay{2025, 5, 31},
			To:   &model.YearMonthDay{2025, 6, 2},
		},
		{ // 国庆+中秋
			From: &model.YearMonthDay{2025, 10, 1},
			To:   &model.YearMonthDay{2025, 10, 8},
		},
	}
)
