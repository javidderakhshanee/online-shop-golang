
namespace OnlineShopOrders.Core.Domain.Exceptions;
public sealed class OrderNotFoundException : DomainException
{
    public OrderNotFoundException(ulong orderId) : base($"Order not found {orderId}")
    {
        
    }
}