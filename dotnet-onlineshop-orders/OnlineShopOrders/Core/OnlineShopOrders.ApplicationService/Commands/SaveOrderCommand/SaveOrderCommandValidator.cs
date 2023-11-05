
using FluentValidation;
using OnlineShopOrders.Core.ApplicationService.Exceptions;

namespace OnlineShopOrders.Core.ApplicationService.Commands.SaveOrder;
public class SaveOrderCommandValidator : AbstractValidator<SaveOrderCommand>
{
    public SaveOrderCommandValidator()
    {
        RuleFor(v => v.CustomerId)
            .NotEmpty()
            .WithMessage($"{nameof(SaveOrderCommand.CustomerId)} is requierd!")
            .Must(x=>x>0)
            .WithMessage($"{nameof(SaveOrderCommand.CustomerId)} must be greater than zero!");

        RuleFor(v => v.AddressId)
            .NotEmpty()
            .WithMessage($"{nameof(SaveOrderCommand.AddressId)} is requierd!")
            .Must(x=>x>0)
            .WithMessage($"{nameof(SaveOrderCommand.AddressId)} must be greater than zero!");
        
        RuleForEach(x => x.OrderLines)
            .NotNull()
            .WithMessage($"{nameof(SaveOrderCommand.OrderLines)} is required!");

        RuleForEach(p => p.OrderLines)
            .ChildRules(child =>
                          {
                             child.RuleFor(x=>x.ProductId)
                             .NotEmpty()
                             .WithMessage($"{nameof(SaveOrderLine.ProductId)} is requierd!")
                             .Must(x=>x>0)
                             .WithMessage($"{nameof(SaveOrderLine.ProductId)} must be greater than zero!");
                               
                               child.RuleFor(x=>x.ProductName)
                               .NotEmpty()
                               .WithMessage($"{nameof(SaveOrderLine.ProductName)} is required!"); 

                               child.RuleFor(x=>x.Quantity)
                               .NotEmpty()
                               .WithMessage($"{nameof(SaveOrderLine.Quantity)} is requierd!")
                               .Must(x=>x>0)
                               .WithMessage($"{nameof(SaveOrderLine.Quantity)} must be greater than zero!"); 

                                child.RuleFor(x=>x.Price)
                               .NotEmpty()
                               .WithMessage($"{nameof(SaveOrderLine.Price)} is requierd!")
                               .Must(x=>x>0)
                               .WithMessage($"{nameof(SaveOrderLine.Price)} must be greater than zero!");
                          }
                       );
    }
}
