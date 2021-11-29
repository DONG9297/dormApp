package model

import "dorm/utils"

type Building struct {
	ID   int
	Name string
}

func GetBuildingByName(name string) (building *Building, err error) {
	sqlStr := "select building_id, name from buildings where name = ?"
	row := utils.Db.QueryRow(sqlStr, name)
	building = &Building{}
	err = row.Scan(&building.ID, &building.Name)
	if err != nil {
		return nil, err
	}
	return building, nil
}

func GetBuildingByID(ID int) (building *Building, err error) {
	sqlStr := "select building_id, name from buildings where building_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	building = &Building{}
	err = row.Scan(&building.ID, &building.Name)
	if err != nil {
		return nil, err
	}
	return building, nil
}

func GetAllBuildings() (buildings []*Building, err error) {
	sql := "select building_id, name from buildings"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		building := &Building{}
		err := rows.Scan(&building.ID, &building.Name)
		if err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}
	return buildings, nil
}
