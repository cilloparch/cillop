package mongodb

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	EntityCreator[Entity any]                    func() *Entity
	EntityTransformer[Entity any, Transform any] func(*Entity) Transform
)

type Helper[Entity any, Transform any] interface {
	GetFilterCount(ctx context.Context, filter bson.M, options ...*options.CountOptions) (int64, *i18np.Error)
	GetListFilter(ctx context.Context, filter bson.M, options ...*options.FindOptions) ([]Entity, *i18np.Error)
	GetListFilterTransform(ctx context.Context, filter bson.M, transform EntityTransformer[Entity, Transform], options ...*options.FindOptions) ([]Transform, *i18np.Error)
	GetAggregateListFilter(ctx context.Context, pipeline mongo.Pipeline, options ...*options.AggregateOptions) ([]Entity, *i18np.Error)
	GetAggregateListFilterTransform(ctx context.Context, pipeline mongo.Pipeline, transform EntityTransformer[Entity, Transform], options ...*options.AggregateOptions) ([]Transform, *i18np.Error)
	GetFilter(ctx context.Context, filter bson.M, options ...*options.FindOneOptions) (*Entity, bool, *i18np.Error)
	UpdateOne(ctx context.Context, filter bson.M, setter bson.M, options ...*options.UpdateOptions) *i18np.Error
}

type helper[Entity any, Transform any] struct {
	collection *mongo.Collection
	creator    EntityCreator[Entity]
}

func NewHelper[Entity any, Transform any](collection *mongo.Collection, creator EntityCreator[Entity]) Helper[Entity, Transform] {
	return &helper[Entity, Transform]{
		collection: collection,
		creator:    creator,
	}
}

func (h *helper[Entity, Transform]) GetFilterCount(ctx context.Context, filter bson.M, options ...*options.CountOptions) (int64, *i18np.Error) {
	count, err := h.collection.CountDocuments(ctx, filter, options...)
	if err != nil {
		return 0, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
	}
	return count, nil
}

func (h *helper[Entity, Transform]) GetListFilter(ctx context.Context, filter bson.M, options ...*options.FindOptions) ([]Entity, *i18np.Error) {
	var l []Entity
	cur, err := h.collection.Find(ctx, filter, options...)
	if err != nil {
		return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
	}
	for cur.Next(ctx) {
		o := h.creator()
		err := cur.Decode(o)
		if err != nil {
			return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
		}
		l = append(l, *o)
	}
	return l, nil
}

func (h *helper[Entity, Transform]) GetListFilterTransform(ctx context.Context, filter bson.M, transform EntityTransformer[Entity, Transform], options ...*options.FindOptions) ([]Transform, *i18np.Error) {
	var l []Transform
	cur, err := h.collection.Find(ctx, filter, options...)
	if err != nil {
		return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
	}
	for cur.Next(ctx) {
		o := h.creator()
		err := cur.Decode(o)
		if err != nil {
			return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
		}
		l = append(l, transform(o))
	}
	return l, nil
}

func (h *helper[Entity, Transform]) GetAggregateListFilter(ctx context.Context, pipeline mongo.Pipeline, options ...*options.AggregateOptions) ([]Entity, *i18np.Error) {
	var l []Entity
	cur, err := h.collection.Aggregate(ctx, pipeline, options...)
	if err != nil {
		return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
	}
	for cur.Next(ctx) {
		o := h.creator()
		err := cur.Decode(o)
		if err != nil {
			return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
		}
		l = append(l, *o)
	}
	return l, nil
}

func (h *helper[Entity, Transform]) GetAggregateListFilterTransform(ctx context.Context, pipeline mongo.Pipeline, transform EntityTransformer[Entity, Transform], options ...*options.AggregateOptions) ([]Transform, *i18np.Error) {
	var l []Transform
	cur, err := h.collection.Aggregate(ctx, pipeline, options...)
	if err != nil {
		return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
	}
	for cur.Next(ctx) {
		o := h.creator()
		err := cur.Decode(o)
		if err != nil {
			return nil, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
		}
		l = append(l, transform(o))
	}
	return l, nil
}

func (h *helper[Entity, Transform]) GetFilter(ctx context.Context, filter bson.M, options ...*options.FindOneOptions) (*Entity, bool, *i18np.Error) {
	o := h.creator()
	res := h.collection.FindOne(ctx, filter, options...)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil
		}
		return nil, true, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
	}
	err := res.Decode(&o)
	if err != nil {
		return nil, true, i18np.NewError("get_query_failed", i18np.P{"Error": err.Error()})
	}
	return o, true, nil
}

func (h *helper[Entity, Transform]) UpdateOne(ctx context.Context, filter bson.M, setter bson.M, options ...*options.UpdateOptions) *i18np.Error {
	res, err := h.collection.UpdateOne(ctx, filter, setter, options...)
	if err != nil {
		return i18np.NewError("update_query_failed", i18np.P{"Error": err.Error()})
	}
	if res.MatchedCount == 0 {
		return i18np.NewError("update_query_not_found", i18np.P{"Error": "No document matched"})
	}
	return nil
}
