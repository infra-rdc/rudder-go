# RunParameters

Parameters to configure when the data source is fetched to update node properties.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `OnGeneration`                                              | **bool*                                                     | :heavy_minus_sign:                                          | Trigger a fetch at the beginning of a policy generation     | true                                                        |
| `OnNewNode`                                                 | **bool*                                                     | :heavy_minus_sign:                                          | Trigger a fetch when a new node is accepted, for that node  | true                                                        |
| `Schedule`                                                  | [*components.Schedule](../../models/components/schedule.md) | :heavy_minus_sign:                                          | Configure if data source should be fetch periodically       |                                                             |