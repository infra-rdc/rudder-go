# DirectiveRuleCompliance


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ID`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | id of the rule                                                               | 32377fd7-02fd-43d0-aab7-28460a91347b                                         |
| `Name`                                                                       | *string*                                                                     | :heavy_check_mark:                                                           | Name of the rule                                                             | Global configuration for all nodes                                           |
| `Mode`                                                                       | [*components.Mode](../../models/components/mode.md)                          | :heavy_minus_sign:                                                           | N/A                                                                          |                                                                              |
| `Compliance`                                                                 | *float32*                                                                    | :heavy_check_mark:                                                           | Directive compliance level                                                   | 83.34                                                                        |
| `ComplianceDetails`                                                          | [components.ComplianceDetails](../../models/components/compliancedetails.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `Components`                                                                 | [][components.Components](../../models/components/components.md)             | :heavy_minus_sign:                                                           | N/A                                                                          |                                                                              |