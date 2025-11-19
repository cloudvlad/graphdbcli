```json
"defaultNS":{
	"name":"defaultNS",
	"label":"Default namespaces for imports('\'';'\'' delimited)",
	"value":""
},
...
"imports": {
    "label":"Imported RDF files('\'';'\'' delimited)",
    "name": "imports",
    "value": ""
},
```

The label for the propety, and the delimiter defined in it, should not be used for using the correct delimiter.
Also, the quotation might be problematic, especially relevant for long requests send via cURL.
