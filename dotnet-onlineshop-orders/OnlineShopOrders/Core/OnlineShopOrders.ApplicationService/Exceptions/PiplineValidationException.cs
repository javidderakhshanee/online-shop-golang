using System.ComponentModel.DataAnnotations;

namespace OnlineShopOrders.Core.ApplicationService.Exceptions;

public sealed class PiplineValidationException : ValidationException 
{
    public PiplineValidationException(string message) : base(message)
    {
    }
}