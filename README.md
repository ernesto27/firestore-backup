# FIRESTORE BACKUP

This is a simple tool to create backups from a firestore database.

## Installation

You need to have installed golang and configured $GOPATH environment variable.

After that run this command 

```bash
go install github.com/ernesto27/firestore-backup
```

## Usage

```bash
firestore-backup [firebase-credentials.json] [collections..]
```

First parameter "firebase-credentials.json" is the config firebase file that you can download from your firebase console.

```
Project settings -> Service accounts -> Firebase admin SDK
```

"collections.." is the name of the collections that you want to backup.

## Example

```bash
firestore-backup firebase-credentials.json users posts
```

This will create two files "users.json" and "posts.json" with the data of the collections.

Example users.json
```json
{
  "fNttol9VYulnwkFSir6p": {
    "name": "name1",
    "uid": "oU4IC2RLSEQvGITjZLcPplMmjZs1"
  },
  "oZs16BA72KgmFdVoGdYh": {
    "name": "name2",
    "uid": "BJT295K2wpQkewA8kvMQeJrEhdk1"
  }
}
```