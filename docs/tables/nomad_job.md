# Table: nomad_job

A Nomad job represents a groups of tasks to be run within a Nomad cluster.

## Examples

### List job datacenters

```sql
select
  name,
  datacenter
from
  nomad_job
```
