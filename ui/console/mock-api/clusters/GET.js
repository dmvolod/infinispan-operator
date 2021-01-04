module.exports = (_req, res) => {
  return res.status(200).json([{
    "name": "example-infinispan",
    "namespace": "infinispan-operator",
    "type": "DataGrid",
    "status": "HEALTHY",
    "console": "http://www.example.com",
  }, {
    "name": "new-example",
    "namespace": "default",
    "type": "Cache",
    "status": "UNKNOWN",
    "console": "http://www.example.com",
  }, {
    "name": "not-exposed-cluster",
    "namespace": "default",
    "type": "Cache",
    "status": "UNKNOWN",
    "console": "",
  },
  ]);
}
