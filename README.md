# steampipe-plugin-nomad

## Development

1. Install steampipe
2. Run `make install`
3. Create a `~/.steampipe/config/nomad.spc` file:

  ```hcl
  connection "nomad" {
    plugin = "local/nomad"
  }
  ```
