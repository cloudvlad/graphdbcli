package s3cmd

var examples = `# Create Backup using Username and Password
graphdbcli backup create s3 /
-u admin -p root /
-r life-data,repo2 /
--bucket my-bucket /
--location http://localhost:7200 /
--access-key-id 0123456789 /
--access-key-token 9876543210 /
--backupName my-backup.tar

# Create Backup using Authentication Token 
graphdbcli backup create s3 /
--authToken abcd1234567890 /
-r life-data,repo2 /
--bucket my-bucket /
--location http://localhost:7200 /
--access-key-id 0123456789 /
--access-key-token 9876543210 /
--backupName my-backup.tar

# Create Backup using Authentication Token and S3 Access Key ID
# and Access Token set as environment variable -
# S3_ACCESS_KEY_ID and S3_ACCESS_KEY_TOKEN
graphdbcli backup create s3 /
--authToken abcd1234567890 /
-r life-data,repo2 /
--bucket my-bucket /
--location http://localhost:7200 /
--backupName my-backup.tar
`

var shortDescription = `Create cloud backups utilizing S3 service`

var longDescription = `Create cloud backups that are uploaded to a S3-compatible service.
By default it uses the AWS S3 service, however this behaviour can be changed by setting the the --service flag.

Often the authentication trough the S3-compatible service happens by using Access Key ID and Access Key Token.
Those could be passed by specifying them using the --access-key-id, and --access-key-token flags respectively.
However, in cases where those cannot be based directory though the CLI, environment variables could be used
to extract them. Those are "S3_ACCESS_KEY_ID" and "S3_ACCESS_KEY_TOKEN".
`
