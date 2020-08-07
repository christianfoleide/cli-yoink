## Yoink

### A command line tool for making requests to a restful service

#### Usage

GET requests

```bash
foo@bar:~$ yoink resource/uri -p
```

The ``-p`` flag is for pretty-printing the result, and is optional.

POST or PUT requests

```bash
foo@bar:~$ yoink -m post resource/uri path/to/file.json
```
The ``-m`` flag should be followed by your specified request method

If the json-data you wish to send is in the current directory, only the filename (with extension) is needed.

