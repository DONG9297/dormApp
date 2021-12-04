package dao

import (
	"dorm/model"
	"dorm/utils"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	sql := "insert into order_items(order_id, user_id) values(?,?)"
	_, err := utils.Db.Exec(sql, orderItem.OrderID, orderItem.UserID)
	return err
}

func GetItemsByOrderID(orderID string) (orderItems []*model.OrderItem, err error) {
	sqlStr := "select item_id, order_id, user_id from order_items where order_id = ?"
	rows, err := utils.Db.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := &model.OrderItem{}
		err = rows.Scan(&item.ID, &item.OrderID, &item.UserID)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, item)
	}
	return orderItems, nil
}
func GetOrderIDsByUserID(userID string) (orderIDs []string, err error) {
	sqlStr := "select  order_id from order_items where user_id = ?"
	rows, err := utils.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var orderID string
		err = rows.Scan(&orderID)
		if err != nil {
			return nil, err
		}
		orderIDs = append(orderIDs, orderID)
	}
	return orderIDs, nil
}
