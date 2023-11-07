## Gserve

Gserve is a cli tool to create a simple, zero-configuration HTTP file server.

**Usage:**

```sh
gserve [flags]
```

**Flags:**

* `-d, --directory string` - The directory to serve files from. (default: current folder)
* `-p, --port string` - The port to serve files on. (default: 7777)
* `-v, --version` - version for gserve
* `-h, --help` - Display help message.

**Examples:**

To start a file server on the current directory on port 7777, run the following command:

```sh
gserve
```

To start a file server on a different directory, use the `-d` flag:

```sh
gserve -d /path/to/directory
```

To start a file server on a different port, use the `-p` flag:

```sh
gserve -p 8080
```
