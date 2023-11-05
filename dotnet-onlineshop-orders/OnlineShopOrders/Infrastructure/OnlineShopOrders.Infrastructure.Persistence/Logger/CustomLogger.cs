
using Microsoft.Extensions.Logging;

namespace OnlineShopOrders.Infrastructure.Persistence.Logger;
public sealed class CustomLogger : ICustomLogger
{
     private readonly ILogger<CustomLogger> logger;

    public CustomLogger(ILogger<CustomLogger> logger)
    {
        this.logger = logger;
    }

    public void LogError(Exception exception, string message, params object[] args)
    {
        logger.LogError(exception, message, args);
    }

    public void LogInformation(string message, params object[] args)
    {
        logger.LogInformation(message, args);
    }

    public void LogWarning(string message, Exception? exception = null, params object[] args)
    {
        logger.LogWarning(message: message, exception: exception, args: args);
    }
}