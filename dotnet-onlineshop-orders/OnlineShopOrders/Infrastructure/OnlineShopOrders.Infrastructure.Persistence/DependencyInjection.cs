using Microsoft.Extensions.DependencyInjection;
using OnlineShopOrders.Core.Domain.Repository;
using OnlineShopOrders.Infrastructure.Persistence.Repository;

namespace OnlineShopOrders.Infrastructure.Persistence;

public static class DepencenyInjection
{
     public static IServiceCollection AddPersistence(this IServiceCollection services)
    {
        return services
                .AddScoped<IOrderRepository, OrderRepository>();
    }
}