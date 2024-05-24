# APIEndpoints

objects with two fields, the first one has the endpoint name as key and its description as value, the second one has HTTP verb to use (GET, POST PUT, DELETE) as key and the supported version an API path for value.


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `EndpointName`                                          | **string*                                               | :heavy_minus_sign:                                      | The endpoint name for key and its description for value |
| `HTTPVerb`                                              | *any*                                                   | :heavy_minus_sign:                                      | N/A                                                     |