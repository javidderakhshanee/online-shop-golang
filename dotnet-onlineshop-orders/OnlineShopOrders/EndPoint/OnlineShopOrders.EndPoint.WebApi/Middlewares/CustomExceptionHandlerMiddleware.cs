using System.ComponentModel.DataAnnotations;
using System.Text;
using OnlineShopOrders.Core.Domain.Exceptions;
using OnlineShopOrders.Core.Domain.Extentions;

public sealed class CustomExceptionHandlerMiddleware
{
    private readonly RequestDelegate _next;

    public CustomExceptionHandlerMiddleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task Invoke(HttpContext httpContext, ICustomLogger logger)
    {
        try
        {
            await _next(httpContext);
        }
        catch (Exception ex)
        {
            await HandleExceptionAsync(httpContext, ex, logger);
        }
    }
    private static bool CanLogException(Exception exception)
    {
        if (exception is DomainException)
            return false;

        return true;
    }

    private Task HandleExceptionAsync(HttpContext context, Exception ex, ICustomLogger logger)
    {
        if (CanLogException(ex))
            logger.LogError(ex, "Uncaught exception occurred");

        if (ex.InnerException is DomainException domainException)
        {
            var messageDomain = new { ErrorType = domainException.GetType().Name, ErrorMessage = domainException?.Message };
            context.Response.StatusCode = StatusCodes.Status400BadRequest;
            context.Response.WriteAsync(messageDomain.Serialize());

            GenerateErrorMessage(context,domainException,domainException?.Message,StatusCodes.Status400BadRequest);

            return Task.CompletedTask;
        }

        if (ex is FluentValidation.ValidationException validationException)
        {
            var msg = new StringBuilder();
            foreach (var error in validationException.Errors)
                msg.Append(error.ErrorMessage);

            GenerateErrorMessage(context,validationException,string.Join(Environment.NewLine,msg),StatusCodes.Status400BadRequest);

            return Task.CompletedTask;
        }

        GenerateErrorMessage(context,ex,"Internal Server Error",StatusCodes.Status500InternalServerError);

        return Task.CompletedTask;
    }

    private static void GenerateErrorMessage(HttpContext context,Exception ex,string message,int statusCode)
    {
        var errorMessage = new { ErrorType = ex.GetType().Name, ErrorMessage = message};
        context.Response.StatusCode = statusCode;
        context.Response.WriteAsync(errorMessage.Serialize());
    }
}
