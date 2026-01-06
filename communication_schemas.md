# API Communication Standards

## Requests

All requests require the `type` key, which is either `search` or `execute`

| Type    | Description                                       | Additional Required Keys | Optional Keys |
| ------- | ------------------------------------------------- | ------------------------ | ------------- |
| search  | Tells the handler to search using the given query |                          | `query`       |
| execute | Tells the handler to execute the given result     | `selected_id`            |

### Search example

*With no query*:

```json
{
  "type": "search",
}
```

*With query*:

```json
{
  "type": "search",
  "query": "~/Downloads"
}
```

Required keys:

- `type`: Should always be `search`.

Optional keys:

- `query`: A string defining the collector's starting environment, such as the start directory for a file searcher.

### Execute example

```json
{
  "type": "execute",
  "selected_id": 15
}
```

Required keys:

- `type`: Should always be `execute`.
- `selected_id`: The unique identifier of the selected item, (the id will be sent along with the item when search results are returned).

## Responses

All responses require the `status` key, which is either `success` or `error`.

| Status  | Description                                     | Additional Required Keys | Optional Keys         |
| ------- | ----------------------------------------------- | ------------------------ | --------------------- |
| success | Successfully executed command or search         | `payload`                |
| error   | Error occurred completing the command or search | `message`                | `status_code`, `data` |

### Success examples

#### Search success

```json
{
  "status": "success",
  "payload": ["item1", "item2", "item3", ...]
}
```

#### Execute success

```json
{
  "status": "success",
  "payload": null
}
```

Required keys: 

- `status`: Should always be `success`.
- `payload`: Should be the collector's search results for a `search` request and `null` for an `execute` success.

### Error examples

```json
{
  "status": "error",
  "message": "Unable to access specified files, permission denied.",
}
```

Required keys:

- `status`: Should always be `error`.
- `message`: A meaningful, readable message, explaining what went wrong.

Optional keys:

- `status_code`: A numeric code corresponding to the error.
- `data`: Any other information about the error, such as stack traces.

#### Status Code Table

*To be defined*

| Code | Description | Use Case |
| ---- | ----------- | -------- |