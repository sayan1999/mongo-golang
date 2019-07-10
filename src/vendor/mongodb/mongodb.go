package mongodb

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "qilog"
    "fmt"
    "errors"
)

var ctx = context.TODO()
type Collection struct{ collectionPtr *mongo.Collection }
func (c Collection) string() string{
    return "DB: " + c.collectionPtr.Database().Name() + " Collection: " + c.collectionPtr.Name()
}
func (c Collection) IsNIL() bool{
    return c.collectionPtr == nil
}

type Client struct{ clientPtr *mongo.Client }
func (c Client) string() string {
    return fmt.Sprintf("%#v", c.clientPtr)    
}
func (c Client) IsNIL() bool{
    return c.clientPtr == nil
}

func Initialize(hostID string, port string) (Client, error) {
  var err error
  var client Client
  client.clientPtr, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+hostID+":"+port))
  if qilog.Check(err){
    return client, err
  }

  err = client.clientPtr.Ping(ctx, nil)
  if qilog.Check(err){
    return client, err
  }
  return client, nil
}

func (client Client) GetCollection(dbName string, colName string) (Collection, error) {

    collection := Collection{client.clientPtr.Database(dbName).Collection(colName)}
    if collection.collectionPtr == nil{
      return Collection{nil}, errors.New("Collection not found")
    }
    return collection, nil
}

func (collection Collection) Insert(toInsert []interface{}) (*mongo.InsertManyResult, error) {

    insertedResults, err := collection.collectionPtr.InsertMany(ctx, toInsert)
    return insertedResults, err
}

func (collection Collection) Find(filter interface{}, projection interface{}, flag int) ([]interface{}, string) {

    var errorDetected string
    var findResult []interface{}
    if flag == 1 {

        var result interface{}
        err := collection.collectionPtr.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
        if qilog.Check(err) {
            errorDetected = err.Error()
        }
        findResult = append(findResult, result)

    } else {

        cursor, err := collection.collectionPtr.Find(ctx, filter, options.Find().SetProjection(projection))
        if err != nil {
            return nil, err.Error()
        }
        for cursor.Next(ctx) {
            var result interface{}
            err := cursor.Decode(&result)
            if qilog.Check(err) {
                errorDetected += err.Error()
            }           
            findResult = append(findResult, result)
        }
    }
    return findResult, errorDetected
}

func (collection Collection) Update(filter interface{}, update interface{}, flag int) (*mongo.UpdateResult, error) {

    var (
        err           error
        updatedResult *mongo.UpdateResult
    )

    if flag == 1 {
        updatedResult, err = collection.collectionPtr.UpdateOne(ctx, filter, update)
    } else {
        updatedResult, err = collection.collectionPtr.UpdateMany(ctx, filter, update)

    }
    
    return updatedResult, err
}

func (collection Collection) Delete(filter interface{}, flag int) (*mongo.DeleteResult, error) {

    var (
        err           error
        deletedResult *mongo.DeleteResult
    )

    if flag == 1 {
        deletedResult, err = collection.collectionPtr.DeleteOne(ctx, filter)
    } else {
        deletedResult, err = collection.collectionPtr.DeleteMany(ctx, filter)
    }

    if qilog.Check(err) {
      return nil, err
    }
    return deletedResult, nil
}
