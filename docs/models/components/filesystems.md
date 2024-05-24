# FileSystems


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Name`                                     | **string*                                  | :heavy_minus_sign:                         | Type of file system (`ext4`, `swap`, etc.) | ext4                                       |
| `MountPoint`                               | **string*                                  | :heavy_minus_sign:                         | Mount point                                | /srv                                       |
| `Description`                              | **string*                                  | :heavy_minus_sign:                         | Description of the file system             |                                            |
| `FileCount`                                | **int64*                                   | :heavy_minus_sign:                         | Number of files                            | 1456                                       |
| `FreeSpace`                                | **int64*                                   | :heavy_minus_sign:                         | Free space remaining                       | 3487                                       |
| `TotalSpace`                               | **int64*                                   | :heavy_minus_sign:                         | Total space                                | 208869                                     |