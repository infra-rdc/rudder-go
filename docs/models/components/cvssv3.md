# Cvssv3

CVSS V3 of the CVE


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `BaseScore`                                                         | **float64*                                                          | :heavy_minus_sign:                                                  | CVSS V3 base score                                                  | 9.8                                                                 |
| `Vector`                                                            | **string*                                                           | :heavy_minus_sign:                                                  | CVSS V3 vector                                                      | CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H                        |
| `BaseSeverity`                                                      | [*components.BaseSeverity](../../models/components/baseseverity.md) | :heavy_minus_sign:                                                  | CVSS V3 Severity                                                    |                                                                     |