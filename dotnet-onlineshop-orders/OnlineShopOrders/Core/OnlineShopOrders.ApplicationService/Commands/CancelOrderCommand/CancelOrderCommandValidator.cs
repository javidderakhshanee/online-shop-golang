using FluentValidation;

namespace OnlineShopOrders.Core.ApplicationService.Commands.CancelOrder;

public sealed class CancelOrderCommandValidator:AbstractValidator<CancelOrderCommand>
{
    public CancelOrderCommandValidator()
    {
        RuleFor(x=>x.OrderId)
        .NotEmpty()
        .WithMessage($"{nameof(CancelOrderCommand.OrderId)} is required!");
    }
}