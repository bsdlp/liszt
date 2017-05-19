package registry

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

type dynamodbUnit struct {
	ID         string `dynamodbav:"unit_id"`
	Name       string
	BuildingID string    `dynamodbav:"building_id"`
	Residents  []string  `dynamodb:",stringset"`
	UpdatedAt  time.Time `dynamodbav:",unixtime"`
}

// ListBuildingUnits implements Registrar
func (dr *DynamoRegistrar) ListBuildingUnits(ctx context.Context, buildingID string) (units []*Unit, err error) {
	out, err := dr.DB.QueryWithContext(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(dr.Config.UnitTableName),
		IndexName:              aws.String(buildingUnitsGSIName),
		KeyConditionExpression: aws.String("#building_id=:building_id"),
		ExpressionAttributeNames: map[string]*string{
			"#building_id": aws.String(buildingIDAttributeName),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":building_id": {S: aws.String(buildingID)},
		},
	})
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	dbUnits := make([]*dynamodbUnit, aws.Int64Value(out.Count))
	err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &dbUnits)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	units = make([]*Unit, len(dbUnits))
	for i, v := range dbUnits {
		units[i] = &Unit{
			ID:   v.ID,
			Name: v.Name,
		}
	}
	return
}

// RegisterUnit implements Registrar
func (dr *DynamoRegistrar) RegisterUnit(ctx context.Context, buildingID string, in *Unit) (unit *Unit, err error) {
	id := getULID().String()
	item, err := dynamodbattribute.MarshalMap(&dynamodbUnit{
		ID:         id,
		Name:       in.Name,
		BuildingID: buildingID,
		UpdatedAt:  time.Now().UTC(),
	})
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// omitempty does not do the needful
	// https://github.com/aws/aws-sdk-go/issues/682
	delete(item, "Residents")

	params := &dynamodb.PutItemInput{
		TableName: aws.String(dr.Config.UnitTableName),
		Item:      item,
	}
	_, err = dr.DB.PutItemWithContext(ctx, params)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	unit = &Unit{
		ID:   id,
		Name: in.Name,
	}
	return
}

// DeregisterUnit implements Registrar
func (dr *DynamoRegistrar) DeregisterUnit(ctx context.Context, unitID string) (err error) {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(dr.Config.UnitTableName),
		Key: map[string]*dynamodb.AttributeValue{
			unitIDAttributeName: {S: aws.String(unitID)},
		},
	}
	_, err = dr.DB.DeleteItemWithContext(ctx, params)
	err = errors.WithStack(err)
	return
}

// ListUnitResidents implements Registrar
func (dr *DynamoRegistrar) ListUnitResidents(ctx context.Context, unitID string) (residents []*Resident, err error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(dr.Config.UnitTableName),
		Key: map[string]*dynamodb.AttributeValue{
			unitIDAttributeName: {S: aws.String(unitID)},
		},
	}
	out, err := dr.DB.GetItemWithContext(ctx, params)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	obj := new(dynamodbUnit)
	err = dynamodbattribute.UnmarshalMap(out.Item, obj)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	residents, err = dr.batchGetResidents(ctx, obj.Residents)
	return
}

// MoveResidentIn implements Registrar
func (dr *DynamoRegistrar) MoveResidentIn(ctx context.Context, residentID, unitID string) (err error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(dr.Config.UnitTableName),
		Key: map[string]*dynamodb.AttributeValue{
			unitIDAttributeName: {S: aws.String(unitID)},
		},
		UpdateExpression: aws.String("ADD Residents :residents SET UpdatedAt = :timestamp"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":residents": {
				SS: []*string{aws.String(residentID)},
			},
			":timestamp": {N: aws.String(timestamp)},
		},
	}
	_, err = dr.DB.UpdateItemWithContext(ctx, params)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// MoveResidentOut implements Registrar
func (dr *DynamoRegistrar) MoveResidentOut(ctx context.Context, residentID, unitID string) (err error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(dr.Config.UnitTableName),
		Key: map[string]*dynamodb.AttributeValue{
			unitIDAttributeName: {S: aws.String(unitID)},
		},
		UpdateExpression: aws.String("SET UpdatedAt = :timestamp DELETE Residents :resident"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":resident":  {SS: []*string{aws.String(residentID)}},
			":timestamp": {N: aws.String(timestamp)},
		},
	}
	_, err = dr.DB.UpdateItemWithContext(ctx, params)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
