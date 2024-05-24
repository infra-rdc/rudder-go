# GetCampaignResultsRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `ID`                                              | *string*                                          | :heavy_check_mark:                                | Id of the campaign                                | 0076a379-f32d-4732-9e91-33ab219d8fde              |
| `Limit`                                           | **int64*                                          | :heavy_minus_sign:                                | Max number of elements in response                |                                                   |
| `Offset`                                          | **int64*                                          | :heavy_minus_sign:                                | Offset of data in response (skip X elements)      |                                                   |
| `Before`                                          | [*types.Date](../../types/date.md)                | :heavy_minus_sign:                                | N/A                                               |                                                   |
| `After`                                           | [*types.Date](../../types/date.md)                | :heavy_minus_sign:                                | N/A                                               |                                                   |
| `Order`                                           | **string*                                         | :heavy_minus_sign:                                | N/A                                               |                                                   |
| `Asc`                                             | [*components.Asc](../../models/components/asc.md) | :heavy_minus_sign:                                | N/A                                               |                                                   |