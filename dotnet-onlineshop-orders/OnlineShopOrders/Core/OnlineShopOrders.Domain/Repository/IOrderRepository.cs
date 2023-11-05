namespace OnlineShopOrders.Core.Domain.Repository;

public interface IOrderRepository
{
    Task<List<Order>> GetOrders(ulong customerId,CancellationToken cancellationToken);
    Task<Order> Get(ulong id,CancellationToken cancellationToken);
    Task<Order> Save(Order order,CancellationToken cancellationToken);    
    Task Cancel(ulong id,CancellationToken cancellationToken);
}