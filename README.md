## Yoink

### A command line tool for making requests to a restful service

#### Usage

From the directory in which your ``config.json`` resides

### The config command

```bash
foo@bar:~$ yoink config list
```

Lists the current configuration

### Using config flags

```bash
foo@bar:~$ yoink config --set-hostname=<yourhostname>
```
Sets hostname in the configuration file


Using config to make requests

```bash
foo@bar:~$ yoink --use-config
```



