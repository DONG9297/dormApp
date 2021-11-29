package model

import "dorm/utils"

type Unit struct {
	ID         int
	Name       string
	BuildingID int
}

func GetUnitsByBuilding(BuildingID int) (units []*Unit, err error) {
	sqlStr := "select unit_id, name, building_id from units where building_id = ?"
	rows, err := utils.Db.Query(sqlStr, BuildingID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		unit := &Unit{}
		err = rows.Scan(&unit.ID, &unit.Name, &unit.BuildingID)
		if err != nil {
			return nil, err
		}
		units = append(units, unit)
	}
	return units, err
}

func GetUnitByID(ID int) (unit *Unit, err error) {
	sqlStr := "select unit_id, name, building_id from units where unit_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	unit = &Unit{}
	err = row.Scan(&unit.ID, &unit.Name, &unit.BuildingID)
	if err != nil {
		return nil, err
	}
	return unit, nil
}
