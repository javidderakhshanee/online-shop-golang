using System.Text.Json;

namespace OnlineShopOrders.Core.Domain.Extentions;

public static class CustomExtention
{
   public static T Deserialized<T>(this string json)
    {
        if (string.IsNullOrWhiteSpace(json)) 
           return default;

        return JsonSerializer.Deserialize<T>(json);
    }
    public static string Serialize(this object obj)
    {
        return JsonSerializer.Serialize(obj);
    }
}