## DB

### `residents`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `string` | internal id for a resident |
| `name` | `string` | name of the resident |

### `units_residents`

| column | type | description |
| ------ | ---- | ----------- |
| `unit` | `string` | unit id |
| `resident` | `string` | resident id |

### `units`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `string` | internal id for a unit |
| `name` | `string` | unit name, usually a number |

### `buildings`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `string` | internal id for a building |
| `name` | `string` | name of the building |
| `address` | `string` | address of the building |
