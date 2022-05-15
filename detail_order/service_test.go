package detail

import (
	"github.com/stretchr/testify/mock"
)

var repoDetail = &RepositoryMock{Mock: mock.Mock{}}
var serviceDetail = detailOrderService{repository: repoDetail}

// func TestSaveDetailOrder(t *testing.T) {
// 	var input = []InputNewDetailOrder{
// 		{ItemID: 1, Quantity: 2, Note: "note"},
// 		//{ItemID: 2, Quantity: 1, Note: "note"},
// 	}

// 	var detail []order.DetailOrder
// 	for _, v := range input {
// 		detail = append(detail, order.DetailOrder{OrderID: 1, ItemID: v.ItemID, Quantity: v.Quantity, Note: v.Note})
// 	}

// 	repoDetail.Mock.On("Save", detail).Return(detail)

// 	result, err := serviceDetail.SaveDetailOrder(1, input)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, detail[0])
// 	assert.Equal(t, detail[0].ID, result[0].ID)
// 	assert.Equal(t, detail[0].ItemID, result[0].ItemID)
// }
