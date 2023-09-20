# Virtual File System (VFS) CLI

A simple command-line interface (CLI) to interact with a virtual file system. This CLI enables user management, folder management, and file management.

## Table of Contents

- [Virtual File System (VFS) CLI](#virtual-file-system-vfs-cli)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
    - [User Management](#user-management)
    - [Folder Management](#folder-management)
    - [File Management](#file-management)
    - [Input Validation and Restrictions](#input-validation-and-restrictions)
  - [Contributing](#contributing)
  - [License](#license)

## Features

- **User Management**: Register users with unique, case-insensitive usernames.
- **Folder Management**: Create, delete, and rename folders with optional descriptions. Folder names are unique within a user's scope.
- **File Management**: Create, delete, and list files within specified folders. Files have optional description fields and are unique within a folder.

## Installation

1. Make sure you have Go 1.20+ installed.
   - Check your Go version with: `go version`
   - If you haven't installed Go or need to upgrade, visit [The official Go download page](https://golang.org/dl/).

1. Navigate to the project directory and build the CLI:

   ```bash
   make build
   ```

1. To start using the VFS CLI, simply run:

   ```bash
   ./vfs
   ```

## Usage

### User Management

1. **Register a User**

    ```
    register [username]
    ```

    Register a new user with the given username.

### Folder Management

1. **Create a Folder**

    ```
    create-folder [username] [foldername] [description]?
    ```

    Create a new folder for the specified user. The description is optional.

2. **Delete a Folder**

    ```
    delete-folder [username] [foldername]
    ```

    Delete an existing folder for the specified user.

3. **List Folders**

    ```
    list-folders [username] [--sort-name|--sort-created] [asc|desc]?
    ```

    List all folders for the given user. The sort and order flags are optional.

4. **Rename a Folder**

    ```
    rename-folder [username] [foldername] [new-folder-name]
    ```

    Rename an existing folder for the specified user.

### File Management

1. **Create a File**

    ```
    create-file [username] [foldername] [filename] [description]?
    ```

    Create a new file within the specified folder for the given user. The description is optional.

2. **Delete a File**

    ```
    delete-file [username] [foldername] [filename]
    ```

    Delete an existing file within the specified folder for the given user.

3. **List Files**

    ```
    list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]?
    ```

    List all files within the specified folder for the given user. The sort and order flags are optional.

### Input Validation and Restrictions

- `[username]`, `[foldername]`, `[filename]`, and `[new-folder-name]` should not contain whitespace characters and should have a maximum length of 20 characters.
- Valid characters for `[username]`, `[foldername]`, `[filename]`, and `[new-folder-name]` are alphanumeric characters (`A-Z`, `a-z`, `0-9`) and underscores (`_`).
- `[description]` have a maximum length of 255 characters.

## Contributing

We welcome any contributions! Please fork the repository, make your changes, and submit a pull request. Make sure to write tests for any new features or changes in behavior.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
