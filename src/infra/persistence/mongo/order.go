package mongo

import (
	"context"
	"log"

	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	repos "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/repositories"
	order_vo "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"
	infra_model "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderDocumentRepository struct {
	collection *mongo.Collection
}

func NewOrderDocumentRepository(db *mongo.Client) *OrderDocumentRepository {
	return &OrderDocumentRepository{
		collection: db.Database("orders").Collection("order_models"),
	}
}

func ToDocument(order entities.Order) (*infra_model.OrderDocument, error) {
	var items []infra_model.OrderItemDocument
	for _, row := range order.GetItems() {
		or, _ := ItemToDocument(row)
		items = append(items, *or)
	}

	return &infra_model.OrderDocument{
		ID:        order.GetID().String(),
		Items:     items,
		Status:    order.GetStatus().EnumIndex(),
		CreatedAt: order.GetCreatedAt(),
		UpdatedAt: order.GetUpdatedAt(),
	}, nil
}

func ItemToDocument(orderItem order_vo.OrderItem) (*infra_model.OrderItemDocument, error) {
	return &infra_model.OrderItemDocument{
		Name:  orderItem.GetName(),
		Qty:   orderItem.GetQty(),
		Price: orderItem.GetPrice(),
		Total: orderItem.GetTotalPrice(),
	}, nil
}

// implement interface
var _ repos.OrderRepository = &OrderDocumentRepository{}

func (r *OrderDocumentRepository) Persist(context context.Context, order entities.Order) (*entities.Order, error) {
	var row *infra_model.OrderDocument

	or, err := ToDocument(order)
	if err != nil {
		return nil, err
	}

	r.collection.FindOne(context, bson.M{"id": order.GetID().String()}).Decode(&row)
	if row != nil {
		r.collection.FindOneAndReplace(context, bson.M{"id": order.GetID().String()}, or)
	} else {
		r.collection.InsertOne(context, or)
	}

	return &order, nil
}

func (r *OrderDocumentRepository) Detail(context context.Context, ID string) (*entities.Order, error) {
	var row *infra_model.OrderDocument

	res := r.collection.FindOne(context, bson.M{"id": ID})
	if res.Err() != nil {
		return nil, res.Err()
	}

	res.Decode(&row)

	order, err := row.ToEntity()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderDocumentRepository) Delete(context context.Context, ID string) error {
	orderDocument := r.collection.FindOneAndDelete(context, bson.M{"id": ID}).Err()

	if orderDocument != nil {
		return orderDocument
	}

	return nil
}

func (r *OrderDocumentRepository) List(context context.Context, page int, perPage int, filter *repos.OrderRepositoryFilter) (*[]entities.Order, *int64, error) {
	var orderEntities []entities.Order

	result, err := r.collection.Find(context, filter)
	if err != nil {
		return nil, nil, err
	}

	for result.Next(context) {
		var item infra_model.OrderDocument
		err = result.Decode(&item)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		or, _ := item.ToEntity()
		orderEntities = append(orderEntities, *or)
	}

	count, _ := r.collection.CountDocuments(context, bson.M{"status": filter.Status.EnumIndex()})

	return &orderEntities, &count, nil
}
