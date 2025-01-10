## API Usage

```bash
curl -X POST http://localhost:4000/api/auth/register -d '{"username":"testuser", "password":"testpassword", "email":"test@example.com"}' -H "Content-Type: application/json"
```

```bash
curl -X POST http://localhost:4000/api/auth/login -d '{"username":"testuser", "password":"testpassword"}' -H "Content-Type: application/json"
```

```bash
curl -H "Authorization: Bearer <JWT>" http://localhost:4000/api/user/1
```
