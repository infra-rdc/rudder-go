# SoftwareUpdate


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Name`                                                         | **string*                                                      | :heavy_minus_sign:                                             | name of software that can be updated                           |
| `Version`                                                      | **string*                                                      | :heavy_minus_sign:                                             | available version for software                                 |
| `Arch`                                                         | **string*                                                      | :heavy_minus_sign:                                             | CPU architecture of the update                                 |
| `From`                                                         | **string*                                                      | :heavy_minus_sign:                                             | tool that discovered that update                               |
| `Kind`                                                         | [*components.Kind](../../models/components/kind.md)            | :heavy_minus_sign:                                             | if available, kind of patch provided by that update, else none |
| `Source`                                                       | **string*                                                      | :heavy_minus_sign:                                             | information about the source providing that update             |
| `Description`                                                  | **string*                                                      | :heavy_minus_sign:                                             | details about the content of the update, if available          |
| `Severity`                                                     | [*components.Severity](../../models/components/severity.md)    | :heavy_minus_sign:                                             | if available, the severity of the update                       |
| `Ids`                                                          | []*string*                                                     | :heavy_minus_sign:                                             | N/A                                                            |