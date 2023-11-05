using MediatR;
using OnlineShopOrders.Core.Domain;
using OnlineShopOrders.Core.Domain.Repository;

namespace OnlineShopOrders.Core.ApplicationService.Commands.SaveOrder;

public record SaveOrderLine(ulong ProductId,string ProductName,uint Quantity,decimal Price,uint Discount);
public record SaveOrderCommand : IRequest<Order>
{
    public ulong CustomerId { get;  init; }
    public ulong AddressId { get;  init; }
    public List<SaveOrderLine> OrderLines { get;  init; }
}

public sealed class SaveOrderCommandHandler : IRequestHandler<SaveOrderCommand, Order>
{
    private readonly IOrderRepository  _orderRepository;

    public SaveOrderCommandHandler(IOrderRepository orderRepository)
    {
        _orderRepository = orderRepository;
    }

    public async Task<Order> Handle(SaveOrderCommand request, CancellationToken cancellationToken)
    {
        var order = new Order{
            CustomerId= request.CustomerId,
            AddressId=request.AddressId,
            CreatedDate=DateTime.Now
        };
        
        foreach(var line in request.OrderLines.Select(x=>new OrderLine{ProductId=x.ProductId,
                                                          ProductName=x.ProductName,
                                                          Quantity=x.Quantity,
                                                          Price=x.Price,
                                                          Discount=x.Discount}))
            order.AddItem(line);

        await _orderRepository.Save(order, cancellationToken);

        return order;
    }
}
