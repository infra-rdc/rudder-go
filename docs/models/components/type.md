# Type

Define and configure data source type.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Name`                                                                              | [*components.DatasourceName](../../models/components/datasourcename.md)             | :heavy_minus_sign:                                                                  | Data source type name                                                               | HTTP                                                                                |
| `Parameters`                                                                        | [*components.DatasourceParameters](../../models/components/datasourceparameters.md) | :heavy_minus_sign:                                                                  | You can use Rudder variable expansion (`${rudder.node`, `${node.properties...}`)    |                                                                                     |