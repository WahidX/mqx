# Docs

This file will contain the assumptions which are taken to design mqx.

## Message file

### Name

Name will be in this format `files/<topic_name>.msg`

### Format

Format for each message in the file will be in this format:

```plaintext
<4 bytes mentioning message length>
<message content>
```

Later headers will be added.
