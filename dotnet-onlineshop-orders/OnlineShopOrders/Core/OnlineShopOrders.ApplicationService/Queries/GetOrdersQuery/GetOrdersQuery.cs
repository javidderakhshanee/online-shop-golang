
using MediatR;
using OnlineShopOrders.Core.Domain;
using OnlineShopOrders.Core.Domain.Repository;

namespace OnlineShopOrders.Core.ApplicationService.Queries.GetOrders;

public record GetOrdersQuery(ulong CustomerId) : IRequest<List<Order>>;

public class GetOrdersHandler : IRequestHandler<GetOrdersQuery, List<Order>>
{
    private readonly IOrderRepository _orderRepository;

    public GetOrdersHandler(IOrderRepository orderRepository)
    {
        _orderRepository = orderRepository;
    }

    public async Task<List<Order>> Handle(GetOrdersQuery request, CancellationToken cancellationToken)
    {
        return  await _orderRepository.GetOrders(request.CustomerId,cancellationToken);
    }
}
