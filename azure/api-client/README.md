# Azure AD - API Client

```csharp
services.AddApiClient(o =>
{
    o.Instance = Environment.GetEnvironmentVariable("API_INSTANCE");
    o.TenantId = Environment.GetEnvironmentVariable("API_TENANT_ID");
    o.ClientId = Environment.GetEnvironmentVariable("API_CLIENT_ID");
    o.ClientSecret = Environment.GetEnvironmentVariable("API_CLIENT_SECRET");
    o.Scope = Environment.GetEnvironmentVariable("API_SCOPE");
    o.BaseAddress = Environment.GetEnvironmentVariable("API_BASE_ADDRESS");
});
```
