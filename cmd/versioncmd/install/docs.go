package install

var examples = `graphdbcli version install 11.0.0
# Shortcut alternative for installing a version
graphdbcli v i 11.0.0
graphdbcli version install 11.0.0 --integrity-check
# Shortcut version for doing an integrity check
graphdbcli version install 11.0.0 -c
graphdbcli version install
`

var shortDescription = `Installs a specified version of the Platform Independent Distribution File locally.`

var longDescription = `Installs a specified version of the Platform Independent Distribution File locally.

The distributions file are located at the CLI home directory, by default .gdb, at the dist/ folder.
Once installed, the distribution will be called the version name plus the zip extension.
`
