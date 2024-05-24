# CveCheck


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `CveID`                                                      | **string*                                                    | :heavy_minus_sign:                                           | CVE id                                                       | CVE-2019-5953                                                |
| `Score`                                                      | [*components.Score](../../models/components/score.md)        | :heavy_minus_sign:                                           | CVE score                                                    |                                                              |
| `Nodes`                                                      | []*string*                                                   | :heavy_minus_sign:                                           | Id of Nodes affected by this CVE                             |                                                              |
| `Packages`                                                   | [][components.Packages](../../models/components/packages.md) | :heavy_minus_sign:                                           | Packages affected by this CVE                                |                                                              |