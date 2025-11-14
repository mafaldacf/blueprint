package foobar

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	//"go.mongodb.org/mongo-driver/bson"
)

type Foo struct {
	FooID    string
	Text     string
	BarItems []Bar
	FooMap   map[string]string
	//OtherFoo *Foo
}

func (f *Foo) GetID() string {
	return f.FooID
}

type FooPtr struct {
	FooMap  map[string]*Foo
	FooMap2 map[string]*Foo
}

type FooService interface {
	WriteFoo(ctx context.Context, id string, text string, key string, val string, bars []Bar) (Foo, error)
	//ReadFoo(ctx context.Context, id string) (Foo, error)
}

type FooServiceImpl struct {
	fooDb      backend.NoSQLDatabase
	barService BarService
}

func NewFooServiceImpl(ctx context.Context, barService BarService, fooDb backend.NoSQLDatabase) (FooService, error) {
	d := &FooServiceImpl{barService: barService, fooDb: fooDb}
	return d, nil
}

func (s *FooServiceImpl) WriteFoo(ctx context.Context, id string, text string, key string, val string, barItems []Bar) (Foo, error) {
	collection, err := s.fooDb.GetCollection(ctx, "foo_db", "foo")
	if err != nil {
		return Foo{}, err
	}

	/* var barItems []Bar
	for _, item := range bars {
		barItems = append(barItems, item)
	} */

	/* for _, item := range barItems {
		_, err := s.barService.ReadBar(ctx, item.BarID)
		if err != nil {
			return Foo{}, err
		}
	} */

	// --------
	// ORIGINAL
	// --------
	foo := Foo{
		FooID:    id,
		Text:     text,
		BarItems: barItems,
		FooMap:   make(map[string]string),
	}
	foo.FooMap["first_key"] = id
	foo.FooMap["second_key"] = text
	foo.FooMap[key] = val
	err = collection.InsertOne(ctx, foo)
	if err != nil {
		return Foo{}, err
	}

	return foo, nil

	// ------------
	// EXPERIMENT 1
	// ------------

	/* foo1 := Foo{
		FooID: id,
		Text:  text,
	}

	foo2 := Foo{
		FooID: id,
		Text:  text,
	}

	fooPtr := &FooPtr{
		FooMap: make(map[string]*Foo),
	}
	fooPtr.FooMap["key"] = &foo1

	fooPtr.FooMap["key"] = &foo2
	fooPtr.FooMap["key"].Text = "new text 1!"

	fooPtr.FooMap["key"] = &foo2
	fooPtr.FooMap["key"].Text = "new text 2!"

	err = collection.InsertOne(ctx, fooPtr.FooMap["key"])
	if err != nil {
		return Foo{}, err
	}
	return foo1, nil */

	// ------------
	// EXPERIMENT 2
	// ------------

	/* var foo0, foo1, foo2 Foo
	foo1.FooID = "myid1"

	err = collection.InsertOne(ctx, foo1)
	if err != nil {
		return Foo{}, err
	}

	foo1.OtherFoo = &foo0
	foo0.FooID = "myid0"
	foo1.OtherFoo.FooID = "myotherfooid0"

	foo2 = foo1

	foo1.FooID = id
	foo1.Text = text

	return foo2, nil */

	// ------------
	// EXPERIMENT 3
	// ------------

	/* var foo0 = &Foo{}
	var foo1 = &Foo{}
	foo1.SetOtherFoo(foo0)
	foo0.FooID = "myid0"

	err = collection.InsertOne(ctx, foo1.GetOtherFoo())
	if err != nil {
		return Foo{}, err
	}

	return *foo1, nil */
}

/* func (s *FooServiceImpl) ReadFoo(ctx context.Context, id string) (Foo, error) {
	var foo Foo

	collection, err := s.fooDb.GetCollection(ctx, "foo_db", "foo")
	if err != nil {
		return Foo{}, err
	}

	query := bson.D{{Key: "FooID", Value: id}}
	cursor, err := collection.FindOne(ctx, query)
	if err != nil {
		return Foo{}, err
	}

	res, err := cursor.One(ctx, &foo)
	if !res || err != nil {
		return Foo{}, err
	}

	return foo, nil
}
*/
