package resolver

//
//import (
//"context"
//"testing"
//
//"github.com/dictyBase/go-genproto/dictybaseapis/order"
//
//"github.com/dictyBase/graphql-server/internal/graphql/mocks"
//"github.com/dictyBase/graphql-server/internal/graphql/models"
//"github.com/stretchr/testify/assert"
//)
//
//func TestOrder(t *testing.T) {
//	t.Parallel()
//	assert := assert.New(t)
//	ord := &QueryResolver{
//		Registry: &mocks.MockRegistry{},
//		Logger:   mocks.TestLogger(),
//	}
//	id := "999"
//	o, err := ord.Order(context.Background(), id)
//	assert.NoError(err, "expect no error from getting order information")
//	assert.Exactly(o.Data.Id, id, "should match id")
//	assert.Exactly(o.Data.Attributes.Courier, mocks.MockOrderAttributes.Courier, "should match courier")
//	assert.Exactly(o.Data.Attributes.CourierAccount, mocks.MockOrderAttributes.CourierAccount, "should match courier account")
//	assert.Exactly(o.Data.Attributes.Comments, mocks.MockOrderAttributes.Comments, "should match comments")
//	assert.Exactly(o.Data.Attributes.Payment, mocks.MockOrderAttributes.Payment, "should match payment")
//	assert.Exactly(o.Data.Attributes.PurchaseOrderNum, mocks.MockOrderAttributes.PurchaseOrderNum, "should match purchase order number")
//	assert.Exactly(o.Data.Attributes.Status, order.OrderStatus_In_preparation, "should match status")
//	assert.Exactly(o.Data.Attributes.Consumer, mocks.MockOrderAttributes.Consumer, "should match consumer")
//	assert.Exactly(o.Data.Attributes.Payer, mocks.MockOrderAttributes.Payer, "should match payer")
//	assert.Exactly(o.Data.Attributes.Purchaser, mocks.MockOrderAttributes.Purchaser, "should match purchaser")
//	assert.ElementsMatch(o.Data.Attributes.Items, mocks.MockOrderAttributes.Items, "should match items")
//}
//
//func TestListOrders(t *testing.T) {
//	t.Parallel()
//	assert := assert.New(t)
//	ord := &QueryResolver{
//		Registry: &mocks.MockRegistry{},
//		Logger:   mocks.TestLogger(),
//	}
//	cursor := 0
//	limit := 10
//	filter := "type===strain"
//	o, err := ord.ListOrders(context.Background(), &cursor, &limit, &filter)
//	assert.NoError(err, "expect no error from getting list of orders")
//	assert.Exactly(o.Limit, &limit, "should match limit")
//	assert.Exactly(o.PreviousCursor, 0, "should match previous cursor")
//	assert.Exactly(o.NextCursor, 10000, "should match next cursor")
//	assert.Exactly(o.TotalCount, 3, "should match total count (length) of items")
//	assert.Len(o.Orders, 3, "should have three orders")
//}
//
//func TestCreateOrder(t *testing.T) {
//	t.Parallel()
//	assert := assert.New(t)
//	ord := &MutationResolver{
//		Registry: &mocks.MockRegistry{},
//		Logger:   mocks.TestLogger(),
//	}
//	comments := "first order"
//	pon := "987654"
//	id := "DBS123456"
//	input := &models.CreateOrderInput{
//		Courier:          "USPS",
//		CourierAccount:   "123456",
//		Comments:         &comments,
//		Payment:          "credit",
//		PurchaseOrderNum: &pon,
//		Status:           models.StatusEnumInPreparation,
//		Consumer:         "art@vandelayindustries.com",
//		Payer:            "george@costanza.com",
//		Purchaser:        "thatsgold@jerry.org",
//		Items:            []*string{&id},
//	}
//	o, err := ord.CreateOrder(context.Background(), input)
//	assert.NoError(err, "expect no error from creating an order")
//	assert.Exactly(o.Data.Attributes.Courier, input.Courier, "should match courier")
//	assert.Exactly(o.Data.Attributes.CourierAccount, input.CourierAccount, "should match courier account")
//	assert.Exactly(&o.Data.Attributes.Comments, input.Comments, "should match comments")
//	assert.Exactly(o.Data.Attributes.Payment, input.Payment, "should match payment")
//	assert.Exactly(&o.Data.Attributes.PurchaseOrderNum, input.PurchaseOrderNum, "should match purchase order number")
//	assert.Exactly(o.Data.Attributes.Status, order.OrderStatus_In_preparation, "should match status")
//	assert.Exactly(o.Data.Attributes.Consumer, input.Consumer, "should match consumer")
//	assert.Exactly(o.Data.Attributes.Payer, input.Payer, "should match payer")
//	assert.Exactly(o.Data.Attributes.Purchaser, input.Purchaser, "should match purchaser")
//	assert.ElementsMatch(o.Data.Attributes.Items, []string{"DBS123456"}, "should match items")
//}
//
//func TestUpdateOrder(t *testing.T) {
//	t.Parallel()
//	assert := assert.New(t)
//	ord := &MutationResolver{
//		Registry: &mocks.MockRegistry{},
//		Logger:   mocks.TestLogger(),
//	}
//	courier := "FedEx"
//	courierAccount := "444444"
//	comments := "Please send ASAP"
//	status := models.StatusEnumGrowing
//	o, err := ord.UpdateOrder(
//		context.Background(),
//		"999",
//		&models.UpdateOrderInput{
//			Courier:        &courier,
//			CourierAccount: &courierAccount,
//			Comments:       &comments,
//			Status:         &status,
//		},
//	)
//	assert.NoError(err, "expect no error from updating an order")
//	assert.Exactly(o.Data.Attributes.Courier, courier, "should match updated courier")
//	assert.Exactly(o.Data.Attributes.CourierAccount, courierAccount, "should match updated courier account")
//	assert.Exactly(o.Data.Attributes.Comments, comments, "should match updated comments")
//	assert.Exactly(o.Data.Attributes.PurchaseOrderNum, "987654", "should match purchase order number")
//	assert.Exactly(o.Data.Attributes.Status, order.OrderStatus_Growing, "should match updated status")
//	assert.Exactly(o.Data.Attributes.Payment, mocks.MockOrderAttributes.Payment, "should match existing payment")
//	assert.Exactly(o.Data.Attributes.Consumer, mocks.MockOrderAttributes.Consumer, "should match existing consumer")
//	assert.Exactly(o.Data.Attributes.Payer, mocks.MockOrderAttributes.Payer, "should match existing payer")
//	assert.Exactly(o.Data.Attributes.Purchaser, mocks.MockOrderAttributes.Purchaser, "should match existing purchaser")
//	assert.ElementsMatch(o.Data.Attributes.Items, mocks.MockOrderAttributes.Items, "should match existing items")
//}
