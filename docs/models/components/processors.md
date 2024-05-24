# Processors


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Name`                                    | **string*                                 | :heavy_minus_sign:                        | CPU name                                  | Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz |
| `Arch`                                    | **string*                                 | :heavy_minus_sign:                        | CPU architecture                          | i386                                      |
| `Model`                                   | **int64*                                  | :heavy_minus_sign:                        | CPU model                                 | 158                                       |
| `FamilyName`                              | **string*                                 | :heavy_minus_sign:                        | CPU family                                |                                           |
| `Core`                                    | **int64*                                  | :heavy_minus_sign:                        | Number of core for that CPU               | 1                                         |
| `Speed`                                   | **int64*                                  | :heavy_minus_sign:                        | Speed (frequency) of the CPU              | 2800                                      |
| `Thread`                                  | **int64*                                  | :heavy_minus_sign:                        | Number of thread by core for the CPU      | 1                                         |
| `Stepping`                                | **int64*                                  | :heavy_minus_sign:                        | Stepping or power management information  | 9                                         |
| `Manufacturer`                            | **string*                                 | :heavy_minus_sign:                        | CPU manufacturer                          | Intel                                     |
| `Quantity`                                | **int64*                                  | :heavy_minus_sign:                        | Number of CPU with these features         | 1                                         |
| `Cpuid`                                   | **string*                                 | :heavy_minus_sign:                        | Identifier of the CPU                     |                                           |
| `ExternalClock`                           | **string*                                 | :heavy_minus_sign:                        | External clock used by the CPU            |                                           |
| `Description`                             | **string*                                 | :heavy_minus_sign:                        | System provided description of the CPU    |                                           |