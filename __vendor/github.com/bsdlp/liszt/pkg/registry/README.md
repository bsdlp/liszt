## DB

### `residents`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `int` | internal id for a resident |
| `firstname` | `string` | first name of the resident |
| `middlename` | `string` | middle name of the resident |
| `lastname` | `string` | last name of the resident |

### `units_residents`

| column | type | description |
| ------ | ---- | ----------- |
| `unit` | `int` | unit id |
| `resident` | `int` | resident id |

### `units`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `int` | internal id for a unit |
| `display_name` | `string` | unit name, usually a number |

### `buildings`

| column | type | description |
| ------ | ---- | ----------- |
| `id` | `int` | internal id for a building |
| `name` | `string` | name of the building |
| `address` | `string` | address of the building |
