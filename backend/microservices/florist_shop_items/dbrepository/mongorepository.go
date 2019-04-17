package dbrepository

import (
	shoprepo "OnlineFlorist/backend/microservices/florist_shop/dbrepository"
	shop "OnlineFlorist/backend/microservices/florist_shop/domain"
	shopmongoutils "OnlineFlorist/backend/microservices/florist_shop/utils"
	domain "OnlineFlorist/backend/microservices/florist_shop_items/domain"

	"bufio"
	"encoding/json"
	"fmt"

	"os"
	"strings"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "floristshopitems"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

//create restaurant(reader)
func (r *MongoRepository) Create(b *domain.FloristShopItems) (domain.ItemID, error) {

	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ItemID(0) == b.ItemID {
		b.ItemID = domain.NewID()
	}
	_, err := coll.UpsertId(b.ItemID, b)

	if err != nil {
		return domain.ItemID(0), err
	}
	return b.ItemID, nil
}

func (r *MongoRepository) Insert(filename string) (int, error) {
	fname, _ := os.Open(filename)
	defer fname.Close()
	fp := bufio.NewScanner(fname)
	line := ""
	var final domain.FloristShopItems
	rcnt := 0
	for fp.Scan() {
		rcnt += 1
		line = fp.Text()
		json.Unmarshal([]byte(line), &final)
		final.ItemID = domain.NewID()
		p, _ := r.Store(&final)
		fmt.Println(p)
	}
	return rcnt, nil
}

//Find a Restaurant(reader)
func (r *MongoRepository) Get(id domain.ItemID) (*domain.FloristShopItems, error) {
	result := domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	fmt.Println("GTEE", err, result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//get all restaurants(reader)
func (r *MongoRepository) GetAll() ([]*domain.FloristShopItems, error) {
	result := []*domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	fmt.Println(coll.Find(bson.M{}).All(&result))
	err := coll.Find(bson.M{}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

func (r *MongoRepository) SortByRating() ([]*domain.FloristShopItems, error) {
	result := []*domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	// fmt.Println(coll.Find(bson.M{}).Sort("category").All(&result))
	err := coll.Find(bson.M{}).Sort("-rating").All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Find a Restaurant By Name(reader)
func (r *MongoRepository) FindByCategory(category string) ([]*domain.FloristShopItems, error) {
	// fmt.Println("Name : Findig", name)
	result := []*domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	//err := coll.Find(bson.M{"name":bson.RegEx{name,"i"}}).All(&result)

	err := coll.Find(bson.M{"category": category}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Store a Restaurant record(writer)
func (r *MongoRepository) Store(b *domain.FloristShopItems) (domain.ItemID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ItemID(0) == b.ItemID {
		b.ItemID = domain.NewID()
	}
	_, err := coll.UpsertId(b.ItemID, b)

	if err != nil {
		return domain.ItemID(0), err
	}
	return b.ItemID, nil
}

func (r *MongoRepository) FindShop(id domain.ItemID) *shop.FloristShop {
	result := domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	// fmt.Println("shop in find:", result, err)
	// shop1 := shoprepo.get(result["_shopid"])
	// return shop1
	switch err {
	case nil:

		mongoSession, _ := shopmongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

		dbname := "floristshopitem"
		repoaccess := shoprepo.NewMongoRepository(mongoSession, dbname)

		// var Shop = shop.FloristShop{}
		// errS := ""
		// if err != nil {
		// var err3 error
		Shop, err1 := repoaccess.Get(result.ShopID)
		// fmt.Println(Shop, err)
		if err1 != nil {
			return &shop.FloristShop{}
		}
		return Shop
	default:
		return nil
	}
}

//Delete a Restaurant record(writer)
func (r *MongoRepository) Delete(id domain.ItemID) error {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Remove(bson.M{"_id": id})
	switch err {
	case nil:
		return nil
	case mgo.ErrNotFound:
		return domain.ErrNotFound
	default:
		return err
	}
}

//Find a Restaurant By Type Of Food(filter)
func (r *MongoRepository) FindByTypeOfFood(foodtype string) ([]*domain.FloristShopItems, error) {
	result := []*domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"type_of_food": foodtype}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Find a Restaurant By Type Of Post Code(filter)
func (r *MongoRepository) FindByTypeOfPostCode(postcode string) ([]*domain.FloristShopItems, error) {
	result := []*domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"postcode": postcode}).All(result)

	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Count number of Restaurant By Type Of Food(filter)
func (r *MongoRepository) CountByTypeOfFood(foodtype string) (int, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	fcnt, err := coll.Find(bson.M{"type_of_food": foodtype}).Count()
	switch err {
	case nil:
		return fcnt, nil
	case mgo.ErrNotFound:
		return 0, domain.ErrNotFound
	default:
		return 1, err
	}
}

//Count number of Restaurant By Type Of Post Code(filter)
func (r *MongoRepository) CountByTypeOfPostCode(postcode string) (int, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	fcnt, err := coll.Find(bson.M{"postcode": postcode}).Count()
	switch err {
	case nil:
		return fcnt, nil
	case mgo.ErrNotFound:
		return 0, domain.ErrNotFound
	default:
		return 1, err
	}
}

//Search on Restaurant (filter)
func (r *MongoRepository) Search(query string) ([]*domain.FloristShopItems, error) {
	result := []*domain.FloristShopItems{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)

	arr := strings.Split(query, "=")
	key := arr[0]
	value := arr[1]
	//        fmt.Println(key,value)
	err := coll.Find(bson.M{key: bson.M{"$regex": bson.RegEx{value, ""}}}).All(&result)
	fmt.Println("RESULT :: ", result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}
