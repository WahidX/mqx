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

- [x] - Add graceful shutdown
- [x] - Add topic mgmt

- Need a generic ErrorHandler
  Which will handle all the connection close erorrs

## Own datastorage design

- For each topic we will have one file where we will keep all the messages with right delimeters.
- We will have a index file for each topic which will keep the offset of the messages in the data file.
- On server start we will load all the topics and their index files in memory.
- On each message publish we will write the message in the data file and update the index file.

### Thoughts

We need to store mainly 3 things.

- The messages in [topic].msg file
- The index of the messages in file
- The consumer offsets

All messages will come and get stored in a file for the topic. All the messages will be stored in a file with a delimiter. And the index file will have the offset of the messages in the data file.

### Memory data structure

- We will have a map of topic name to topic struct. Where the topic struct will have the index file and msg file.
