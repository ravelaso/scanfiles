# ScanFiles

ScanFiles is a command-line tool to search for strings in files within a specified directory.

## Requirements

- Go 1.23 or later

## Building from Source

Follow these steps to build `ScanFiles` for Linux:

1. **Clone the repository:**

   ```sh
   git clone https://github.com/ravelaso/scanfiles.git
   cd scanfiles
   ```

2. **Build the application:**

   Use the following command to build the application for Linux:

   ```sh
   GOOS=linux GOARCH=amd64 go build -o scanfiles
   ```

   This will generate a binary named `scanfiles` in the current directory.

3. **Package as a .deb file (optional):**

   If you want to distribute your application as a `.deb` package, you can package it as follows:

   - Create the necessary directory structure:

     ```sh
     mkdir -p scanfiles-linux/usr/local/bin
     mkdir -p scanfiles-linux/DEBIAN
     ```

   - Move the binary into place:

     ```sh
     mv scanfiles scanfiles-linux/usr/local/bin/
     ```

   - Create a `control` file with package metadata:

     ```sh
     cat << EOF > scanfiles-linux/DEBIAN/control
     Package: scanfiles
     Version: 1.0
     Section: utils
     Priority: optional
     Architecture: amd64
     Maintainer: Your Name <your.email@example.com>
     Description: A tool to search for strings in files.
     EOF
     ```

   - Build the `.deb` package:

     ```sh
     dpkg-deb --build scanfiles-linux
     ```

   This will create `scanfiles-linux.deb`, which can be installed using `dpkg`:

   ```sh
   sudo dpkg -i scanfiles-linux.deb
   ```

## Usage

Run `scanfiles` with a directory and search string as parameters:

```sh
scanfiles /path/to/directory searchString
```
