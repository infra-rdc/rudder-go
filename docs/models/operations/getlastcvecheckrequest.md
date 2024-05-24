# GetLastCVECheckRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `GroupID`                                                         | **string*                                                         | :heavy_minus_sign:                                                | Id of node groups you want to get from last CVE check             |
| `NodeID`                                                          | **string*                                                         | :heavy_minus_sign:                                                | Id of nodes you want to get from last CVE check                   |
| `CveID`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Id of CVE you want to get from last CVE check                     |
| `Package`                                                         | **string*                                                         | :heavy_minus_sign:                                                | Name of packages you want to get from last CVE check              |
| `Severity`                                                        | [*components.CveSeverity](../../models/components/cveseverity.md) | :heavy_minus_sign:                                                | Severity of the CVE you want to get from last CVE check           |