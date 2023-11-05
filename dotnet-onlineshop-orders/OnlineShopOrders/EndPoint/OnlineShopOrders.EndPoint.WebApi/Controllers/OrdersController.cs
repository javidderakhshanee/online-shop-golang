using MediatR;
using Microsoft.AspNetCore.Mvc;
using OnlineShopOrders.Core.ApplicationService.Commands.CancelOrder;
using OnlineShopOrders.Core.ApplicationService.Commands.SaveOrder;
using OnlineShopOrders.Core.ApplicationService.Queries.GetOrder;
using OnlineShopOrders.Core.ApplicationService.Queries.GetOrders;

namespace OnlineShopOrders.EndPoint.WebApi.Controllers;

[ApiController]
[Route("[controller]")]
public class OrdersController : ControllerBase
{
    private readonly IMediator _mediator;

    public OrdersController(IMediator mediator)
    {
        _mediator = mediator;
    }

    [HttpGet("{customerId}")]
    public  async Task<IActionResult> GetOrders(ulong customerId)
    {
       if(customerId==0)
         return BadRequest();

       var result=await _mediator.Send(new GetOrdersQuery(customerId));

       return Ok(result);
    }

    [HttpGet("{id}")]
    public async Task<IActionResult> GetOrder(ulong id)
    {
         var order=await _mediator.Send(new GetOrderQuery(id));
         
         return order is null ? NotFound() : Ok(order); 
    }

    [HttpPost]
    public async Task<IActionResult> SaveOrder(SaveOrderCommand order)
    {
       if(order is null)
         return BadRequest(); 

       var result = await _mediator.Send(order);

       return Ok(result);
    }

    [HttpDelete("{id}")]
    public async Task<IActionResult> CancelOrder(ulong id)
    {
       await _mediator.Send(new CancelOrderCommand(id));

       return NoContent();
    }

}

