package dao

import (
	"dorm/model"
	"dorm/utils"
)

func AddOrder(order *model.Order) error {
	sql := "insert into orders(order_id, user_id, count, building_id, gender, create_time, state) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, order.ID, order.UserID, order.Count, order.BuildingID, order.Gender, order.CreateTime, order.State)
	if err != nil {
		return err
	}
	return nil
}

func GetAllOrders() (orders []*model.Order, err error) {
	sql := "select order_id, user_id, count, building_id, gender, create_time, state from orders"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Count, &order.BuildingID, &order.Gender, &order.CreateTime, &order.State)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
func GetUnprocessedOrdersBefore(timeStr string) (orders []*model.Order, err error) {
	sql := "select order_id, user_id, count, building_id, gender, create_time, state from orders where create_time<? and state = 0 ORDER BY  create_time ASC"
	rows, err := utils.Db.Query(sql, timeStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Count, &order.BuildingID, &order.Gender, &order.CreateTime, &order.State)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func GetOrdersByUserID(userID int) (orders []*model.Order, err error) {
	sql := "select order_id, user_id, count, building_id, gender, create_time, state from orders where user_id = ?"
	rows, err := utils.Db.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Count, &order.BuildingID, &order.Gender, &order.CreateTime, &order.State)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func UpdateOrderState(orderID string, state int) error {
	//写sql语句
	sql := "update orders set state = ? where order_id = ?"
	//执行
	_, err := utils.Db.Exec(sql, state, orderID)
	return err
}
