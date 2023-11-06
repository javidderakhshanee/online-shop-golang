using FakeItEasy;
using FluentAssertions;
using MediatR;
using Microsoft.AspNetCore.Http.HttpResults;
using Microsoft.AspNetCore.Mvc;
using OnlineShopOrders.Core.ApplicationService.Commands.SaveOrder;
using OnlineShopOrders.Core.Domain;
using OnlineShopOrders.EndPoint.WebApi.Controllers;

namespace OnlineShopOrdersTest;

public class OrdersControllerTest
{
    private readonly IMediator _mediator;
    public OrdersControllerTest()
    {
        _mediator=A.Fake<IMediator>();
    }
    
    [Fact]
    public async Task GetOrders_Return_BadRequest_Result()
    {
        //Arrange
        var controller=new OrdersController(_mediator);
        
        //Act
        var result = await controller.GetOrders(0);

        //Assert

        result.Should().BeOfType<BadRequestResult>();

    }

    [Fact]
    public async Task GetOrders_Return_List_Orders()
    {
        //Arrange
        var controller=new OrdersController(_mediator);

        //Act
        var result = await controller.GetOrders(1);

        //Assert
        
        result.Should().BeOfType<OkObjectResult>();
        result.Should().NotBeNull();
    }

    [Fact]
    public async Task SaveOrder_Params_Return_Ok_Result()
    {
        //Arrange
        var controller=new OrdersController(_mediator);
        var command=A.Fake<SaveOrderCommand>();

        //Act
        var result = await controller.SaveOrder(command);

        //Assert
        
        result.Should().BeOfType<OkObjectResult>();
    }
    
     [Fact]
    public async Task CancelOrder_Return_NoContent_Result()
    {
        //Arrange
        var controller=new OrdersController(_mediator);

        //Act
        var result = await controller.CancelOrder(1);

        //Assert
        
        result.Should().BeOfType<NoContentResult>();
    }

}