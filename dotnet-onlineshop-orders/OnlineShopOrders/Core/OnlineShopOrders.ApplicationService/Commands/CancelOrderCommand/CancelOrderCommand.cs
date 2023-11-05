using MediatR;
using OnlineShopOrders.Core.Domain.Repository;

namespace OnlineShopOrders.Core.ApplicationService.Commands.CancelOrder;

public record CancelOrderCommand(ulong OrderId):IRequest;

public sealed class CancelOrderCommandHandler : IRequestHandler<CancelOrderCommand>
{
    private readonly IOrderRepository _orderRepository;

    public CancelOrderCommandHandler(IOrderRepository orderRepository)
    {
        _orderRepository = orderRepository;
    }
    public async Task Handle(CancelOrderCommand request, CancellationToken cancellationToken)
    {
        await _orderRepository.Cancel(request.OrderId,cancellationToken);
    }
}