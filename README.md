# LCS

A simple LaTeX conversion service

Requires `pandoc` and LaTeX to be installed

Start the server with:

```
go run main.go
```

By default, the server is accessed through port `8080`

---

## Testing locally

To test everything locally, you'll need a `sample.tex` file.

First, you'll need to make sure the directories are set up:

```
./setup.sh
```

Next, you'll need to run the server:

```
go run main.go
```

or

```
./run-server.sh
```

Finally, you can easily upload the `sample.tex` file and download the resulting `.pdf` file:

```
./run-download.sh
```

---

## Uploading a file

Files are uploaded through POST requests to the `/upload` endpoint. The file data is received with the form item name `uploadedFile`.

When a file is uploaded, the server will return the location to use to access the file. For example, if a file is uploaded to `http://localhost:8080/upload` and the server responds with `/pdf/ea9cea512170c60dce1209f6.pdf`, then the `.pdf` file can be accessed at `http://localhost:8080/pdf/ea9cea512170c60dce1209f6.pdf`.

If any errors occurred, the response will begin with `ERROR:`.

---

## Convenience scripts

A number of scripts are included for convenience

### setup-server.sh

Creates directories for the server

### setup-client.sh

Creates directory for `run-download.sh`

### setup.sh

Runs the other setup scripts

### cleanup.sh

Removes the directories created by the setup scripts

### run-server.sh

Runs the server

### run-upload.sh

Uploads `sample.tex` to the local server

### run-download.sh

Runs `run-upload.sh` and downloads the file returned by it
