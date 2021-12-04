package dao

import (
	"dorm/model"
	"dorm/utils"
)

func GetDormList(gender string) (dormList []*model.DormListItem, err error) {
	sqlStr := "select building_name, available_beds, count from dorm_count where gender = ? and available_beds>0"
	rows, err := utils.Db.Query(sqlStr, gender)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := &model.DormListItem{}
		err = rows.Scan(&item.BuildingName, &item.AvailableBeds, &item.Count)
		if err != nil {
			return nil, err
		}
		dormList = append(dormList, item)
	}
	return dormList, nil
}

func GetBuildingList(gender string) (buildingList []string, err error) {
	sqlStr := "select distinct building_name from dorm_count where gender = ? and available_beds>0"
	rows, err := utils.Db.Query(sqlStr, gender)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			return nil, err
		}
		buildingList = append(buildingList, item)
	}
	return buildingList, nil
}
