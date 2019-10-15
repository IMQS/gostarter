The XXXXXXX APIs are accessible from _/xxxxxxx/api/_

For example, `https://demo.imqs.co.za/xxxxx/api/frogs/list?type=2,3`

# GET /ping
Ping the service to see if it is alive. Returns a JSON object with the current unix time in seconds, such as `{"Timestamp": 1509373918}`

# GET /frog/list
List frogs. Returns a JSON array with the frogs that match the criteria specified in the request.

### Query Parameters
|            |                   |
| ---------- | ----------------- |
| type       | Comma separated list of frog type IDs |

# POST /frog/add
Add a new frog

Example request:
```json
[
    {"FrogTypeID": 123, "Description": "bullfrog"},
    {"FrogTypeID": 666, "Description": "toad"}
]
```