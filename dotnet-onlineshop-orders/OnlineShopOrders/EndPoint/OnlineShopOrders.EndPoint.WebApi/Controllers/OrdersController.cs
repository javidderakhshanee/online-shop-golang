using MediatR;
using Microsoft.AspNetCore.Mvc;
using OnlineShopOrders.Core.ApplicationService.Commands.CancelOrder;
using OnlineShopOrders.Core.ApplicationService.Commands.SaveOrder;
using OnlineShopOrders.Core.ApplicationService.Queries.GetOrder;
using OnlineShopOrders.Core.ApplicationService.Queries.GetOrders;
using OnlineShopOrders.Core.Domain;

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

    [HttpGet]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(List<Order>))]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    public  async Task<IActionResult> GetOrders([FromQuery]ulong customerId)
    {
       if(customerId==0)
         return BadRequest();

       var result=await _mediator.Send(new GetOrdersQuery(customerId));

       return Ok(result);
    }

    [HttpGet("{id}")]
    [ProducesResponseType(StatusCodes.Status200OK, Type =typeof(Order))]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    public async Task<IActionResult> GetOrder(ulong id)
    {
         var order=await _mediator.Send(new GetOrderQuery(id));
         
         return order is null ? NotFound() : Ok(order); 
    }

    [HttpPost]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(Order))]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    public async Task<IActionResult> SaveOrder(SaveOrderCommand order)
    {
       var result = await _mediator.Send(order);

       return Ok(result);
    }

    [HttpDelete("{id}")]
    [ProducesResponseType(StatusCodes.Status204NoContent)]
    public async Task<IActionResult> CancelOrder(ulong id)
    {
       await _mediator.Send(new CancelOrderCommand(id));

       return NoContent();
    }

}

