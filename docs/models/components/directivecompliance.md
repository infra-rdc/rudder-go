# DirectiveCompliance


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | id of the directive                                                                                                    | 9a1773c9-0889-40b6-be89-f6504443ac1b                                                                                   |
| `Name`                                                                                                                 | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Name of the directive                                                                                                  | test directive                                                                                                         |
| `Mode`                                                                                                                 | [components.ComplianceDirectiveIDMode](../../models/components/compliancedirectiveidmode.md)                           | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `Compliance`                                                                                                           | *float32*                                                                                                              | :heavy_check_mark:                                                                                                     | Directive compliance level                                                                                             | 83.34                                                                                                                  |
| `ComplianceDetails`                                                                                                    | [components.ComplianceDirectiveIDComplianceDetails](../../models/components/compliancedirectiveidcompliancedetails.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `Rules`                                                                                                                | [components.DirectiveRuleCompliance](../../models/components/directiverulecompliance.md)                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `Nodes`                                                                                                                | [components.DirectiveNodeCompliance](../../models/components/directivenodecompliance.md)                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |