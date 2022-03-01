# Differential backup
This script allows to back up files by "differential backup" mode (Only files that have a different hash and path will be backed up in a new "snapshot")

## How it works - V3.0
The script will perform a backup by creating a new directory based on the current day and by copying only the new or changed files in the destination folder. Restore functionality will restore all contents from the chosen date to the oldest by keeping the newest files

## How it works - V2.0
The script will create an index database to trace which files have been already backed up and when performing a backup will create a new directory based on the current day and will copy only the new or changed files

### Example - V2.0
#### Init
```
test
├── backup
│   ├── IMPORTANT.txt
│   └── index.db
└── toBackup
    ├── 1.txt
    ├── 2
    │   ├── 2.txt
    │   └── 3.txt
    └── 3.txt
```
    
#### Backup on 12-05-2021
```
test
├── backup
│   ├── 12-05-2021
│   │   └── toBackup
│   │       ├── 1.txt
│   │       └── 2
│   │           ├── 2.txt
│   │           └── 3.txt
│   ├── IMPORTANT.txt
│   └── index.db
└── toBackup
    ├── 1.txt
    └── 2
        ├── 2.txt
        └── 3.txt
```

#### Backup on 13-05-2021
```
test
├── backup
│   ├── 12-05-2021
│   │   └── toBackup
│   │       ├── 1.txt
│   │       └── 2
│   │           ├── 2.txt
│   │           └── 3.txt
│   ├── 13-05-2021
│   │   └── toBackup
│   │       ├── 2
│   │       │   └── 3.txt
│   │       └── 3.txt
│   ├── IMPORTANT.txt
│   └── index.db
└── toBackup
    ├── 1.txt --> same as 12-05-2021
    ├── 2
    │   ├── 2.txt --> same as 12-05-2021
    │   └── 3.txt --> different from 12-05-2021
    └── 3.txt --> new file (same hash of 3.txt)
```
#### Restore
Back up directory
```
backup
├── 15-05-2021
│   └── toBackup
│       ├── 1.txt
│       ├── 2
│       │   ├── 2.txt
│       │   └── 3.txt --> content "ciao"
│       └── 3.txt --> content "ciao"
├── 16-05-2021
│   └── toBackup
│       ├── 2
│       │   └── 3.txt --> content "new content"
│       ├── 3.txt --> content "new content"
│       └── 4.txt
├── IMPORTANT.txt
└── index.db
```
Restoring 15-05-2021
```
restore
└── toBackup
    ├── 1.txt
    ├── 2
    │   ├── 2.txt
    │   └── 3.txt --> content "ciao"
    └── 3.txt --> content "ciao"
```
Restoring 16-05-2021
```
restore
└── toBackup
    ├── 1.txt
    ├── 2
    │   ├── 2.txt
    │   └── 3.txt --> content "new content"
    ├── 3.txt --> content "new content"
    └── 4.txt
```
## How to use - V2.0/V3.0
```
diff-backup is a backup tool that perform incremental backup for a specific directory

Usage:
  diff-backup [command]

Available Commands:
  backup      backup command will perform the backup of the <source> directory to the <destination> directory
  help        Help about any command
  init        init command will initialize the <source> directory that will used for backups
  restore     restore command will restore the backup of the <source> directory to the <destination> directory from a certain <date>
  version     Print the version number of diff-backup

Flags:
  -h, --help   help for diff-backup

```

## How to use - V1.0
### Init
Create a new backup repository
![Init](assets/init.gif)
### Backup
Performing backup
![Init](assets/backup.gif)
### Restore
Performing restore
![Init](assets/restore.gif)
### Known bugs - V1.0
Not supporting space on \<source\> and \<destination\>

## Dependencies
- sqlite drivers for your OS
- github.com/c-bata/go-prompt
- github.com/kalafut/imohash
- github.com/mattn/go-sqlite3
- github.com/schollz/progressbar/v3
- github.com/spf13/cobra

#### Support the project
[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif)](https://www.paypal.com/donate?hosted_button_id=8EWYXPED4ZU5E)