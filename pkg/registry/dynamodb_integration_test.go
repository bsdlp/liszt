package registry

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var testRegistrar = &DynamoRegistrar{
	DB: dynamodb.New(session.New(aws.NewConfig().WithRegion("us-west-2"))),
	Config: &DynamoConfig{
		BuildingTableName: "liszt-buildings-testing",
		UnitTableName:     "liszt-units-testing",
		ResidentTableName: "liszt-residents-testing",
		AWSRegion:         "us-west-2",
	},
}
