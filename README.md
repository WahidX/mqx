# mqx

- env vars based on config file
- grpc server

## Main server apis

- `ping`
- `publish`
- `listen`
- topic mgmt apis:
  - `createTopic`
  - `listTopics`
  - `deleteTopic`
  - `getMessages(topic)`

## Data Storage

Data in each Topic will be stored in file based db. Log DB may be.
And the data can be retrived in one direction only. sort of append only logs in linked list.

### TODO

- Move from rest to binary protocol
  - [x] create net.Connection in main.go
  - [x] write a base handler
  - [] for each command forward to a handler (convert the rest handlers)
  - [] Update topicHub to use net.Conn
  - [] write the listening logic, publishing logic for handler
- Listener management

- Need a generic ErrorHandler\
  Which will handle all the connection close erorrs
