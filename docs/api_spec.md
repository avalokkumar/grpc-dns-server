# API Specification

## gRPC Service: `DnsService`

### Method: `Resolve`
- **Request**:
  - `domain` (string): The domain name to resolve.
- **Response**:
  - `ip` (string): The resolved IP address.

### Example
```json
Request:
{
  "domain": "example.com"
}

Response:
{
  "ip": "93.184.216.34"
}
```