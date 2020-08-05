# goreplacedir

Generate file outputs from go text templates under a target directory.

Pulls values out of environment variables.

```shell script
./goreplacedir-<arch> <target directory> <output directory>
```

If you have a file in a directory called "files" that looks like this:
```yaml
email:
    username: {{ .username }}
    password: {{ .password }}
```
You could generate it like this:

```shell script
export username="hello@example.com"
export password="supersecurepassword"
./goreplacedir-linux-amd64 files files_generated
```

Your output will look like this:

```yaml
email:
    username: hello@example.com
    password: supersecurepassword
```