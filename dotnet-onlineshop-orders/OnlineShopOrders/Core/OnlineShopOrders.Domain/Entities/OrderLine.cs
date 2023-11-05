
namespace OnlineShopOrders.Core.Domain;

public sealed class OrderLine: ValueObject
{
    public Order Order { get; set; }
    public ulong OrderId { get; set; }
    public ulong ProductId { get; set; }
    public string ProductName { get; set; }
    public uint Quantity { get; set; }
    public decimal Price { get; set; }
    public uint Discount { get; set; }
    public decimal Total=>Price*Quantity;
    public decimal TotalWithDiscount=>Total*Discount/100;

    public override IEnumerable<object> GetAtomicValues()
    {
         yield return ProductId;
         yield return ProductName;
         yield return Quantity;
    }
}