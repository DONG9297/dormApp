package dao

import (
	"dorm/model"
	"dorm/utils"
)

func GetBuildingByName(name string) (building *model.Building, err error) {
	sqlStr := "select building_id, name from buildings where name = ?"
	row := utils.Db.QueryRow(sqlStr, name)
	building = &model.Building{}
	err = row.Scan(&building.ID, &building.Name)
	if err != nil {
		return nil, err
	}
	return building, nil
}

func GetBuildingByID(ID int) (building *model.Building, err error) {
	sqlStr := "select building_id, name from buildings where building_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	building = &model.Building{}
	err = row.Scan(&building.ID, &building.Name)
	if err != nil {
		return nil, err
	}
	return building, nil
}

func GetAllBuildings() (buildings []*model.Building, err error) {
	sql := "select building_id, name from buildings"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		building := &model.Building{}
		err := rows.Scan(&building.ID, &building.Name)
		if err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}
	return buildings, nil
}
