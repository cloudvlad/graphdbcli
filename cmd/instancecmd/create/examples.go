package create

var examples = `
graphdbcli instance create -l graphdb11.license -v 11.0.2
graphdbcli instance create -v 10.8.6
graphdbcli instance create -l graphdb11.license
graphdbcli instance create --name star-wars --license graphdb11.license
graphdbcli instance create -p 7200
graphdbcli instance create --properties ~/local.properties
graphdbcli instance create --properties /home/mark/local.properties
# Create instance (all properties specified)
graphdbcli instance create \
--name star-wars \
--license graphdb11.license \
--version 11.0.2 \
--port 7200 \ 
--activate true \
--properties ./graphdb.properties
`
