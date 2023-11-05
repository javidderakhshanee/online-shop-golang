public interface ICustomLogger
{
    void LogError(Exception exception, string message, params object[] args);
    void LogInformation(string message, params object[] args);
    void LogWarning(string message, Exception? exception = null, params object[] args);
}