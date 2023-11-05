using Microsoft.EntityFrameworkCore;
using OnlineShopOrders.Infrastructure.Configuration;
using OnlineShopOrders.Core.ApplicationService;
using OnlineShopOrders.Infrastructure.Persistence;
using OnlineShopOrders.Logger;

var builder = WebApplication.CreateBuilder(args);

builder.Services
       .AddCustomLogger()
       .AddApplicationService()
       .AddPersistence();

 builder.Services.AddControllers();
 builder.Services.AddDbContext<OnlineShopOrdersContext>(option =>
        {
            option.UseMySQL(connectionString: builder.Configuration.GetConnectionString("OnlineShopOrdersMySqlConnection"));
            option.UseQueryTrackingBehavior(QueryTrackingBehavior.NoTracking);
        });

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseMiddleware<CustomExceptionHandlerMiddleware>();

app.UseAuthorization();

app.MapControllers();


app.Run();
