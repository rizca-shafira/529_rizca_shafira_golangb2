package controllers

// import package
import (
	"ass2/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// struct data
type Data struct {
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (idb *InDB) CreateOrder(c *gin.Context) {

	var (
		order  structs.Order
		item   structs.Item
		result gin.H
	)

	customer_name := c.PostForm("customer_name")
	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	stringQty := c.PostForm("quantity")

	quantity, _ := strconv.Atoi(stringQty)

	order.Ordered_At = time.Now()
	order.Customer_Name = customer_name
	item.Item_Code = item_code
	item.Description = description
	item.Quantity = quantity

	var data []Data
	dataStruct := Data{
		Item_Code:   item.Item_Code,
		Description: item.Description,
		Quantity:    quantity,
	}
	data = append(data, dataStruct)

	err := idb.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Create Data Not Successfully",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		item.Order_ID = order.Order_ID

		err = idb.DB.Create(&item).Error
		if err != nil {
			result = gin.H{
				"result": "Create Data Not Successfully",
			}
			c.JSON(http.StatusNoContent, result)
		} else {

			c.JSON(http.StatusOK, gin.H{
				"orderId":       order.Order_ID,
				"orderedAt":     order.Ordered_At,
				"customer_name": order.Customer_Name,
				"items":         data,
			})
		}
	}
}

func (idb *InDB) GetOrders(c *gin.Context) {

	var (
		orders []structs.Order
		result gin.H
	)
	idb.DB.Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
		c.JSON(http.StatusOK, result)
	}
}

// update order
func (idb *InDB) UpdateOrder(c *gin.Context) {

	id := c.Param("orderId")

	type Data2 struct {
		LineItemId  int    `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}

	var (
		order    structs.Order
		item     structs.Item
		newOrder structs.Order
		result   gin.H
	)

	err := idb.DB.Where("id = ?", id).First(&order).Error
	if err != nil {

		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		err = idb.DB.Where("order_id = ?", id).First(&item).Error
		if err != nil {

			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)

		} else {

			orderid, _ := strconv.Atoi(id)
			customer_name := c.PostForm("customer_name")
			newOrder.Order_ID = orderid
			newOrder.Customer_Name = customer_name
			newOrder.Ordered_At = order.Ordered_At

			var data []Data2
			dataStruct := Data2{
				LineItemId:  item.Item_ID,
				ItemCode:    item.Item_Code,
				Description: item.Description,
				Quantity:    item.Quantity,
			}

			data = append(data, dataStruct)

			err = idb.DB.Model(&order).Updates(newOrder).Error
			if err != nil {

				result = gin.H{
					"result": "Update Failed",
				}
				c.JSON(http.StatusNoContent, result)
			} else {

				result = gin.H{
					"result":        "successfully updated data",
					"orderId":       order.Order_ID,
					"orderedAt":     order.Ordered_At,
					"customer_name": order.Customer_Name,
					"data":          data,
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}
}

// delete data with {id}
func (idb *InDB) DeleteOrder(c *gin.Context) {

	var (
		order  structs.Order
		item   structs.Item
		result gin.H
	)

	orderId := c.Param("orderId")

	err := idb.DB.Where("id = ?", orderId).First(&order).Error
	if err != nil {

		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		err = idb.DB.Where("order_id = ?", orderId).First(&item).Error
		if err != nil {

			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)
		} else {

			err = idb.DB.Delete(&item).Error
			if err != nil {

				result = gin.H{
					"result": "Delete item Failed",
				}
				c.JSON(http.StatusBadRequest, result)
			} else {

				err = idb.DB.Delete(&order).Error
				if err != nil {

					result = gin.H{
						"result": "Delete order Failed",
					}
					c.JSON(http.StatusBadRequest, result)
				} else {

					result = gin.H{
						"result": "data deleted successfully",
						"data":   order,
					}
					c.JSON(http.StatusOK, result)
				}
			}
		}
	}
}
