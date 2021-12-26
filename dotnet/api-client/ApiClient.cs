using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using Microsoft.Identity.Client;
using Polly;
using System;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Threading;
using System.Threading.Tasks;

namespace Sample;

public class ApiClientOptions
{
    public string Instance { get; set; }
    public string TenantId { get; set; }
    public string ClientId { get; set; }
    public string ClientSecret { get; set; }
    public string Scope { get; set; }
    public string BaseAddress { get; set; }
}

public class ApiClient
{
    private readonly ApiClientOptions _options;
    private readonly HttpClient _httpClient;

    public ApiClient(IOptions<ApiClientOptions> options, HttpClient httpClient)
    {
        _options = options.Value;

        httpClient.BaseAddress = new Uri(_options.BaseAddress);

        _httpClient = httpClient;
    }

    public async Task PostDemo()
    {
        var body = new { };

        var response = await _httpClient.PostAsJsonAsync("demo", body);

        Console.WriteLine(response.StatusCode);
        Console.WriteLine(await response.Content.ReadAsStringAsync());
    }
}

public class ApiClientDelegatingHandler : DelegatingHandler
{
    private readonly ApiClientOptions _options;

    public ApiClientDelegatingHandler(IOptions<ApiClientOptions> options)
    {
        _options = options.Value;
    }

    protected override async Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
    {
        if (request.Headers.Authorization == null)
        {
            var accessToken = await GetAccessTokenAsync();
            request.Headers.Authorization = new AuthenticationHeaderValue("Bearer", accessToken);
        }

        return await base.SendAsync(request, cancellationToken);
    }

    private async Task<string> GetAccessTokenAsync()
    {
        var authority = $"{_options.Instance}/{_options.TenantId}";
        var scopes = new[] { _options.Scope };

        var app = ConfidentialClientApplicationBuilder
            .Create(_options.ClientId)
            .WithClientSecret(_options.ClientSecret)
            .WithAuthority(new Uri(authority))
            .Build();

        var authResult = await app
            .AcquireTokenForClient(scopes)
            .ExecuteAsync();

        return authResult.AccessToken;
    }
}

public static class ApiClientServiceCollectionExtensions
{
    public static IServiceCollection AddApiClient(this IServiceCollection services, Action<ApiClientOptions> setupAction)
    {
        _ = services ?? throw new ArgumentNullException(nameof(services));
        _ = setupAction ?? throw new ArgumentNullException(nameof(setupAction));

        services.AddOptions();
        services.Configure(setupAction);

        services.AddScoped<ApiClientDelegatingHandler>();

        services.AddHttpClient<ApiClient>()
            .AddHttpMessageHandler<ApiClientDelegatingHandler>()
            .AddTransientHttpErrorPolicy(p =>
                p.WaitAndRetryAsync(3, i =>
                    TimeSpan.FromSeconds(Math.Pow(2, i))));

        return services;
    }
}
