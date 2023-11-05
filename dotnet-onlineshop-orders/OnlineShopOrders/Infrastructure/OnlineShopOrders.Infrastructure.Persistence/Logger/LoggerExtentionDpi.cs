
using Microsoft.Extensions.DependencyInjection;
using OnlineShopOrders.Infrastructure.Persistence.Logger;

namespace OnlineShopOrders.Logger;

public static class LoggerExtentionDpi
{
    public static IServiceCollection AddCustomLogger(this IServiceCollection services)
    {
        return services.AddScoped<ICustomLogger, CustomLogger>();
    }
}
