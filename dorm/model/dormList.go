package model

import "dorm/utils"

type DormListItem struct {
	BuildingName  string `json:"building_name"`
	AvailableBeds int    `json:"available_beds"`
	Count         int    `json:"count"`
}

func GetDormList(gender string) (dormList []*DormListItem, err error) {
	sqlStr := "select building_name, available_beds, count from dorm_count where gender = ? and available_beds>0"
	rows, err := utils.Db.Query(sqlStr, gender)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := &DormListItem{}
		err = rows.Scan(&item.BuildingName, &item.AvailableBeds, &item.Count)
		if err != nil {
			return nil, err
		}
		dormList = append(dormList, item)
	}
	return dormList, nil
}
