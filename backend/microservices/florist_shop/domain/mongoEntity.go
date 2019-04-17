package domain

import (
	"gopkg.in/mgo.v2/bson"
)

//ID type
type ShopID bson.ObjectId

//ToString convert an ID in a string
func (i ShopID) String() string {
	return bson.ObjectId(i).Hex()
}

// MarshalJSON will marshal ID to Json
func (i ShopID) MarshalJSON() ([]byte, error) {
	return bson.ObjectId(i).MarshalJSON()
}

// UnmarshalJSON will convert a string to an ID
func (i *ShopID) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = s[1 : len(s)-1]
	if bson.IsObjectIdHex(s) {
		*i = ShopID(bson.ObjectIdHex(s))
	}

	return nil
}

// GetBSON implements bson.Getter.
func (i ShopID) GetBSON() (bson.ObjectId, error) {
	if i == "" {
		return "", nil
	}
	return bson.ObjectId(i), nil
}

// SetBSON implements bson.Setter.
func (i *ShopID) SetBSON(raw bson.Raw) error {
	decoded := new(string)
	bsonErr := raw.Unmarshal(decoded)
	if bsonErr == nil {
		*i = ShopID(bson.ObjectId(*decoded))
		return nil
	}
	return bsonErr
}

//StringToID convert a string to an ID
func StringToID(s string) ShopID {
	return ShopID(bson.ObjectIdHex(s))
}

//IsValidID check if is a valid ID
func IsValidID(s string) bool {
	return bson.IsObjectIdHex(s)
}

//NewID create a new id
func NewID() ShopID {
	return StringToID(bson.NewObjectId().Hex())
}
