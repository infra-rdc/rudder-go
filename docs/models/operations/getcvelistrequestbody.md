# GetCVEListRequestBody

List of CVE ids you want


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `CveIds`                                                         | []*string*                                                       | :heavy_minus_sign:                                               | N/A                                                              |                                                                  |
| `OnlyScore`                                                      | **bool*                                                          | :heavy_minus_sign:                                               | Only send score of the cve, and not the whole detailed list      | true                                                             |
| `MinScore`                                                       | **string*                                                        | :heavy_minus_sign:                                               | Only send CVE with a score higher than the value                 | 7.5                                                              |
| `MaxScore`                                                       | **string*                                                        | :heavy_minus_sign:                                               | Only send CVE with a score lower than the value                  | 8.5                                                              |
| `PublishedDate`                                                  | [*types.Date](../../types/date.md)                               | :heavy_minus_sign:                                               | Only send CVE with a publication date more recent than the value |                                                                  |