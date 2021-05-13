# Differential backup
This script allows to back up files by "differential backup" mode (Only files that have a different hash and path will be backed up in a new "snapshot")

## How it works
The script will create an index database to trace which files have been already backed up and when performing a backup will create a new directory based on the current day and will copy only the new or changed files

### Example
#### Init
```
test
├── backup
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
│   └── index.db
└── toBackup
    ├── 1.txt
    ├── 2
    │   ├── 2.txt
    │   └── 3.txt --> chanaged from 12-05-2021
    └── 3.txt --> new file (same hash of 3.txt)
```

## How to use
### Init
Create a new backup repository
![Init](assets/init.gif)
### Backup
Performing backup
![Init](assets/backup.gif)
### Restore
Not implemented yet

#### Support the project
[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif)](https://www.paypal.com/donate?hosted_button_id=8EWYXPED4ZU5E)