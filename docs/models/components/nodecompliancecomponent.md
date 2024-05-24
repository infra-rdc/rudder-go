# NodeComplianceComponent


## Fields

| Field                                                                                                                      | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                                       | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | id of the node                                                                                                             | root                                                                                                                       |
| `Name`                                                                                                                     | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Name of the node                                                                                                           | server.rudder.local                                                                                                        |
| `Mode`                                                                                                                     | [*components.NodeComplianceComponentMode](../../models/components/nodecompliancecomponentmode.md)                          | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |                                                                                                                            |
| `Compliance`                                                                                                               | *float32*                                                                                                                  | :heavy_check_mark:                                                                                                         | Directive compliance level                                                                                                 | 83.34                                                                                                                      |
| `ComplianceDetails`                                                                                                        | [components.NodeComplianceComponentComplianceDetails](../../models/components/nodecompliancecomponentcompliancedetails.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |                                                                                                                            |
| `Values`                                                                                                                   | [][components.Values](../../models/components/values.md)                                                                   | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |                                                                                                                            |