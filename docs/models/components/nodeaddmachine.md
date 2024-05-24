# NodeAddMachine

The kind of machine for the node (use vm for a generic VM)


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Type`                                                           | [components.NodeAddType](../../models/components/nodeaddtype.md) | :heavy_check_mark:                                               | list of groups to include in rule application                    |                                                                  |
| `Provider`                                                       | [*components.Provider](../../models/components/provider.md)      | :heavy_minus_sign:                                               | The kind of virtual machine for the node                         |                                                                  |
| `Manufacturer`                                                   | **string*                                                        | :heavy_minus_sign:                                               | Manufacturer of the machine                                      | corp inc.                                                        |
| `SerialNumber`                                                   | **string*                                                        | :heavy_minus_sign:                                               | Serial number of the machine                                     | ece12459-2b90-49c9-ab1e-72e38f797421                             |