package dbrepository

import (
	domain "OnlineFlorist/backend/microservices/customer/domain"
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

var collectionName = "customer"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

//create restaurant(reader)
func (r *MongoRepository) Create(b *domain.Customer) (domain.ID, error) {

	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ID(0) == b.CustID {
		b.CustID = domain.NewID()
	}
	_, err := coll.UpsertId(b.CustID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.CustID, nil
}

func (r *MongoRepository) Insert(filename string) (int, error) {
	fname, _ := os.Open(filename)
	defer fname.Close()
	fp := bufio.NewScanner(fname)
	line := ""
	var final domain.Customer
	rcnt := 0
	for fp.Scan() {
		rcnt += 1
		line = fp.Text()
		json.Unmarshal([]byte(line), &final)
		final.CustID = domain.NewID()
		p, _ := r.Store(&final)
		fmt.Println(p)
	}
	return rcnt, nil
}

//Find a Restaurant(reader)
func (r *MongoRepository) Get(id domain.ID) (*domain.Customer, error) {
	result := domain.Customer{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
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
func (r *MongoRepository) GetAll() ([]*domain.Customer, error) {
	result := []*domain.Customer{}
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

//Find a Restaurant By Name(reader)
func (r *MongoRepository) FindByName(name string) ([]*domain.Customer, error) {
	fmt.Println("Name : Findig", name)
	result := []*domain.Customer{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	//err := coll.Find(bson.M{"name":bson.RegEx{name,"i"}}).All(&result)

	err := coll.Find(bson.M{"name": name}).All(&result)
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
func (r *MongoRepository) Store(b *domain.Customer) (domain.ID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ID(0) == b.CustID {
		b.CustID = domain.NewID()
	}
	_, err := coll.UpsertId(b.CustID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.CustID, nil
}

//Delete a Restaurant record(writer)
func (r *MongoRepository) Delete(id domain.ID) error {
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
func (r *MongoRepository) FindByTypeOfFood(foodtype string) ([]*domain.Customer, error) {
	result := []*domain.Customer{}
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
func (r *MongoRepository) FindByTypeOfPostCode(postcode string) ([]*domain.Customer, error) {
	result := []*domain.Customer{}
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
func (r *MongoRepository) Search(query string) ([]*domain.Customer, error) {
	result := []*domain.Customer{}
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
