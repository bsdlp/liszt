package registry

// DynamoConfig holds config options for DynamoRegistrar
type DynamoConfig struct {
	BuildingTableName string
	UnitTableName     string
	ResidentTableName string
}
