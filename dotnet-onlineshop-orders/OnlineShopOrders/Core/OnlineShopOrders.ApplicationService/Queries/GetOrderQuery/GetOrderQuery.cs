using MediatR;
using OnlineShopOrders.Core.Domain;
using OnlineShopOrders.Core.Domain.Repository;

namespace OnlineShopOrders.Core.ApplicationService.Queries.GetOrder;

public record GetOrderQuery(ulong OrderId):IRequest<Order>;

public sealed class GetOrderQueryHandler : IRequestHandler<GetOrderQuery, Order>
{
    private readonly IOrderRepository _orderRepository;

    public GetOrderQueryHandler(IOrderRepository orderRepository)
    {
        _orderRepository = orderRepository;
    }

    public async Task<Order> Handle(GetOrderQuery request, CancellationToken cancellationToken)
    {
        return await _orderRepository.Get(request.OrderId,cancellationToken);
    }
}