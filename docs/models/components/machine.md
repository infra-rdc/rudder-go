# Machine

Information about the underlying machine


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ID`                                                                | **string*                                                           | :heavy_minus_sign:                                                  | Rudder unique identifier for the machine                            |                                                                     |
| `Type`                                                              | [*components.NodeFullType](../../models/components/nodefulltype.md) | :heavy_minus_sign:                                                  | Type of the machine                                                 | Virtual                                                             |
| `Provider`                                                          | **string*                                                           | :heavy_minus_sign:                                                  | In the case of VM, the VM technology                                | vbox                                                                |
| `Manufacturer`                                                      | **string*                                                           | :heavy_minus_sign:                                                  | Information about machine manufacturer                              | innotek GmbH                                                        |
| `SerialNumber`                                                      | **string*                                                           | :heavy_minus_sign:                                                  | If available, a unique identifier provided by the machine           | ece12459-2b90-49c9-ab1e-72e38f797421                                |