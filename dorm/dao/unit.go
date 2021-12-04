package dao

import (
	"dorm/model"
	"dorm/utils"
)

func GetUnitsByBuilding(BuildingID int) (units []*model.Unit, err error) {
	sqlStr := "select unit_id, name, building_id from units where building_id = ?"
	rows, err := utils.Db.Query(sqlStr, BuildingID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		unit := &model.Unit{}
		err = rows.Scan(&unit.ID, &unit.Name, &unit.BuildingID)
		if err != nil {
			return nil, err
		}
		units = append(units, unit)
	}
	return units, err
}

func GetUnitByID(ID int) (unit *model.Unit, err error) {
	sqlStr := "select unit_id, name, building_id from units where unit_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	unit = &model.Unit{}
	err = row.Scan(&unit.ID, &unit.Name, &unit.BuildingID)
	if err != nil {
		return nil, err
	}
	return unit, nil
}
