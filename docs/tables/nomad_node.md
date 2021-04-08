# Table: nomad_node

A Nomad node represents a client agent within a Nomad cluster.

## Examples

### Join Nodes With AWS Resources

```sql
select
  aws_ec2_instance.instance_id,
  aws_ec2_instance.key_name,
  nomad_node.id
from
  nomad_node
join
  aws_ec2_instance
on
  (
    nomad_node.attributes ->> 'unique.platform.aws.instance-id'
    =
    aws_ec2_instance.instance_id
  )
```
