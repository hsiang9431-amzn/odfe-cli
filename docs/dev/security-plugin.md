This issue will be used for tracking the progress of adding `odfe-security` plugin support to `odfe-cli`. Following table elaborates all the granulated tasks and respectful progress

Task | sub task | Description
-----|----------|-------------
`cobra` | \-  | code for adding cobra commands
\- | `--user` | user flag for command
\- | `--password` | password flag for command
\- | `--json-file` | input json file flag for command
Account | GET | end point `_opendistro/_security/api/account`
\- | PUT
Action groups | GET | end point `_opendistro/_security/api/actiongroups/<action-group>`
\- | DELETE
\- | PUT
\- | PATCH
\- | GET (multiple) | end point `_opendistro/_security/api/actiongroups/`
\- | PATCH (multiple)
Users | GET | end point `_opendistro/_security/api/internalusers/<username>`
\- | DELETE | end point
\- | PUT
\- | PATCH
\- | GET (multiple) | end point `_opendistro/_security/api/internalusers/`
\- | PATCH (multiple)
Roles | GET | end point `_opendistro/_security/api/roles/<role>`
\- | DELETE
\- | PUT
\- | PATCH
\- | GET (multiple) | end point `_opendistro/_security/api/roles/`
\- | PATCH (multiple)
Role mappings | GET | end point `_opendistro/_security/api/rolesmapping/<role>`
\- | DELETE
\- | PUT
\- | PATCH
\- | GET (multiple) | end point `_opendistro/_security/api/rolesmapping/`
\- | PATCH (multiple)
Tenants | GET | end point `_opendistro/_security/api/tenants/<tenant>`
\- | PATCH
\- | DELETE
\- | PUT
\- | GET (multiple) | end point `_opendistro/_security/api/tenants/`
\- | PATCH (multiple)
Configuration and Certificates| GET | end point `_opendistro/_security/api/securityconfig`
\- | PATCH 
\- | PUT | end point `_opendistro/_security/api/securityconfig/config`
Cache | DELETE | end point `_opendistro/_security/api/cache`
Health | GET | end point `_opendistro/_security/health`
