using FluentValidation;

namespace OnlineShopOrders.Core.ApplicationService.Queries.GetOrders;

public class GetOrdersValidator : AbstractValidator<GetOrdersQuery>
{
    public GetOrdersValidator()
    {
         RuleFor(x => x.CustomerId)
            .NotEmpty()
            .WithMessage($"{nameof(GetOrdersQuery.CustomerId)} is required.");
    }
}