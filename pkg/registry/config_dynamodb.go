package registry

// DynamoConfig holds config options for DynamoRegistrar
type DynamoConfig struct {
	BuildingTableName string `envconfig:"building_table_name" default:"liszt-buildings"`
	UnitTableName     string `envconfig:"unit_table_name" default:"liszt-units"`
	ResidentTableName string `envconfig:"resident_table_name" default:"liszt-residents"`
	AWSRegion         string `envconfig:"aws_region" default:"us-west-2"`
}
