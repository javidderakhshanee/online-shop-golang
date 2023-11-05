using OnlineShopOrders.Core.Domain.Exceptions;

namespace OnlineShopOrders.Core.Domain;

public sealed class Order:AggregateRoot
{
    public ulong CustomerId { get; set; }
    public ulong AddressId { get; set; }
    public DateTime CreatedDate { get; set; }
    public List<OrderLine> OrderLines{get;set;}=new();
    public void AddItem(OrderLine line)
    {
        if(OrderLines.Exists(x=>x.ProductId==line.ProductId))
             throw new ProductExistsInOrderException(line.ProductId,Id);

        OrderLines.Add(line);
    }
    public void RemoveItem(ulong productId)
    {
        var item = OrderLines.FirstOrDefault(x => x.ProductId == productId);
        if (item == null) return;

        OrderLines.Remove(item);
    }
    public decimal GetTotal()=> OrderLines.Sum(x => x.Total);
    public decimal GetTotalWithDiscount()=>OrderLines.Sum(x => x.TotalWithDiscount);

    
}
