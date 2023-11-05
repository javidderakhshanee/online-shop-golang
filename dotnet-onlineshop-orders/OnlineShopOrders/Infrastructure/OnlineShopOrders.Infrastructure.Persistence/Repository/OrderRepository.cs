using Microsoft.EntityFrameworkCore;
using OnlineShopOrders.Core.Domain;
using OnlineShopOrders.Core.Domain.Repository;
using OnlineShopOrders.Infrastructure.Configuration;
using OnlineShopOrders.Core.Domain.Exceptions;


namespace OnlineShopOrders.Infrastructure.Persistence.Repository;

public sealed class OrderRepository : IOrderRepository
{
    private readonly OnlineShopOrdersContext _context;

    public OrderRepository(OnlineShopOrdersContext context)
    {
        this._context = context;
    }
 
    public async Task<Order> Get(ulong id,CancellationToken cancellationToken)
    {
        return await _context.Orders
                     .Where(x=>x.Id==id)
                     .Include(x=>x.OrderLines)
                     .FirstOrDefaultAsync(cancellationToken);
    }

    public async Task<List<Order>> GetOrders(ulong customerId,CancellationToken cancellationToken)
    {
        return await _context.Orders
                    .Where(x=>x.CustomerId==customerId)
                    .Include(x=>x.OrderLines)
                    .OrderByDescending(x=>x.CreatedDate)
                    .ToListAsync(cancellationToken);
    }

    public async Task<Order> Save(Order order,CancellationToken cancellationToken)
    {
        _context.Orders.Add(order);

        _context.ChangeTracker.DetectChanges();

        await _context.SaveChangesAsync(cancellationToken);

        return order;
    }

    public async Task Cancel(ulong id,CancellationToken cancellationToken)
    {
        var order=await Get(id,cancellationToken);
        if(order is null) 
           throw new OrderNotFoundException(id);
        
        _context.Entry(order).State = EntityState.Deleted;

        await _context.SaveChangesAsync(cancellationToken);
    }
}