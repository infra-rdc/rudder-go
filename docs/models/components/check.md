# Check


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Name`                                                           | *string*                                                         | :heavy_check_mark:                                               | Name of the check                                                | RAM available                                                    |
| `Msg`                                                            | *string*                                                         | :heavy_check_mark:                                               | Message about the check                                          | Only 2GB of RAM left                                             |
| `Status`                                                         | [components.CheckStatus](../../models/components/checkstatus.md) | :heavy_check_mark:                                               | The severity status of the check                                 |                                                                  |