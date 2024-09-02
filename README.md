# go-mq

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
