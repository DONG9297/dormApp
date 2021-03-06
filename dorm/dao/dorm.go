package dao

import (
	"dorm/model"
	"dorm/utils"
)

func GetDormByID(ID int) (dorm *model.Dorm, err error) {
	sqlStr := "select dorm_id, name, gender, total_beds, available_beds, unit_id from dorms where dorm_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	dorm = &model.Dorm{}
	err = row.Scan(&dorm.ID, &dorm.Name, &dorm.Gender, &dorm.TotalBeds, &dorm.AvailableBeds, &dorm.UnitID)
	if err != nil {
		return nil, err
	}
	return dorm, nil
}

func UpdateDormAvailableBeds(dormID, availableBeds int) (err error) {
	sqlStr := "update dorms set available_beds = ? where dorm_id = ?"
	_, err = utils.Db.Exec(sqlStr, availableBeds, dormID)
	if err != nil {
		return err
	}
	return nil
}

func GetAvailableDormInfos(availableBeds int, buildingName, gender string) (dormInfos []*model.DormInfo, err error) {
	sqlStr := "select building_name, unit_name, dorm_name, dorm_id, available_beds from dorm_list where building_name = ? and gender = ? and available_beds >= ?"
	rows, err := utils.Db.Query(sqlStr, buildingName, gender, availableBeds)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		dormInfo := &model.DormInfo{}
		err = rows.Scan(&dormInfo.BuildingName, &dormInfo.UnitName, &dormInfo.DormName, &dormInfo.DormID, &dormInfo.AvailableBeds)
		if err != nil {
			return nil, err
		}
		dormInfos = append(dormInfos, dormInfo)
	}
	return dormInfos, nil
}

func GetDormInfoByDormID(ID int) (*model.DormInfo, error) {
	sqlStr := "select building_name, unit_name, dorm_name, dorm_id from dorm_list where dorm_id =?"
	row := utils.Db.QueryRow(sqlStr, ID)
	dormInfo := &model.DormInfo{}
	err := row.Scan(&dormInfo.BuildingName, &dormInfo.UnitName, &dormInfo.DormName, &dormInfo.DormID)
	if err != nil {
		return nil, err
	}
	return dormInfo, nil
}
