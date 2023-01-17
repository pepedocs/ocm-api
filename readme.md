# ocm-api
ocm-api is a toy CLI program that has a sole purpose for querying information about OCM APIs. It is therefore can be treated as the CLI version of [api.openshift.com](https://api.openshift.com).

# Examples
1. Query all the available clusters management APIs.

```
$ ocm-api list api clusters_mgmt | jq -r ".paths | keys[]"

/api/clusters_mgmt/v1/version_gates
/api/clusters_mgmt/v1/version_gates/{version_gate_id}
/api/clusters_mgmt/v1/versions
/api/clusters_mgmt/v1/versions/{version_id}
...

```

2. Query the clusters management APIs metadata

```
$ ocm-api list api clusters_mgmt | jq .info

{
  "version": "v1",
  "title": "clusters_mgmt",
  "license": {
    "name": "Apache 2.0",
    "url": "http://www.apache.org/licenses/LICENSE-2.0"
  },
  "contact": {
    "name": "OCM Feedback",
    "email": "ocm-feedback@redhat.com"
  }
}


```
