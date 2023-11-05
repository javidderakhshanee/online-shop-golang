namespace OnlineShopOrders.Core.Domain.Exceptions;

public sealed class ProductExistsInOrderException:DomainException
{
    public ProductExistsInOrderException(ulong productId,ulong orderId):base($"This product --> {productId} exists in order {orderId}.")
    {
        
    }
}