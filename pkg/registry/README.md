## DB

### `residents`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `int` | internal id for a resident |
| `first_name` | `string` | first name of the resident |
| `middle_name` | `string` | middle name of the resident |
| `last_name` | `string` | last name of the resident |

### `units_residents`

| column | type | description |
| ------ | ---- | ----------- |
| `unit` | `int` | unit id |
| `resident` | `int` | resident id |

### `units`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `int` | internal id for a unit |
| `name` | `string` | unit name, usually a number |

### `buildings`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `int` | internal id for a building |
| `name` | `string` | name of the building |
| `address` | `string` | address of the building |
