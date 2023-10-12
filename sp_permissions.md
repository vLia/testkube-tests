## Users

If nothing is set on the environment level, the permissions are inherited from the organization. The environment level permissions can overwrite what is inherited from the organization.

| Org role    | Env role | List | Read private | Read public | Create | Update | Delete |
| ----------- | ---------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | 
| not in org  | not in env | | | ✓ | | | | |
| owner | no role in env, but implicitly in it | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| admin | no role in env, but implicitly in it | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| member | no role in env, but implicitly in it | | | ✓ | | | |
| biller | no role in env, but implicitly in it | | | ✓ | | | |
| owner | read | ✓ | ✓ | ✓ | | | |
| owner | run | ✓ | ✓ | ✓ | | | |
| owner | write | ✓ | ✓ | ✓ | ✓ | ✓ | |
| owner | admin | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| admin | read | ✓ | ✓ | ✓ | | | |
| admin | run | ✓ | ✓ | ✓ | | | |
| admin | write | ✓ | ✓ | ✓ | ✓ | ✓ | |
| admin | admin | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| member | read | ✓ | ✓ | ✓ | | | |
| member | run | ✓ | ✓ | ✓ | | | |
| member | write | ✓ | ✓ | ✓ | ✓ | ✓ | |
| member | admin | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| biller | read | ✓ | ✓ | ✓ | | | |
| biller | run | ✓ | ✓ | ✓ | | | |
| biller | write | ✓ | ✓ | ✓ | ✓ | ✓ | |
| biller | admin | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |

## API Tokens
They behave the same way
