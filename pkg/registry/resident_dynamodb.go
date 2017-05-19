package registry

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

// GetResidentByID implements registrar
func (dr *DynamoRegistrar) GetResidentByID(ctx context.Context, residentID string) (resident *Resident, err error) {
	out, err := dr.DB.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(dr.Config.ResidentTableName),
		Key: map[string]*dynamodb.AttributeValue{
			residentIDAttributeName: {S: aws.String(residentID)},
		},
	})
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if out.Item == nil {
		return
	}

	resident = new(Resident)
	err = dynamodbattribute.UnmarshalMap(out.Item, resident)
	if err != nil {
		err = errors.WithStack(err)
		resident = nil
		return
	}
	return
}

// RegisterResident implements Registrar
func (dr *DynamoRegistrar) RegisterResident(ctx context.Context, in *Resident) (out *Resident, err error) {
	out = new(Resident)
	*out = *in
	out.ID = getULID().String()

	residentAV, err := dynamodbattribute.MarshalMap(out)
	if err != nil {
		out = nil
		err = errors.WithStack(err)
		return
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String(dr.Config.ResidentTableName),
		Item:      residentAV,
	}
	_, err = dr.DB.PutItemWithContext(ctx, params)
	if err != nil {
		out = nil
		err = errors.WithStack(err)
		return
	}
	return
}

// DeregisterResident implements Registrar
func (dr *DynamoRegistrar) DeregisterResident(ctx context.Context, residentID string) (err error) {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(dr.Config.ResidentTableName),
		Key: map[string]*dynamodb.AttributeValue{
			residentIDAttributeName: {S: aws.String(residentID)},
		},
	}
	_, err = dr.DB.DeleteItemWithContext(ctx, params)
	err = errors.WithStack(err)
	return
}

func (dr *DynamoRegistrar) batchGetResidents(ctx context.Context, residentIDs []string) (residents []*Resident, err error) {
	if residentIDs == nil || len(residentIDs) == 0 {
		residents = []*Resident{}
		return
	}

	keys := make([]map[string]*dynamodb.AttributeValue, len(residentIDs))
	for i, v := range residentIDs {
		keys[i] = map[string]*dynamodb.AttributeValue{
			residentIDAttributeName: {S: aws.String(v)},
		}
	}

	params := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			dr.Config.ResidentTableName: {Keys: keys},
		},
	}
	out, err := dr.DB.BatchGetItemWithContext(ctx, params)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	residentAVs, ok := out.Responses[dr.Config.ResidentTableName]
	if !ok {
		residents = []*Resident{}
		return
	}

	residents = make([]*Resident, len(residents))
	err = dynamodbattribute.UnmarshalListOfMaps(residentAVs, &residents)
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}
