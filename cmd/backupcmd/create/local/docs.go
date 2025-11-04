package local

var examples = `# Create Backup using Username and Password
graphdbcli backup create s3 /
-u admin -p root /
-r life-data,repo2 /
--location http://localhost:7200 /
--backupName my-backup.tar

# Backup all repositories on locally hosted instance without security enabled
graphdbcli backup create local

# Specifying local directory for storing the backup
graphdbcli backup create local /
--backupSaveDirPath ./mybackup-2025

# Specifying multiple repositories on multiple lines
graphdbcli backup create local /
-r repo1 /
-r repo2 
`

var shortDescription = `Create local backups stored on the host machine`

var longDescription = `Create local backups stored on the host machine`
