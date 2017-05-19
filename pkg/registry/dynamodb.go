package registry

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

// DynamoRegistrar implements Registrar using dynamodb
type DynamoRegistrar struct {
	DB     dynamodbiface.DynamoDBAPI
	Config *DynamoConfig
}

const (
	buildingIDAttributeName = "building_id"
	unitIDAttributeName     = "unit_id"
	residentIDAttributeName = "resident_id"
	buildingUnitsGSIName    = "building_unit_gsi"
)
