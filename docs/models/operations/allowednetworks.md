# AllowedNetworks


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Add`                                                                     | []*string*                                                                | :heavy_minus_sign:                                                        | List of networks to add to existing allowed networks for that server      | [<br/>"192.168.2.0/24",<br/>"192.168.0.0/16"<br/>]                        |
| `Delete`                                                                  | []*string*                                                                | :heavy_minus_sign:                                                        | List of networks to remove from existing allowed networks for that server | [<br/>"162.168.1.0/24"<br/>]                                              |